package mns

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gogap/errors"
)

type QueueManager interface {
	CreateSimpleQueue(queueName string) error
	CreateQueue(queueName string, delaySeconds int32, maxMessageSize int32, messageRetentionPeriod int32, visibilityTimeout int32, pollingWaitSeconds int32, slices int32) error
	SetQueueAttributes(queueName string, delaySeconds int32, maxMessageSize int32, messageRetentionPeriod int32, visibilityTimeout int32, pollingWaitSeconds int32, slices int32) error
	GetQueueAttributes(queueName string) (*QueueAttribute, error)
	DeleteQueue(queueName string) error
	ListQueue(nextMarker string, retNumber int32, prefix string) (*Queues, error)
	ListQueueDetail(nextMarker string, retNumber int32, prefix string) (*QueueDetails, error)
}

type queueManager struct {
	cli     Client
	decoder Decoder
}

var _ QueueManager = (*queueManager)(nil)

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

func checkMaxMessageSize(maxSize int32) (err error) {
	if maxSize < 1024 || maxSize > 262144 {
		err = ERR_MNS_MAX_MESSAGE_SIZE_RANGE_ERROR.New()
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

func NewQueueManager(client Client) *queueManager {
	return &queueManager{
		cli:     client,
		decoder: NewDecoder(),
	}
}

func checkAttributes(delaySeconds int32, maxMessageSize int32, messageRetentionPeriod int32, visibilityTimeout int32, pollingWaitSeconds int32) (err error) {
	if err = checkDelaySeconds(delaySeconds); err != nil {
		return
	}
	if err = checkMaxMessageSize(maxMessageSize); err != nil {
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

func (p *queueManager) CreateSimpleQueue(queueName string) (err error) {
	return p.CreateQueue(queueName, 0, 65536, 345600, 30, 0, 2)
}

func (p *queueManager) CreateQueue(queueName string, delaySeconds int32, maxMessageSize int32, messageRetentionPeriod int32, visibilityTimeout int32, pollingWaitSeconds int32, slices int32) (err error) {
	queueName = strings.TrimSpace(queueName)

	if err = checkQueueName(queueName); err != nil {
		return
	}

	if err = checkAttributes(delaySeconds,
		maxMessageSize,
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

	var code int
	code, err = send(p.cli, p.decoder, PUT, nil, &message, "queues/"+queueName, nil)

	if code == http.StatusNoContent {
		err = ERR_MNS_QUEUE_ALREADY_EXIST_AND_HAVE_SAME_ATTR.New(errors.Params{"name": queueName})
		return
	}

	return
}

func (p *queueManager) SetQueueAttributes(queueName string, delaySeconds int32, maxMessageSize int32, messageRetentionPeriod int32, visibilityTimeout int32, pollingWaitSeconds int32, slices int32) (err error) {
	queueName = strings.TrimSpace(queueName)

	if err = checkQueueName(queueName); err != nil {
		return
	}

	if err = checkAttributes(delaySeconds,
		maxMessageSize,
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

func (p *queueManager) GetQueueAttributes(queueName string) (*QueueAttribute, error) {
	queueName = strings.TrimSpace(queueName)

	if err := checkQueueName(queueName); err != nil {
		return nil, err
	}

	attr := &QueueAttribute{}
	_, err := send(p.cli, p.decoder, GET, nil, nil, "queues/"+queueName, attr)
	return attr, err
}

func (p *queueManager) DeleteQueue(queueName string) (err error) {
	queueName = strings.TrimSpace(queueName)

	if err = checkQueueName(queueName); err != nil {
		return
	}

	_, err = send(p.cli, p.decoder, DELETE, nil, nil, "queues/"+queueName, nil)

	return
}

func (p *queueManager) ListQueue(nextMarker string, retNumber int32, prefix string) (*Queues, error) {
	header := map[string]string{}
	queues := &Queues{}

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
			return nil, ERR_MNS_RET_NUMBER_RANGE_ERROR.New()
		}
	}

	prefix = strings.TrimSpace(prefix)
	if prefix != "" {
		header["x-mns-prefix"] = prefix
	}

	_, err := send(p.cli, p.decoder, GET, header, nil, "queues", queues)
	return queues, err
}

func (p *queueManager) ListQueueDetail(nextMarker string, retNumber int32, prefix string) (*QueueDetails, error) {
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
			return nil, ERR_MNS_RET_NUMBER_RANGE_ERROR.New()
		}
	}

	prefix = strings.TrimSpace(prefix)
	if prefix != "" {
		header["x-mns-prefix"] = prefix
	}

	header["x-mns-with-meta"] = "true"

	queueDetails := &QueueDetails{}
	_, err := send(p.cli, p.decoder, GET, header, nil, "queues", queueDetails)
	return queueDetails, err
}
