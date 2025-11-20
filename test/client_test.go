package test

import (
	"os"
	"testing"

	"github.com/aliyun/aliyun-mns-go-sdk"
)

func TestNewClient(t *testing.T) {
	// 临时设置环境变量
    os.Setenv("ALIBABA_CLOUD_ACCESS_KEY_ID", "test-access-key-id")
    os.Setenv("ALIBABA_CLOUD_ACCESS_KEY_SECRET", "test-access-key-secret")
    defer func() {
        // 测试结束后清理环境变量
        os.Unsetenv("ALIBABA_CLOUD_ACCESS_KEY_ID")
        os.Unsetenv("ALIBABA_CLOUD_ACCESS_KEY_SECRET")
    }() 
	// 测试使用 NewClient 创建客户端
	endpoint := "http://xxx.mns.cn-hangzhou.aliyuncs.com"
	client, err := ali_mns.NewClient(endpoint, "cn-hangzhou")
	if err != nil {
		t.Errorf("Failed to create client with NewClient: %v", err)
	}

	if client == nil {
		t.Error("Client should not be nil")
	}

	// 验证 region 是否正确设置
	region := client.GetRegion()
	if region != "cn-hangzhou" {
		t.Errorf("Expected region cn-hangzhou, got %s", region)
	}
}

func TestNewClientWithToken(t *testing.T) {
	// 测试使用 NewClientWithToken 创建客户端
	endpoint := "http://xxx.mns.cn-hangzhou.aliyuncs.com"
	client, err := ali_mns.NewClientWithToken(endpoint, "test-token", "cn-hangzhou")
	if err != nil {
		t.Errorf("Failed to create client with NewClientWithToken: %v", err)
	}

	if client == nil {
		t.Error("Client should not be nil")
	}

	// 验证 region 是否正确设置
	region := client.GetRegion()
	if region != "cn-hangzhou" {
		t.Errorf("Expected region cn-hangzhou, got %s", region)
	}
}

func TestNewAliMNSClientWithConfig(t *testing.T) {
	// 测试使用 NewAliMNSClientWithConfig 创建客户端
	endpoint := "http://xxx.mns.cn-hangzhou.aliyuncs.com"
	client, err := ali_mns.NewAliMNSClientWithConfig(ali_mns.AliMNSClientConfig{
		EndPoint:         endpoint,
		AccessKeyId:      os.Getenv("ALIBABA_CLOUD_ACCESS_KEY_ID"),
		AccessKeySecret:  os.Getenv("ALIBABA_CLOUD_ACCESS_KEY_SECRET"),
		Region:           "cn-hangzhou",
	})
	if err != nil {
		t.Errorf("Failed to create client with NewAliMNSClientWithConfig: %v", err)
	}

	if client == nil {
		t.Error("Client should not be nil")
	}

	// 验证 region 是否正确设置
	region := client.GetRegion()
	if region != "cn-hangzhou" {
		t.Errorf("Expected region cn-hangzhou, got %s", region)
	}
}

func TestNewAliMNSClientWithConfigRegionMismatch(t *testing.T) {
	// 测试 endpoint 和 Region 中的 region 信息不一致的场景
	// endpoint 中是 cn-hangzhou，但 Region 设置为 cn-beijing
	endpoint := "http://xxx.mns.cn-hangzhou.aliyuncs.com"
	client, err := ali_mns.NewAliMNSClientWithConfig(ali_mns.AliMNSClientConfig{
		EndPoint:         endpoint,
		AccessKeyId:      os.Getenv("ALIBABA_CLOUD_ACCESS_KEY_ID"),
		AccessKeySecret:  os.Getenv("ALIBABA_CLOUD_ACCESS_KEY_SECRET"),
		Region:           "cn-beijing",
	})
	
	if err != nil {
		t.Errorf("Failed to create client with mismatched regions: %v", err)
	}

	if client == nil {
		t.Error("Client should not be nil even with mismatched regions")
	}

	// 验证实际使用的 region 是配置中指定的值，而不是从 endpoint 解析的值
	region := client.GetRegion()
	if region != "cn-beijing" {
		t.Errorf("Expected region cn-beijing (from config), got %s", region)
	}
}

func TestNewAliMNSClientWithConfigWithoutRegion(t *testing.T) {
	// 测试使用 NewAliMNSClientWithConfig 创建客户端但不设置 region
	endpoint := "http://xxx.mns.cn-hangzhou.aliyuncs.com"
	_, err := ali_mns.NewAliMNSClientWithConfig(ali_mns.AliMNSClientConfig{
		EndPoint:         endpoint,
		AccessKeyId:      os.Getenv("ALIBABA_CLOUD_ACCESS_KEY_ID"),
		AccessKeySecret:  os.Getenv("ALIBABA_CLOUD_ACCESS_KEY_SECRET"),
		Region:           "",
	})
	if err == nil {
		t.Error("Expected error when region is empty, but got nil")
	}
}

func TestNewAliMNSClientWithConfigWithoutEndpoint(t *testing.T) {
	// 测试使用 NewAliMNSClientWithConfig 创建客户端但不设置 endpoint
	_, err := ali_mns.NewAliMNSClientWithConfig(ali_mns.AliMNSClientConfig{
		EndPoint:         "",
		AccessKeyId:      os.Getenv("ALIBABA_CLOUD_ACCESS_KEY_ID"),
		AccessKeySecret:  os.Getenv("ALIBABA_CLOUD_ACCESS_KEY_SECRET"),
		Region:           "cn-hangzhou",
	})
	if err == nil {
		t.Error("Expected error when endpoint is empty, but got nil")
	}
}

