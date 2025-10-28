// test/topic_error_test.go
package test

import (
	"testing"
	
	ali_mns "github.com/aliyun/aliyun-mns-go-sdk"
)

func TestNewMNSTopic_ErrorCases(t *testing.T) {
	// 创建一个模拟的MNSClient用于测试
	endpoint := "http://1234567890123456.mns.cn-hangzhou.aliyuncs.com"
	config := ali_mns.AliMNSClientConfig{
			EndPoint:        endpoint,
			AccessKeyId:     "test-access-key-id",
			AccessKeySecret: "test-access-key-secret",
		}
	client, _ := ali_mns.NewAliMNSClientWithConfigAndOptionsWithError(config, nil)	

	t.Run("Empty topic name should return error", func(t *testing.T) {
		topic, err := ali_mns.NewMNSTopicWithError("", client)
		
		// 验证返回了错误
		if err == nil {
			t.Error("Expected error for empty topic name, but got nil")
		}
		
		// 验证没有创建topic
		if topic != nil {
			t.Error("Expected nil topic for empty name, but got a topic")
		}
		
		// 验证错误信息是否正确
		expectedErrMsg := "ali_mns: topic name could not be empty"
		if err.Error() != expectedErrMsg {
			t.Errorf("Expected error message '%s', but got '%s'", expectedErrMsg, err.Error())
		}
	})

	t.Run("Empty topic name with QPS parameter should return error", func(t *testing.T) {
		topic, err := ali_mns.NewMNSTopicWithError("", client, 100)
		
		// 验证返回了错误
		if err == nil {
			t.Error("Expected error for empty topic name with QPS parameter, but got nil")
		}
		
		// 验证没有创建topic
		if topic != nil {
			t.Error("Expected nil topic for empty name with QPS parameter, but got a topic")
		}
		
		// 验证错误信息是否正确
		expectedErrMsg := "ali_mns: topic name could not be empty"
		if err.Error() != expectedErrMsg {
			t.Errorf("Expected error message '%s', but got '%s'", expectedErrMsg, err.Error())
		}
	})

	t.Run("Nil client should not panic and return error", func(t *testing.T) {
		// 测试nil client的情况
		defer func() {
			if r := recover(); r != nil {
				t.Errorf("NewMNSTopic with nil client should not panic, but panicked with: %v", r)
			}
		}()
		
		topic, err := ali_mns.NewMNSTopicWithError("test-topic", nil)
		
		// 在实际实现中，NewMNSTopic不会检查client是否为nil，所以这里不会返回错误
		// 但我们可以验证函数不会panic并且返回了topic对象
		if topic == nil {
			t.Error("Expected topic to be created even with nil client")
		} else if topic.Name() != "test-topic" {
			t.Errorf("Expected topic name to be 'test-topic', but got '%s'", topic.Name())
		}
		
		// err应该为nil，因为NewMNSTopic不检查client是否为nil
		if err != nil {
			t.Errorf("Expected no error for nil client, but got: %v", err)
		}
	})
}

func TestAliMNSClientConfig_Region(t *testing.T) {
	endpoint := "http://1234567890123456.mns.cn-hangzhou.aliyuncs.com"
	
	t.Run("Region from endpoint when not explicitly set", func(t *testing.T) {
		config := ali_mns.AliMNSClientConfig{
			EndPoint:        endpoint,
			AccessKeyId:     "test-access-key-id",
			AccessKeySecret: "test-access-key-secret",
		}
		
		client, err := ali_mns.NewAliMNSClientWithConfigAndOptionsWithError(config, nil)
		if err != nil {
			t.Fatalf("Failed to create client: %v", err)
		}
		
		expectedRegion := "cn-hangzhou"
		actualRegion := client.GetRegion()
		if actualRegion != expectedRegion {
			t.Errorf("Expected region '%s', but got '%s'", expectedRegion, actualRegion)
		}
	})
	
	t.Run("Explicitly set region overrides parsed region", func(t *testing.T) {
		config := ali_mns.AliMNSClientConfig{
			EndPoint:        endpoint,
			AccessKeyId:     "test-access-key-id",
			AccessKeySecret: "test-access-key-secret",
		}
		
		options := &ali_mns.ClientOptions{
			Region: "cn-beijing", // Explicitly set different region
		}
		
		client, err := ali_mns.NewAliMNSClientWithConfigAndOptionsWithError(config, options)
		if err != nil {
			t.Fatalf("Failed to create client: %v", err)
		}
		
		expectedRegion := "cn-beijing"
		actualRegion := client.GetRegion()
		if actualRegion != expectedRegion {
			t.Errorf("Expected region '%s', but got '%s'", expectedRegion, actualRegion)
		}
	})
	
	t.Run("Explicitly set region with internal endpoint", func(t *testing.T) {
		internalEndpoint := "http://1234567890123456.mns.cn-hangzhou-internal.aliyuncs.com"
		config := ali_mns.AliMNSClientConfig{
			EndPoint:        internalEndpoint,
			AccessKeyId:     "test-access-key-id",
			AccessKeySecret: "test-access-key-secret",
		}
		
		options := &ali_mns.ClientOptions{
			Region: "cn-shanghai", // Explicitly set region
		}
		
		client, err := ali_mns.NewAliMNSClientWithConfigAndOptionsWithError(config, options)
		if err != nil {
			t.Fatalf("Failed to create client: %v", err)
		}
		
		expectedRegion := "cn-shanghai"
		actualRegion := client.GetRegion()
		if actualRegion != expectedRegion {
			t.Errorf("Expected region '%s', but got '%s'", expectedRegion, actualRegion)
		}
	})
}