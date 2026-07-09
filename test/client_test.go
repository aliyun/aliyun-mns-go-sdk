package test

import (
	"testing"

	"github.com/aliyun/aliyun-mns-go-sdk"
)

func TestNewAliMNSClientWithConfig(t *testing.T) {
	// 测试使用 NewAliMNSClientWithConfig 创建客户端
	endpoint := "http://xxx.mns.cn-hangzhou.aliyuncs.com"
	client, err := ali_mns.NewAliMNSClientWithConfig(ali_mns.AliMNSClientConfig{
		EndPoint:        endpoint,
		AccessKeyId:     "ak",
		AccessKeySecret: "sk",
		Region:          "cn-hangzhou",
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
		EndPoint:        endpoint,
		AccessKeyId:     "ak",
		AccessKeySecret: "sk",
		Region:          "cn-beijing",
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

func TestNewAliMNSClientWithConfigIDPTEndpoints(t *testing.T) {
	testCases := []struct {
		name     string
		endpoint string
	}{
		{
			name:     "idpt public endpoint",
			endpoint: "https://123.ap-paris-idpt.mns.idptcloud06api.alibaba",
		},
		{
			name:     "idpt oxs private endpoint",
			endpoint: "https://123.mns-intranet.ap-paris-idpt.mns.idptcloud06api.alibaba",
		},
		{
			name:     "idpt vpc private endpoint",
			endpoint: "https://123.mns-bind-vpc.ap-paris-idpt.mns.idptcloud06api.alibaba",
		},
		{
			name:     "idpt com root endpoint",
			endpoint: "http://123.mns-bind-vpc.ap-jakarta-idpt2.mns.idptcloud03api.com",
		},
		{
			name:     "jakarta idpt2 oxs private alibaba endpoint",
			endpoint: "https://123.mns-intranet.ap-jakarta-idpt2.mns.idptcloud03api.alibaba",
		},
		{
			name:     "jakarta idpt legacy oxs private com endpoint",
			endpoint: "https://123.ap-jakarta-idpt.mns-intranet.idptcloud02cs.com",
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			client, err := ali_mns.NewAliMNSClientWithConfig(ali_mns.AliMNSClientConfig{
				EndPoint:        tt.endpoint,
				AccessKeyId:     "ak",
				AccessKeySecret: "sk",
				Region:          "ap-paris-idpt",
			})
			if err != nil {
				t.Fatalf("Failed to create client with IDPT endpoint %s: %v", tt.endpoint, err)
			}
			if client.GetAccountId() != "123" {
				t.Fatalf("Expected accountId 123, got %s", client.GetAccountId())
			}
		})
	}
}

func TestNewAliMNSClientWithConfigBareEndpoint(t *testing.T) {
	client, err := ali_mns.NewAliMNSClientWithConfig(ali_mns.AliMNSClientConfig{
		EndPoint:        "xxx.mns.cn-hangzhou.aliyuncs.com",
		AccessKeyId:     "ak",
		AccessKeySecret: "sk",
		Region:          "cn-hangzhou",
	})
	if err != nil {
		t.Fatalf("Failed to create client with bare endpoint: %v", err)
	}
	if client.GetAccountId() != "xxx" {
		t.Fatalf("Expected accountId xxx, got %s", client.GetAccountId())
	}
}

func TestNewAliMNSClientWithConfigInvalidEndpointWithoutHost(t *testing.T) {
	_, err := ali_mns.NewAliMNSClientWithConfig(ali_mns.AliMNSClientConfig{
		EndPoint:        "http:///queues",
		AccessKeyId:     "ak",
		AccessKeySecret: "sk",
		Region:          "cn-hangzhou",
	})
	if err == nil {
		t.Error("Expected error when endpoint has scheme but no host, but got nil")
	}
}

func TestNewAliMNSClientWithConfigWithoutRegion(t *testing.T) {
	// 测试使用 NewAliMNSClientWithConfig 创建客户端但不设置 region
	endpoint := "http://xxx.mns.cn-hangzhou.aliyuncs.com"
	_, err := ali_mns.NewAliMNSClientWithConfig(ali_mns.AliMNSClientConfig{
		EndPoint:        endpoint,
		AccessKeyId:     "ak",
		AccessKeySecret: "sk",
		Region:          "",
	})
	if err == nil {
		t.Error("Expected error when region is empty, but got nil")
	}
}

func TestNewAliMNSClientWithConfigWithoutEndpoint(t *testing.T) {
	// 测试使用 NewAliMNSClientWithConfig 创建客户端但不设置 endpoint
	_, err := ali_mns.NewAliMNSClientWithConfig(ali_mns.AliMNSClientConfig{
		EndPoint:        "",
		AccessKeyId:     "ak",
		AccessKeySecret: "sk",
		Region:          "cn-hangzhou",
	})
	if err == nil {
		t.Error("Expected error when endpoint is empty, but got nil")
	}
}
