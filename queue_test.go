package mns_test

import (
	"fmt"
	mns "github.com/aliyun/aliyun-mns-go-sdk"
	"github.com/gogap/logs"
	"github.com/stretchr/testify/suite"
	"log"
	"net/http"
	"testing"
	"time"
)

type queueTestSuite struct {
	suite.Suite
}

func TestQueue(t *testing.T) {
	suite.Run(t, new(queueTestSuite))
}

func (s *queueTestSuite) TestQueueExample() {
	go func() {
		log.Println(http.ListenAndServe("localhost:8080", nil))
	}()

	conf := newAppConf()
	client := mns.NewClient(conf.Url,
		conf.AccessKeyId,
		conf.AccessKeySecret)

	msg := mns.MessageSendRequest{
		MessageBody:  "hello <\"aliyun-mns-go-sdk\">",
		DelaySeconds: 0,
		Priority:     8}

	queueManager := mns.NewMNSQueueManager(client)

	err := queueManager.CreateQueue("test", 0, 65536, 345600, 30, 0, 3)

	time.Sleep(time.Duration(2) * time.Second)

	if err != nil && !mns.ERR_MNS_QUEUE_ALREADY_EXIST_AND_HAVE_SAME_ATTR.IsEqual(err) {
		fmt.Println(err)
		return
	}

	queue := mns.NewMNSQueue("test", client)

	for i := 1; i < 10000; i++ {
		ret, err := queue.SendMessage(msg)

		go func() {
			fmt.Println(queue.QPSMonitor().QPS())
		}()

		if err != nil {
			fmt.Println(err)
		} else {
			logs.Pretty("response:", ret)
		}

		endChan := make(chan int)
		respChan := make(chan mns.MessageReceiveResponse)
		errChan := make(chan error)
		go func() {
			select {
			case resp := <-respChan:
				{
					logs.Pretty("response:", resp)
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
