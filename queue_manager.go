package ali_mns

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gogap/errors"
)

type AliQueueManager interface {
	CreateSimpleQueue(queueName string) (err error)
	CreateQueue(queueName string, delaySeconds int32, maxMessageSize int32, messageRetentionPeriod int32, visibilityTimeout int32, pollingWaitSeconds int32, slices int32) (err error)
	CreateQueueWithOptions(queueName string, options ...QueueOption) (err error)
	SetQueueAttributes(queueName string, delaySeconds int32, maxMessageSize int32, messageRetentionPeriod int32, visibilityTimeout int32, pollingWaitSeconds int32, slices int32) (err error)
	SetQueueAttributesWithOptions(queueName string, options ...QueueOption) (err error)
	GetQueueAttributes(queueName string) (attr QueueAttribute, err error)
	DeleteQueue(queueName string) (err error)
	ListQueue(nextMarker string, retNumber int32, prefix string) (queues Queues, err error)
	ListQueueDetail(nextMarker string, retNumber int32, prefix string) (queueDetails QueueDetails, err error)
}

type MNSQueueManager struct {
	cli     MNSClient
	decoder MNSDecoder
}

type QueueOptions struct {
	delaySeconds           int32
	maxMessageSize         int32
	messageRetentionPeriod int32
	visibilityTimeout      int32
	pollingWaitSeconds     int32
	loggingEnabled         bool
}

type QueueOption func(*QueueOptions, map[string]bool)

func WithDelaySeconds(delay int32) QueueOption {
	return func(o *QueueOptions, tracker map[string]bool) {
		o.delaySeconds = delay
		tracker["delaySeconds"] = true
	}
}

func WithMaxMessageSize(size int32) QueueOption {
	return func(o *QueueOptions, tracker map[string]bool) {
		o.maxMessageSize = size
		tracker["maxMessageSize"] = true
	}
}

func WithMessageRetentionPeriod(period int32) QueueOption {
	return func(o *QueueOptions, tracker map[string]bool) {
		o.messageRetentionPeriod = period
		tracker["messageRetentionPeriod"] = true
	}
}

func WithVisibilityTimeout(timeout int32) QueueOption {
	return func(o *QueueOptions, tracker map[string]bool) {
		o.visibilityTimeout = timeout
		tracker["visibilityTimeout"] = true
	}
}

func WithPollingWaitSeconds(seconds int32) QueueOption {
	return func(o *QueueOptions, tracker map[string]bool) {
		o.pollingWaitSeconds = seconds
		tracker["pollingWaitSeconds"] = true
	}
}

func WithLoggingEnabled(enabled bool) QueueOption {
	return func(o *QueueOptions, tracker map[string]bool) {
		o.loggingEnabled = enabled
		tracker["loggingEnabled"] = true
	}
}

func checkQueueName(queueName string) (err error) {
	if len(queueName) > 256 {
		err = ERR_MNS_QUEUE_NAME_IS_TOO_LONG.New()
		return
	}
	return
}

func checkDelaySeconds(seconds int32) (err error) {
	if seconds > 60480 || seconds < 0 {
		err = ERR_MNS_DELAY_SECONDS_RANGE_ERROR.New()
		return
	}
	return
}

func checkMessageRetentionPeriod(retentionPeriod int32) (err error) {
	if retentionPeriod < 60 || retentionPeriod > 1296000 {
		err = ERR_MNS_MSG_RETENTION_PERIOD_RANGE_ERROR.New()
		return
	}
	return
}

func checkVisibilityTimeout(visibilityTimeout int32) (err error) {
	if visibilityTimeout < 1 || visibilityTimeout > 43200 {
		err = ERR_MNS_MSG_VISIBILITY_TIMEOUT_RANGE_ERROR.New()
		return
	}
	return
}

func checkPollingWaitSeconds(pollingWaitSeconds int32) (err error) {
	if pollingWaitSeconds < 0 || pollingWaitSeconds > 30 {
		err = ERR_MNS_MSG_POOLLING_WAIT_SECONDS_RANGE_ERROR.New()
		return
	}
	return
}

func NewMNSQueueManager(client MNSClient) AliQueueManager {
	return &MNSQueueManager{
		cli:     client,
		decoder: NewAliMNSDecoder(),
	}
}

