package mns

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gogap/errors"
)

type TopicClient interface {
	Name() string
	GenerateQueueEndpoint(queueName string) string
	GenerateMailEndpoint(mailAddress string) string
	PublishMessage(message *MessagePublishRequest) (*MessageSendResponse, error)
	Subscribe(subscriptionName string, message *MessageSubsribeRequest) (err error)
	SetSubscriptionAttributes(subscriptionName string, notifyStrategy NotifyStrategyType) (err error)
	GetSubscriptionAttributes(subscriptionName string) (*SubscriptionAttribute, error)
	Unsubscribe(subscriptionName string) (err error)
	ListSubscriptionByTopic(nextMarker string, retNumber int32, prefix string) (*Subscriptions, error)
	ListSubscriptionDetailByTopic(nextMarker string, retNumber int32, prefix string) (*SubscriptionDetails, error)
}

type topicClient struct {
	name    string
	client  Client
	decoder Decoder

	qpsMonitor *QPSMonitor
}

var _ TopicClient = (*topicClient)(nil)

func NewTopicClient(name string, client Client, qps ...int32) *topicClient {
	if name == "" {
		panic("mns: topic name could not be empty")
	}

	topic := new(topicClient)
	topic.client = client
	topic.name = name
	topic.decoder = NewDecoder()

	qpsLimit := DefaultTopicQPSLimit
	if qps != nil && len(qps) == 1 && qps[0] > 0 {
		qpsLimit = qps[0]
	}
	topic.qpsMonitor = NewQPSMonitor(5, qpsLimit)
	return topic
}

func (p *topicClient) Name() string {
	return p.name
}

func (p *topicClient) GenerateQueueEndpoint(queueName string) string {
	return "acs:mns:" + p.client.GetRegion() + ":" + p.client.GetAccountID() + ":queues/" + queueName
}

func (p *topicClient) GenerateMailEndpoint(mailAddress string) string {
	return "mail:directmail:" + mailAddress
}

func (p *topicClient) PublishMessage(message *MessagePublishRequest) (*MessageSendResponse, error) {
	p.qpsMonitor.checkQPS()
	resp := &MessageSendResponse{}
	_, err := send(p.client, p.decoder, POST, nil, message, fmt.Sprintf("topics/%s/%s", p.name, "messages"), resp)
	return resp, err
}

func (p *topicClient) Subscribe(subscriptionName string, message *MessageSubsribeRequest) (err error) {
	subscriptionName = strings.TrimSpace(subscriptionName)

	if err = checkTopicName(subscriptionName); err != nil {
		return
	}

	p.qpsMonitor.checkQPS()

	var code int
	code, err = send(p.client, p.decoder, PUT, nil, message, fmt.Sprintf("topics/%s/subscriptions/%s", p.name, subscriptionName), nil)

	if code == http.StatusNoContent {
		err = ERR_MNS_SUBSCRIPTION_ALREADY_EXIST_AND_HAVE_SAME_ATTR.New(errors.Params{"name": subscriptionName})
		return
	}
	return
}

func (p *topicClient) SetSubscriptionAttributes(subscriptionName string, notifyStrategy NotifyStrategyType) (err error) {
	subscriptionName = strings.TrimSpace(subscriptionName)

	if err = checkTopicName(subscriptionName); err != nil {
		return
	}

	message := SetSubscriptionAttributesRequest{
		NotifyStrategy: notifyStrategy,
	}

	p.qpsMonitor.checkQPS()
	_, err = send(p.client, p.decoder, PUT, nil, message, fmt.Sprintf("topics/%s/subscriptions/%s?metaoverride=true", p.name, subscriptionName), nil)
	return
}

func (p *topicClient) GetSubscriptionAttributes(subscriptionName string) (*SubscriptionAttribute, error) {
	subscriptionName = strings.TrimSpace(subscriptionName)

	if err := checkTopicName(subscriptionName); err != nil {
		return nil, err
	}

	attr := &SubscriptionAttribute{}
	_, err := send(p.client, p.decoder, GET, nil, nil, fmt.Sprintf("topics/%s/subscriptions/%s", p.name, subscriptionName), attr)
	return attr, err
}

func (p *topicClient) Unsubscribe(subscriptionName string) (err error) {
	subscriptionName = strings.TrimSpace(subscriptionName)

	if err = checkTopicName(subscriptionName); err != nil {
		return
	}

	_, err = send(p.client, p.decoder, DELETE, nil, nil, fmt.Sprintf("topics/%s/subscriptions/%s", p.name, subscriptionName), nil)

	return
}

func (p *topicClient) ListSubscriptionByTopic(nextMarker string, retNumber int32, prefix string) (*Subscriptions, error) {
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

	subscriptions := &Subscriptions{}
	_, err := send(p.client, p.decoder, GET, header, nil, fmt.Sprintf("topics/%s/subscriptions", p.name), subscriptions)
	return subscriptions, err
}

func (p *topicClient) ListSubscriptionDetailByTopic(nextMarker string, retNumber int32, prefix string) (*SubscriptionDetails, error) {
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

	subscriptionDetails := &SubscriptionDetails{}
	_, err := send(p.client, p.decoder, GET, header, nil, fmt.Sprintf("topics/%s/subscriptions", p.name), subscriptionDetails)
	return subscriptionDetails, err
}
