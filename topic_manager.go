package mns

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gogap/errors"
)

type TopicManager interface {
	CreateSimpleTopic(topicName string) (err error)
	CreateTopic(topicName string, maxMessageSize int32, loggingEnabled bool) (err error)
	SetTopicAttributes(topicName string, maxMessageSize int32, loggingEnabled bool) (err error)
	GetTopicAttributes(topicName string) (*TopicAttribute, error)
	DeleteTopic(topicName string) (err error)
	ListTopic(nextMarker string, retNumber int32, prefix string) (*Topics, error)
	ListTopicDetail(nextMarker string, retNumber int32, prefix string) (*TopicDetails, error)
}

type topicManager struct {
	cli     Client
	decoder Decoder
}

var _ TopicManager = (*topicManager)(nil)

func checkTopicName(topicName string) (err error) {
	if len(topicName) > 256 {
		err = ERR_MNS_TOPIC_NAME_IS_TOO_LONG.New()
		return
	}
	return
}

func NewTopicManager(client Client) *topicManager {
	return &topicManager{
		cli:     client,
		decoder: NewDecoder(),
	}
}

func (p *topicManager) CreateSimpleTopic(topicName string) (err error) {
	return p.CreateTopic(topicName, 65536, false)
}

func (p *topicManager) CreateTopic(topicName string, maxMessageSize int32, loggingEnabled bool) (err error) {
	topicName = strings.TrimSpace(topicName)

	if err = checkTopicName(topicName); err != nil {
		return
	}

	if err = checkMaxMessageSize(maxMessageSize); err != nil {
		return
	}

	message := CreateTopicRequest{
		MaxMessageSize: maxMessageSize,
		LoggingEnabled: loggingEnabled,
	}

	var code int
	code, err = send(p.cli, p.decoder, PUT, nil, &message, "topics/"+topicName, nil)

	if code == http.StatusNoContent {
		err = ERR_MNS_TOPIC_ALREADY_EXIST_AND_HAVE_SAME_ATTR.New(errors.Params{"name": topicName})
		return
	}

	return
}

func (p *topicManager) SetTopicAttributes(topicName string, maxMessageSize int32, loggingEnabled bool) (err error) {
	topicName = strings.TrimSpace(topicName)

	if err = checkTopicName(topicName); err != nil {
		return
	}

	if err = checkMaxMessageSize(maxMessageSize); err != nil {
		return
	}

	message := CreateTopicRequest{
		MaxMessageSize: maxMessageSize,
		LoggingEnabled: loggingEnabled,
	}

	_, err = send(p.cli, p.decoder, PUT, nil, &message, fmt.Sprintf("topics/%s?metaoverride=true", topicName), nil)
	return
}

func (p *topicManager) GetTopicAttributes(topicName string) (*TopicAttribute, error) {
	topicName = strings.TrimSpace(topicName)

	if err := checkTopicName(topicName); err != nil {
		return nil, err
	}

	attr := &TopicAttribute{}
	_, err := send(p.cli, p.decoder, GET, nil, nil, "topics/"+topicName, attr)
	return attr, err
}

func (p *topicManager) DeleteTopic(topicName string) (err error) {
	topicName = strings.TrimSpace(topicName)

	if err = checkTopicName(topicName); err != nil {
		return
	}

	_, err = send(p.cli, p.decoder, DELETE, nil, nil, "topics/"+topicName, nil)

	return
}

func (p *topicManager) ListTopic(nextMarker string, retNumber int32, prefix string) (*Topics, error) {

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

	topics := &Topics{}
	_, err := send(p.cli, p.decoder, GET, header, nil, "topics", topics)
	return topics, err
}

func (p *topicManager) ListTopicDetail(nextMarker string, retNumber int32, prefix string) (*TopicDetails, error) {

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

	topicDetails := &TopicDetails{}
	_, err := send(p.cli, p.decoder, GET, header, nil, "topics", topicDetails)
	return topicDetails, err
}
