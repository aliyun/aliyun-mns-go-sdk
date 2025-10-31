// test/client_empty_endpoint_test.go
package test

import (
	"testing"
	
	ali_mns "github.com/aliyun/aliyun-mns-go-sdk"
)

func TestNewAliMNSClientWithConfig_EmptyEndpoint(t *testing.T) {
	t.Run("Empty endpoint should return error", func(t *testing.T) {
		// 测试当EndPoint为空字符串时的情况
		config := ali_mns.AliMNSClientConfig{
			EndPoint:        "", // 空的EndPoint
			AccessKeyId:     "test-access-key-id",
			AccessKeySecret: "test-access-key-secret",
		}
		
		client, err := ali_mns.CreateAliMNSClientWithConfigAndOptions(config, nil)
		t.Logf("client: %v, err: %v", client, err)

		// 验证返回了错误
		if err == nil {
			t.Error("Expected error for empty endpoint, but got nil")
		}
		
		// 验证错误信息是否正确
		expectedErrMsg := "ali-mns: message queue url is empty"
		if err.Error() != expectedErrMsg {
			t.Errorf("Expected error message '%s', but got '%s'", expectedErrMsg, err.Error())
		}
	})

	t.Run("Nil endpoint should return error", func(t *testing.T) {
		// 测试当EndPoint为完全省略(默认零值)时的情况
		config := ali_mns.AliMNSClientConfig{
			// EndPoint未设置，使用默认零值""
			AccessKeyId:     "test-access-key-id",
			AccessKeySecret: "test-access-key-secret",
		}
		
		client, err := ali_mns.CreateAliMNSClientWithConfigAndOptions(config, nil)
		
		// 验证返回了错误
		if err == nil {
			t.Error("Expected error for nil endpoint, but got nil")
		}
		t.Logf("client: %v, err: %v", client, err)
		
		// 验证错误信息是否正确
		expectedErrMsg := "ali-mns: message queue url is empty"
		if err.Error() != expectedErrMsg {
			t.Errorf("Expected error message '%s', but got '%s'", expectedErrMsg, err.Error())
		}
	})
}