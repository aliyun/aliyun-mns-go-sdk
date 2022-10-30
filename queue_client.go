package mns

import (
	"fmt"
	"net/url"
)

var (
	DefaultNumOfMessages int32 = 16
)

type QueueClient interface {
	QPSMonitor() *QPSMonitor
	Name() string
	SendMessage(message *MessageSendRequest) (*MessageSendResponse, error)
	BatchSendMessage(messages ...*MessageSendRequest) (*BatchMessageSendResponse, error)
	ReceiveMessage(respChan chan *MessageReceiveResponse, errChan chan error, waitSeconds ...int64)
	BatchReceiveMessage(respChan chan *BatchMessageReceiveResponse, errChan chan error, numOfMessages int32, waitSeconds ...int64)
	PeekMessage(respChan chan *MessageReceiveResponse, errChan chan error)
	BatchPeekMessage(respChan chan *BatchMessageReceiveResponse, errChan chan error, numOfMessages int32)
	DeleteMessage(receiptHandle string) (err error)
	BatchDeleteMessage(receiptHandles ...string) (*BatchMessageDeleteErrorResponse, error)
	ChangeMessageVisibility(receiptHandle string, visibilityTimeout int64) (*MessageVisibilityChangeResponse, error)
}

type queueClient struct {
	name       string
	client     Client
	decoder    Decoder
	qpsMonitor *QPSMonitor
}

var _ QueueClient = (*queueClient)(nil)

func NewQueueClient(name string, client Client, qps ...int32) *queueClient {
	if name == "" {
		panic("mns: queue name could not be empty")
	}

	queue := new(queueClient)
	queue.client = client
	queue.name = name
	queue.decoder = NewDecoder()

	qpsLimit := DefaultQueueQPSLimit
	if qps != nil && len(qps) == 1 && qps[0] > 0 {
		qpsLimit = qps[0]
	}
	queue.qpsMonitor = NewQPSMonitor(5, qpsLimit)
	return queue
}

func (p *queueClient) QPSMonitor() *QPSMonitor {
	return p.qpsMonitor
}

func (p *queueClient) Name() string {
	return p.name
}

func (p *queueClient) SendMessage(message *MessageSendRequest) (*MessageSendResponse, error) {
	p.qpsMonitor.checkQPS()
	resp := &MessageSendResponse{}
	_, err := send(p.client, p.decoder, POST, nil, message, fmt.Sprintf("queues/%s/%s", p.name, "messages"), resp)
	return resp, err
}

func (p *queueClient) BatchSendMessage(messages ...*MessageSendRequest) (*BatchMessageSendResponse, error) {
	if messages == nil || len(messages) == 0 {
		return nil, nil
	}

	batchRequest := &BatchMessageSendRequest{}
	for _, message := range messages {
		batchRequest.Messages = append(batchRequest.Messages, message)
	}

	p.qpsMonitor.checkQPS()
	resp := &BatchMessageSendResponse{}
	_, err := send(p.client, NewBatchOpDecoder(&resp), POST, nil, batchRequest, fmt.Sprintf("queues/%s/%s", p.name, "messages"), resp)
	return resp, err
}

func (p *queueClient) ReceiveMessage(respChan chan *MessageReceiveResponse, errChan chan error, waitSeconds ...int64) {
	resource := fmt.Sprintf("queues/%s/%s", p.name, "messages")
	if waitSeconds != nil {
		for _, waitSecond := range waitSeconds {
			if waitSecond <= 0 {
				continue
			}
			resource = fmt.Sprintf("queues/%s/%s?waitseconds=%d", p.name, "messages", waitSecond)
			p.qpsMonitor.checkQPS()
			resp := &MessageReceiveResponse{}
			_, err := send(p.client, p.decoder, GET, nil, nil, resource, resp)
			if err != nil {
				// if no
				errChan <- err
			} else {
				respChan <- resp
				// return if success, may be too much msg accumulated
				return
			}
		}
	} else {
		p.qpsMonitor.checkQPS()
		resp := &MessageReceiveResponse{}
		_, err := send(p.client, p.decoder, GET, nil, nil, resource, resp)
		if err != nil {
			errChan <- err
		} else {
			respChan <- resp
		}
	}
	// if no message after waitSecond loop or after once try if no waitSecond offered
	return
}

