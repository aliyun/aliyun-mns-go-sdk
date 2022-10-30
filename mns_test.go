package mns_test

type appConf struct {
	Url             string `json:"url"`
	AccessKeyId     string `json:"access_key_id"`
	AccessKeySecret string `json:"access_key_secret"`
}

func newAppConf() *appConf {
	return &appConf{
		Url:             "",
		AccessKeyId:     "",
		AccessKeySecret: "",
	}
}
