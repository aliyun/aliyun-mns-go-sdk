package main

import (
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
	"time"

	"github.com/aliyun/aliyun-mns-go-sdk"
	"github.com/gogap/logs"
)

func main() {
	go func() {
		log.Println(http.ListenAndServe("localhost:8080", nil))
	}()

	// Replace with your own endpoint.
	endpoint := "http://xxx.mns.cn-hangzhou.aliyuncs.com"
	client := ali_mns.NewClient(endpoint)
	msg := ali_mns.MessageSendRequest{
		MessageBody:  "hello <\"aliyun-mns-go-sdk\">",
		DelaySeconds: 0,
		Priority:     8}

	queueManager := ali_mns.NewMNSQueueManager(client)
	queueName := "test-queue"
	err := queueManager.CreateQueue(queueName, 0, 65536, 345600, 30, 0, 3)
	time.Sleep(time.Duration(2) * time.Second)
	if err != nil && !ali_mns.ERR_MNS_QUEUE_ALREADY_EXIST_AND_HAVE_SAME_ATTR.IsEqual(err) {
		fmt.Println(err)
		return
	}

	queue := ali_mns.NewMNSQueue(queueName, client)
	for i := 1; i < 10000; i++ {
		ret, err := queue.SendMessage(msg)
		go func() {
			fmt.Println(queue.QPSMonitor().QPS())
		}()

		if err != nil {
			fmt.Println(err)
		} else {
			logs.Pretty("response: ", ret)
		}

		endChan := make(chan int)
		respChan := make(chan ali_mns.MessageReceiveResponse)
		errChan := make(chan error)
		go func() {
			select {
			case resp := <-respChan:
				{
					logs.Pretty("response: ", resp)
					logs.Debug("change the visibility: ", resp.ReceiptHandle)
					if ret, e := queue.ChangeMessageVisibility(resp.ReceiptHandle, 5); e != nil {
						fmt.Println(e)
					} else {
						logs.Pretty("visibility changed", ret)
						logs.Debug("delete it now: ", ret.ReceiptHandle)
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
}
