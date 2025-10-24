// test/queue_error_test.go
package test

import (
	"testing"
	
	ali_mns "github.com/aliyun/aliyun-mns-go-sdk"
)

func TestNewMNSQueue_ErrorCases(t *testing.T) {
	// 创建一个模拟的MNSClient用于测试
	endpoint := "http://1234567890123456.mns.cn-hangzhou.aliyuncs.com"
	config := ali_mns.AliMNSClientConfig{
			EndPoint:        endpoint,
			AccessKeyId:     "test-access-key-id",
			AccessKeySecret: "test-access-key-secret",
		}
	client, _ := ali_mns.NewAliMNSClientWithConfig(config)	
	t.Run("Empty queue name should return error", func(t *testing.T) {
		queue, err := ali_mns.NewMNSQueue("", client)
		
		// 验证返回了错误
		if err == nil {
			t.Error("Expected error for empty queue name, but got nil")
		}
		
		// 验证没有创建queue
		if queue != nil {
			t.Error("Expected nil queue for empty name, but got a queue")
		}
		
		// 验证错误信息是否正确
		expectedErrMsg := "ali_mns: queue name could not be empty"
		if err.Error() != expectedErrMsg {
			t.Errorf("Expected error message '%s', but got '%s'", expectedErrMsg, err.Error())
		}
	})

	t.Run("Empty queue name with QPS parameter should return error", func(t *testing.T) {
		queue, err := ali_mns.NewMNSQueue("", client, 100)
		
		// 验证返回了错误
		if err == nil {
			t.Error("Expected error for empty queue name with QPS parameter, but got nil")
		}
		
		// 验证没有创建queue
		if queue != nil {
			t.Error("Expected nil queue for empty name with QPS parameter, but got a queue")
		}
		
		// 验证错误信息是否正确
		expectedErrMsg := "ali_mns: queue name could not be empty"
		if err.Error() != expectedErrMsg {
			t.Errorf("Expected error message '%s', but got '%s'", expectedErrMsg, err.Error())
		}
	})
}