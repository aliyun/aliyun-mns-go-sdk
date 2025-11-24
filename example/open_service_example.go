package main

import (
	"fmt"
	"github.com/aliyun/aliyun-mns-go-sdk"
	"os"
)

func main() {
	// Replace with your own endpoint.
	endpoint := "http://xxx.mns.cn-hangzhou.aliyuncs.com"
	client, err := ali_mns.NewAliMNSClientWithConfig(ali_mns.AliMNSClientConfig{
		EndPoint:         endpoint,
		AccessKeyId:      os.Getenv("ALIBABA_CLOUD_ACCESS_KEY_ID"),
		AccessKeySecret:  os.Getenv("ALIBABA_CLOUD_ACCESS_KEY_SECRET"),
		Region:           "cn-hangzhou",
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	if resp, err := ali_mns.NewAccountManager(client).OpenService(); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("reqId: %s, orderId: %s", resp.RequestId, resp.OrderId)
	}
}