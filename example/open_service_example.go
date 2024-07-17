package main

import (
	"fmt"
	"github.com/aliyun/aliyun-mns-go-sdk"
)

func main() {
	// Replace with your own endpoint.
	endpoint := "http://xxx.mns.cn-hangzhou.aliyuncs.com"
	client := ali_mns.NewClient(endpoint)
	if resp, err := ali_mns.NewAccountManager(client).OpenService(); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("reqId: %s, orderId: %s", resp.RequestId, resp.OrderId)
	}
}
