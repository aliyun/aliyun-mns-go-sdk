package mns_test

import (
	"fmt"
	mns "github.com/aliyun/aliyun-mns-go-sdk"
	"github.com/gogap/logs"
	"github.com/stretchr/testify/suite"
	"testing"
	"time"
)

type topicTestSuite struct {
	suite.Suite
}

func TestTopic(t *testing.T) {
	suite.Run(t, new(topicTestSuite))
}

func (s *topicTestSuite) TestTopicExample() {
	conf := newAppConf()
	client := mns.NewClient(conf.Url,
		conf.AccessKeyId,
		conf.AccessKeySecret)

	// 1. create a queue for receiving pushed messages
	queueManager := mns.NewQueueManager(client)
	err := queueManager.CreateSimpleQueue("testQueue")
	if err != nil && !mns.ERR_MNS_QUEUE_ALREADY_EXIST_AND_HAVE_SAME_ATTR.IsEqual(err) {
		fmt.Println(err)
		return
	}

	// 2. create the topic
	topicManager := mns.NewTopicManager(client)
	// topicManager.DeleteTopic("testTopic")
	err = topicManager.CreateSimpleTopic("testTopic")
	if err != nil && !mns.ERR_MNS_TOPIC_ALREADY_EXIST_AND_HAVE_SAME_ATTR.IsEqual(err) {
		fmt.Println(err)
		return
	}

	// 3. subscribe to topic, the endpoint is set to be a queue in this sample
	topic := mns.NewTopicClient("testTopic", client)
	sub := &mns.MessageSubsribeRequest{
		Endpoint:            topic.GenerateQueueEndpoint("testQueue"),
		NotifyContentFormat: mns.SIMPLIFIED,
	}

	// topic.Unsubscribe("SubscriptionNameA")
	err = topic.Subscribe("SubscriptionNameA", sub)
	if err != nil && !mns.ERR_MNS_SUBSCRIPTION_ALREADY_EXIST_AND_HAVE_SAME_ATTR.IsEqual(err) {
		fmt.Println(err)
		return
	}

	/*
			Please refer to
			https://help.aliyun.com/document_detail/27434.html
			before using mail push

			sub = mns.MessageSubsribeRequest{
		        Endpoint:  topic.GenerateMailEndpoint("a@b.com"),
		        NotifyContentFormat: mns.SIMPLIFIED,
		    }
		    err = topic.Subscribe("SubscriptionNameB", sub)
		    if (err != nil && !mns.ERR_MNS_SUBSCRIPTION_ALREADY_EXIST_AND_HAVE_SAME_ATTR.IsEqual(err)) {
		        fmt.Println(err)
		        return
		    }
	*/

	time.Sleep(time.Duration(2) * time.Second)

	// 4. now publish message
	msg := &mns.MessagePublishRequest{
		MessageBody: "hello topic <\"aliyun-mns-go-sdk\">",
		MessageAttributes: &mns.MessageAttributes{
			MailAttributes: &mns.MailAttributes{
				Subject:     "AAA中文",
				AccountName: "BBB",
			},
		},
	}
	_, err = topic.PublishMessage(msg)
	if err != nil {
		fmt.Println(err)
		return
	}

	// 5. receive the message from queue
	queue := mns.NewQueueClient("testQueue", client)

	endChan := make(chan int)
	respChan := make(chan *mns.MessageReceiveResponse)
	errChan := make(chan error)
	go func() {
		select {
		case resp := <-respChan:
			{
				logs.Pretty("response:", resp)
				fmt.Println("change the visibility: ", resp.ReceiptHandle)
				if ret, e := queue.ChangeMessageVisibility(resp.ReceiptHandle, 5); e != nil {
					fmt.Println(e)
				} else {
					logs.Pretty("visibility changed", ret)
					fmt.Println("delete it now: ", ret.ReceiptHandle)
					if e := queue.DeleteMessage(ret.ReceiptHandle); e != nil {
						fmt.Println(e)
					}
					endChan <- 1
				}
			}
		case err := <-errChan:
			{
				fmt.Println(err)
				endChan <- 1
			}
		}
	}()

	queue.ReceiveMessage(respChan, errChan, 30)
	<-endChan
}
