package mns_test

import (
	"fmt"
	mns "github.com/aliyun/aliyun-mns-go-sdk"
	"github.com/stretchr/testify/suite"
	"testing"
)

type openServiceTestSuite struct {
	suite.Suite
}

func TestOpenService(t *testing.T) {
	suite.Run(t, new(openServiceTestSuite))
}

func (s *openServiceTestSuite) TestOpenServiceExample() {
	conf := newAppConf()
	client := mns.NewClient(conf.Url,
		conf.AccessKeyId,
		conf.AccessKeySecret)

	if resp, err := mns.NewAccountManager(client).OpenService(); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("reqId:%s, orderId:%s", resp.RequestId, resp.OrderId)
	}
}
