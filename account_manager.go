package mns

import "encoding/xml"

type AccountManager struct {
	cli     Client
	decoder Decoder
}

type OpenService struct {
	BaseResponse
	XMLName xml.Name `xml:"OpenService" json:"-"`
	OrderId string   `xml:"OrderId" json:"order_id"`
}

func NewAccountManager(client Client) *AccountManager {
	return &AccountManager{
		cli:     client,
		decoder: NewDecoder(),
	}
}

func (p *AccountManager) OpenService() (attr OpenService, err error) {
	_, err = send(p.cli, p.decoder, POST, nil, nil, "commonbuy/openservice", &attr)
	return
}
