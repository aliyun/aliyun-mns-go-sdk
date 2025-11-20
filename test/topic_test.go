package test

import (
	"os"
	"testing"

	"github.com/aliyun/aliyun-mns-go-sdk"
)

func createTopicTestClient() (ali_mns.MNSClient, error) {
	endpoint := "http://xxx.mns.cn-hangzhou.aliyuncs.com"
	return ali_mns.NewAliMNSClientWithConfig(ali_mns.AliMNSClientConfig{
		EndPoint:         endpoint,
		AccessKeyId:      os.Getenv("ALIBABA_CLOUD_ACCESS_KEY_ID"),
		AccessKeySecret:  os.Getenv("ALIBABA_CLOUD_ACCESS_KEY_SECRET"),
		Region:           "cn-hangzhou",
	})
}

func TestNewMNSTopic(t *testing.T) {
	client, err := createTopicTestClient()
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}

	// 测试使用 NewMNSTopic 创建主题
	topic, err := ali_mns.NewMNSTopic("test-topic", client)
	if err != nil {
		t.Errorf("Failed to create topic with NewMNSTopic: %v", err)
	}

	if topic == nil {
		t.Error("Topic should not be nil")
	}

	if topic.Name() != "test-topic" {
		t.Errorf("Expected topic name test-topic, got %s", topic.Name())
	}
}

func TestNewMNSTopicWithQPS(t *testing.T) {
	client, err := createTopicTestClient()
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}

	// 测试使用 NewMNSTopic 创建带 QPS 限制的主题
	topic, err := ali_mns.NewMNSTopic("test-topic", client, 100)
	if err != nil {
		t.Errorf("Failed to create topic with NewMNSTopic: %v", err)
	}

	if topic == nil {
		t.Error("Topic should not be nil")
	}

	if topic.Name() != "test-topic" {
		t.Errorf("Expected topic name test-topic, got %s", topic.Name())
	}
}

func TestNewMNSTopicEmptyName(t *testing.T) {
	client, err := createTopicTestClient()
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}

	// 测试使用空名称创建主题应该返回错误
	_, err = ali_mns.NewMNSTopic("", client)
	if err == nil {
		t.Error("Expected error when topic name is empty, but got nil")
	}
}

func TestGenerateQueueEndpoint(t *testing.T) {
	client, err := createTopicTestClient()
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}

	topic, err := ali_mns.NewMNSTopic("test-topic", client)
	if err != nil {
		t.Fatalf("Failed to create topic: %v", err)
	}
	
	endpoint := topic.GenerateQueueEndpoint("test-queue")
	expected := "acs:mns:cn-hangzhou:" + client.GetAccountId() + ":queues/test-queue"
	
	if endpoint != expected {
		t.Errorf("Expected endpoint %s, got %s", expected, endpoint)
	}
}

func TestGenerateMailEndpoint(t *testing.T) {
	client, err := createTopicTestClient()
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}

	topic, err := ali_mns.NewMNSTopic("test-topic", client)
	if err != nil {
		t.Fatalf("Failed to create topic: %v", err)
	}
	
	endpoint := topic.GenerateMailEndpoint("test@example.com")
	expected := "mail:directmail:test@example.com"
	
	if endpoint != expected {
		t.Errorf("Expected endpoint %s, got %s", expected, endpoint)
	}
}