func (p *queueClient) BatchReceiveMessage(respChan chan *BatchMessageReceiveResponse, errChan chan error, numOfMessages int32, waitSeconds ...int64) {
	if numOfMessages <= 0 {
		numOfMessages = DefaultNumOfMessages
	}

	resource := fmt.Sprintf("queues/%s/%s?numOfMessages=%d", p.name, "messages", numOfMessages)
	if waitSeconds != nil {
		for _, waitSecond := range waitSeconds {
			if waitSecond <= 0 {
				continue
			}
			resource = fmt.Sprintf("queues/%s/%s?numOfMessages=%d&waitseconds=%d", p.name, "messages", numOfMessages, waitSecond)
			p.qpsMonitor.checkQPS()
			resp := &BatchMessageReceiveResponse{}
			_, err := send(p.client, p.decoder, GET, nil, nil, resource, resp)
			if err != nil {
				errChan <- err
			} else {
				respChan <- resp
				return
			}
		}
	} else {
		p.qpsMonitor.checkQPS()
		resp := &BatchMessageReceiveResponse{}
		_, err := send(p.client, p.decoder, GET, nil, nil, resource, resp)
		if err != nil {
			errChan <- err
		} else {
			respChan <- resp
		}
	}
	return
}

func (p *queueClient) PeekMessage(respChan chan *MessageReceiveResponse, errChan chan error) {
	p.qpsMonitor.checkQPS()
	resp := &MessageReceiveResponse{}
	_, err := send(p.client, p.decoder, GET, nil, nil, fmt.Sprintf("queues/%s/%s?peekonly=true", p.name, "messages"), resp)
	if err != nil {
		errChan <- err
	} else {
		respChan <- resp
	}
	return
}

func (p *queueClient) BatchPeekMessage(respChan chan *BatchMessageReceiveResponse, errChan chan error, numOfMessages int32) {
	if numOfMessages <= 0 {
		numOfMessages = DefaultNumOfMessages
	}

	p.qpsMonitor.checkQPS()
	resp := &BatchMessageReceiveResponse{}
	_, err := send(p.client, p.decoder, GET, nil, nil, fmt.Sprintf("queues/%s/%s?numOfMessages=%d&peekonly=true", p.name, "messages", numOfMessages), resp)
	if err != nil {
		errChan <- err
	} else {
		respChan <- resp
	}
	return
}

func (p *queueClient) DeleteMessage(receiptHandle string) error {
	p.qpsMonitor.checkQPS()
	_, err := send(p.client, p.decoder, DELETE, nil, nil, fmt.Sprintf("queues/%s/%s?ReceiptHandle=%s", p.name, "messages", url.QueryEscape(receiptHandle)), nil)
	return err
}

func (p *queueClient) BatchDeleteMessage(receiptHandles ...string) (*BatchMessageDeleteErrorResponse, error) {
	if receiptHandles == nil || len(receiptHandles) == 0 {
		return nil, nil
	}

	handlers := &ReceiptHandles{}
	for _, handler := range receiptHandles {
		handlers.ReceiptHandles = append(handlers.ReceiptHandles, handler)
	}

	p.qpsMonitor.checkQPS()
	resp := &BatchMessageDeleteErrorResponse{}
	_, err := send(p.client, NewBatchOpDecoder(resp), DELETE, nil, handlers, fmt.Sprintf("queues/%s/%s", p.name, "messages"), nil)
	return resp, err
}

func (p *queueClient) ChangeMessageVisibility(receiptHandle string, visibilityTimeout int64) (*MessageVisibilityChangeResponse, error) {
	p.qpsMonitor.checkQPS()
	resp := &MessageVisibilityChangeResponse{}
	_, err := send(p.client, p.decoder, PUT, nil, nil, fmt.Sprintf("queues/%s/%s?ReceiptHandle=%s&VisibilityTimeout=%d", p.name, "messages", url.QueryEscape(receiptHandle), visibilityTimeout), resp)
	return resp, err
}
