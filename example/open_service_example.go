package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	ali_mns "github.com/aliyun/aliyun-mns-go-sdk"
)

type appConf struct {
	Url             string `json:"url"`
	AccessKeyId     string `json:"access_key_id"`
	AccessKeySecret string `json:"access_key_secret"`
}

func main() {
	conf := appConf{}

	if bFile, e := ioutil.ReadFile("app.conf"); e != nil {
		panic(e)
	} else {
		if e := json.Unmarshal(bFile, &conf); e != nil {
			panic(e)
		}
	}

	client := ali_mns.NewAliMNSClient(conf.Url,
		conf.AccessKeyId,
		conf.AccessKeySecret)

	if resp, err := ali_mns.NewAccountManager(client).OpenService(); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("reqId:%s, orderId:%s", resp.RequestId, resp.OrderId)
	}
}