func checkAttributes(delaySeconds int32, messageRetentionPeriod int32, visibilityTimeout int32, pollingWaitSeconds int32) (err error) {
	if err = checkDelaySeconds(delaySeconds); err != nil {
		return
	}
	if err = checkMessageRetentionPeriod(messageRetentionPeriod); err != nil {
		return
	}
	if err = checkVisibilityTimeout(visibilityTimeout); err != nil {
		return
	}
	if err = checkPollingWaitSeconds(pollingWaitSeconds); err != nil {
		return
	}
	return
}

func (p *MNSQueueManager) CreateSimpleQueue(queueName string) (err error) {
	return p.CreateQueue(queueName, 0, 65536, 345600, 30, 0, 2)
}

func (p *MNSQueueManager) CreateQueue(queueName string, delaySeconds int32, maxMessageSize int32, messageRetentionPeriod int32, visibilityTimeout int32, pollingWaitSeconds int32, slices int32) (err error) {
	queueName = strings.TrimSpace(queueName)

	if err = checkQueueName(queueName); err != nil {
		return
	}

	if err = checkAttributes(delaySeconds,
		messageRetentionPeriod,
		visibilityTimeout,
		pollingWaitSeconds); err != nil {
		return
	}

	message := CreateQueueRequest{
		DelaySeconds:           delaySeconds,
		MaxMessageSize:         maxMessageSize,
		MessageRetentionPeriod: messageRetentionPeriod,
		VisibilityTimeout:      visibilityTimeout,
		PollingWaitSeconds:     pollingWaitSeconds,
		LoggingEnabled:         false,
	}

	var code int
	code, err = send(p.cli, p.decoder, PUT, nil, &message, "queues/"+queueName, nil)

	if code == http.StatusNoContent {
		err = ERR_MNS_QUEUE_ALREADY_EXIST_AND_HAVE_SAME_ATTR.New(errors.Params{"name": queueName})
		return
	}

	return
}

func (p *MNSQueueManager) CreateQueueWithOptions(queueName string, options ...QueueOption) (err error) {
	queueName = strings.TrimSpace(queueName)
	if err = checkQueueName(queueName); err != nil {
		return
	}
	opts := defaultQueueOptions()
	tracker := make(map[string]bool)
	for _, opt := range options {
		opt(&opts, tracker)
	}

	if err = checkAttributes(opts.delaySeconds, opts.messageRetentionPeriod,
		opts.visibilityTimeout, opts.pollingWaitSeconds); err != nil {
		return
	}

	message := CreateQueueRequest{
		DelaySeconds:           opts.delaySeconds,
		MaxMessageSize:         opts.maxMessageSize,
		MessageRetentionPeriod: opts.messageRetentionPeriod,
		VisibilityTimeout:      opts.visibilityTimeout,
		PollingWaitSeconds:     opts.pollingWaitSeconds,
		LoggingEnabled:         opts.loggingEnabled,
	}

	var code int
	code, err = send(p.cli, p.decoder, PUT, nil, &message, "queues/"+queueName, nil)
	if code == http.StatusNoContent {
		err = ERR_MNS_QUEUE_ALREADY_EXIST_AND_HAVE_SAME_ATTR.New(errors.Params{"name": queueName})
		return
	}

	return
}

func (p *MNSQueueManager) SetQueueAttributes(queueName string, delaySeconds int32, maxMessageSize int32, messageRetentionPeriod int32, visibilityTimeout int32, pollingWaitSeconds int32, slices int32) (err error) {
	queueName = strings.TrimSpace(queueName)

	if err = checkQueueName(queueName); err != nil {
		return
	}

	if err = checkAttributes(delaySeconds,
		messageRetentionPeriod,
		visibilityTimeout,
		pollingWaitSeconds); err != nil {
		return
	}

	message := CreateQueueRequest{
		DelaySeconds:           delaySeconds,
		MaxMessageSize:         maxMessageSize,
		MessageRetentionPeriod: messageRetentionPeriod,
		VisibilityTimeout:      visibilityTimeout,
		PollingWaitSeconds:     pollingWaitSeconds,
	}

	_, err = send(p.cli, p.decoder, PUT, nil, &message, fmt.Sprintf("queues/%s?metaoverride=true", queueName), nil)
	return
}

