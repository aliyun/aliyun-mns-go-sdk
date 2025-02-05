package main

import (
	"fmt"
	ali_mns "github.com/aliyun/aliyun-mns-go-sdk"
	"time"
)

func main() {
	endpoint := "http://1202283709788407.mns.cn-hangzhou.aliyuncs.com"
	client := ali_mns.NewClient(endpoint)
	queueManager := ali_mns.NewMNSQueueManager(client)
	queueName := "test-queue-0205-3"
	err := queueManager.CreateQueueWithOptions(queueName,
		ali_mns.WithDelaySeconds(5),
		ali_mns.WithMaxMessageSize(1024),
		ali_mns.WithMessageRetentionPeriod(86400))
	time.Sleep(time.Duration(2) * time.Second)
	if err != nil && !ali_mns.ERR_MNS_QUEUE_ALREADY_EXIST_AND_HAVE_SAME_ATTR.IsEqual(err) {
		fmt.Println("create queue failed: ", err)
		return
	}

	attributes, err := queueManager.GetQueueAttributes(queueName)
	if err != nil {
		fmt.Println("get queue attributes failed: ", err)
		return
	} else {
		fmt.Println("queue attributes:", attributes)
	}

	err = queueManager.SetQueueAttributesWithOptions(queueName,
		ali_mns.WithLoggingEnabled(true),
		ali_mns.WithMessageRetentionPeriod(7200),
		ali_mns.WithPollingWaitSeconds(10))
	attributes, err = queueManager.GetQueueAttributes(queueName)
	if err != nil {
		fmt.Println("get queue attributes failed: ", err)
		return
	} else {
		fmt.Println("queue attributes after set:", attributes)
	}
}
