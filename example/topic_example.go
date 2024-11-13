package main

import (
	"fmt"
	"time"

	"github.com/aliyun/aliyun-mns-go-sdk"
	"github.com/gogap/logs"
)

func main() {
	// Replace with your own endpoint.
	endpoint := "http://xxx.mns.cn-hangzhou.aliyuncs.com"
	queueName := "test-queue"
	topicName := "test-topic"
	queueSubName := "test-sub-queue"
	httpSubName := "test-sub-http"
	client := ali_mns.NewClient(endpoint)

	// 1. create a queue for receiving pushed messages
	queueManager := ali_mns.NewMNSQueueManager(client)
	err := queueManager.CreateSimpleQueue(queueName)
	if err != nil && !ali_mns.ERR_MNS_QUEUE_ALREADY_EXIST_AND_HAVE_SAME_ATTR.IsEqual(err) {
		fmt.Println(err)
		return
	}

	// 2. create the topic
	topicManager := ali_mns.NewMNSTopicManager(client)
	// topicManager.DeleteTopic("testTopic")
	err = topicManager.CreateSimpleTopic(topicName)
	if err != nil && !ali_mns.ERR_MNS_TOPIC_ALREADY_EXIST_AND_HAVE_SAME_ATTR.IsEqual(err) {
		fmt.Println(err)
		return
	}

	topic := ali_mns.NewMNSTopic(topicName, client)
	// 3. subscribe to topic, the endpoint is queue
	queueSub := ali_mns.MessageSubsribeRequest{
		Endpoint:            topic.GenerateQueueEndpoint(queueName),
		NotifyContentFormat: ali_mns.SIMPLIFIED,
	}

	// 4. subscribe to topic, the endpoint is HTTP(S)
	httpSub := ali_mns.MessageSubsribeRequest{
		Endpoint:            "http://www.baidu.com",
		NotifyContentFormat: ali_mns.SIMPLIFIED,
	}

	err = topic.Subscribe(queueSubName, queueSub)
	if err != nil && !ali_mns.ERR_MNS_SUBSCRIPTION_ALREADY_EXIST_AND_HAVE_SAME_ATTR.IsEqual(err) {
		fmt.Println(err)
		return
	}

	err = topic.Subscribe(httpSubName, httpSub)
	if err != nil && !ali_mns.ERR_MNS_SUBSCRIPTION_ALREADY_EXIST_AND_HAVE_SAME_ATTR.IsEqual(err) {
		fmt.Println(err)
		return
	}

	/*
			Please refer to
			https://help.aliyun.com/document_detail/27434.html
			before using mail push

			sub = ali_mns.MessageSubsribeRequest{
		        Endpoint:  topic.GenerateMailEndpoint("a@b.com"),
		        NotifyContentFormat: ali_mns.SIMPLIFIED,
		    }
		    err = topic.Subscribe("SubscriptionNameB", sub)
		    if (err != nil && !ali_mns.ERR_MNS_SUBSCRIPTION_ALREADY_EXIST_AND_HAVE_SAME_ATTR.IsEqual(err)) {
		        fmt.Println(err)
		        return
		    }
	*/

	time.Sleep(time.Duration(2) * time.Second)

	// 5. now publish message
	msg := ali_mns.MessagePublishRequest{
		MessageBody: "hello topic <\"aliyun-mns-go-sdk\">",
		MessageAttributes: &ali_mns.MessageAttributes{
			MailAttributes: &ali_mns.MailAttributes{
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

	// 6. receive the message from queue
	queue := ali_mns.NewMNSQueue(queueName, client)
	endChan := make(chan int)
	respChan := make(chan ali_mns.MessageReceiveResponse)
	errChan := make(chan error)
	go func() {
		select {
		case resp := <-respChan:
			{
				logs.Pretty("response: ", resp)
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