func (p *MNSQueueManager) SetQueueAttributesWithOptions(queueName string, options ...QueueOption) (err error) {
	queueName = strings.TrimSpace(queueName)
	if err = checkQueueName(queueName); err != nil {
		return
	}
	opts := QueueOptions{}
	tracker := make(map[string]bool)
	for _, opt := range options {
		opt(&opts, tracker)
	}

	message := CreateQueueRequest{}
	if tracker["delaySeconds"] {
		if err = checkDelaySeconds(opts.delaySeconds); err != nil {
			return
		}
		message.DelaySeconds = opts.delaySeconds
	}

	if tracker["maxMessageSize"] {
		message.MaxMessageSize = opts.maxMessageSize
	}

	if tracker["messageRetentionPeriod"] {
		if err = checkMessageRetentionPeriod(opts.messageRetentionPeriod); err != nil {
			return
		}
		message.MessageRetentionPeriod = opts.messageRetentionPeriod
	}

	if tracker["visibilityTimeout"] {
		if err = checkVisibilityTimeout(opts.visibilityTimeout); err != nil {
			return
		}
		message.VisibilityTimeout = opts.visibilityTimeout
	}

	if tracker["pollingWaitSeconds"] {
		if err = checkPollingWaitSeconds(opts.pollingWaitSeconds); err != nil {
			return
		}
		message.PollingWaitSeconds = opts.pollingWaitSeconds
	}

	if tracker["loggingEnabled"] {
		message.LoggingEnabled = opts.loggingEnabled
	}

	_, err = send(p.cli, p.decoder, PUT, nil, &message, fmt.Sprintf("queues/%s?metaoverride=true", queueName), nil)
	return
}

func (p *MNSQueueManager) GetQueueAttributes(queueName string) (attr QueueAttribute, err error) {
	queueName = strings.TrimSpace(queueName)

	if err = checkQueueName(queueName); err != nil {
		return
	}

	_, err = send(p.cli, p.decoder, GET, nil, nil, "queues/"+queueName, &attr)

	return
}

func (p *MNSQueueManager) DeleteQueue(queueName string) (err error) {
	queueName = strings.TrimSpace(queueName)

	if err = checkQueueName(queueName); err != nil {
		return
	}

	_, err = send(p.cli, p.decoder, DELETE, nil, nil, "queues/"+queueName, nil)

	return
}

func (p *MNSQueueManager) ListQueue(nextMarker string, retNumber int32, prefix string) (queues Queues, err error) {

	header := map[string]string{}

	marker := strings.TrimSpace(nextMarker)
	if len(marker) > 0 {
		if marker != "" {
			header["x-mns-marker"] = marker
		}
	}

	if retNumber > 0 {
		if retNumber >= 1 && retNumber <= 1000 {
			header["x-mns-ret-number"] = strconv.Itoa(int(retNumber))
		} else {
			err = ERR_MNS_RET_NUMBER_RANGE_ERROR.New()
			return
		}
	}

	prefix = strings.TrimSpace(prefix)
	if prefix != "" {
		header["x-mns-prefix"] = prefix
	}

	_, err = send(p.cli, p.decoder, GET, header, nil, "queues", &queues)

	return
}

func (p *MNSQueueManager) ListQueueDetail(nextMarker string, retNumber int32, prefix string) (queueDetails QueueDetails, err error) {

	header := map[string]string{}

	marker := strings.TrimSpace(nextMarker)
	if len(marker) > 0 {
		if marker != "" {
			header["x-mns-marker"] = marker
		}
	}

	if retNumber > 0 {
		if retNumber >= 1 && retNumber <= 1000 {
			header["x-mns-ret-number"] = strconv.Itoa(int(retNumber))
		} else {
			err = ERR_MNS_RET_NUMBER_RANGE_ERROR.New()
			return
		}
	}

	prefix = strings.TrimSpace(prefix)
	if prefix != "" {
		header["x-mns-prefix"] = prefix
	}

	header["x-mns-with-meta"] = "true"

	_, err = send(p.cli, p.decoder, GET, header, nil, "queues", &queueDetails)

	return
}

func defaultQueueOptions() QueueOptions {
	return QueueOptions{
		delaySeconds:           0,
		maxMessageSize:         65536,
		messageRetentionPeriod: 345600,
		visibilityTimeout:      30,
		pollingWaitSeconds:     0,
		loggingEnabled:         false,
	}
}
