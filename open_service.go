package ali_mns

type OpenServicManager struct {
	cli     MNSClient
	decoder MNSDecoder
}

func NewOpenServiceManager(client MNSClient) *OpenServicManager {
	return &OpenServicManager{
		cli:     client,
		decoder: NewAliMNSDecoder(),
	}
}
func (p *OpenServicManager) OpenService() (attr OpenService, err error) {
	_, err = send(p.cli, p.decoder, POST, nil, nil, "commonbuy/openservice", &attr)
	return
}
