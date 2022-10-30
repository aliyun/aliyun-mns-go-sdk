package mns_test

type appConf struct {
	Url             string `json:"url"`
	AccessKeyId     string `json:"access_key_id"`
	AccessKeySecret string `json:"access_key_secret"`
}

func newAppConf() *appConf {
	return &appConf{
		Url:             "http://5623837437236586.mns.ap-southeast-1.aliyuncs.com",
		AccessKeyId:     "LTAI5t9bff26m8fi71StWN7e",
		AccessKeySecret: "LNcI1L9ZwrEr0g4kvaOBWKmN4Hie8M",
	}
}
