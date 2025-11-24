package test

import (
	"testing"

	"github.com/aliyun/aliyun-mns-go-sdk"
)

func createQueueTestClient() (ali_mns.MNSClient, error) {
	endpoint := "http://xxx.mns.cn-hangzhou.aliyuncs.com"
	return ali_mns.NewAliMNSClientWithConfig(ali_mns.AliMNSClientConfig{
		EndPoint:         endpoint,
		AccessKeyId:      "ak",
		AccessKeySecret:  "sk",
		Region:           "cn-hangzhou",
	})
}

func TestNewMNSQueue(t *testing.T) {
	client, err := createQueueTestClient()
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}

	// 测试使用 NewMNSQueue 创建队列
	queue, err := ali_mns.NewMNSQueue("test-queue", client)
	if err != nil {
		t.Errorf("Failed to create queue with NewMNSQueue: %v", err)
	}

	if queue == nil {
		t.Error("Queue should not be nil")
	}

	if queue.Name() != "test-queue" {
		t.Errorf("Expected queue name test-queue, got %s", queue.Name())
	}
}

func TestNewMNSQueueWithQPS(t *testing.T) {
	client, err := createQueueTestClient()
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}

	// 测试使用 NewMNSQueue 创建带 QPS 限制的队列
	queue, err := ali_mns.NewMNSQueue("test-queue", client, 100)
	if err != nil {
		t.Errorf("Failed to create queue with NewMNSQueue: %v", err)
	}

	if queue == nil {
		t.Error("Queue should not be nil")
	}

	if queue.Name() != "test-queue" {
		t.Errorf("Expected queue name test-queue, got %s", queue.Name())
	}
}

func TestNewMNSQueueEmptyName(t *testing.T) {
	client, err := createQueueTestClient()
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}

	// 测试使用空名称创建队列应该返回错误
	_, err = ali_mns.NewMNSQueue("", client)
	if err == nil {
		t.Error("Expected error when queue name is empty, but got nil")
	}
}