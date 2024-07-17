# Aliyun MNS Go SDK

[![Github version](https://badgen.net/badge/color/1.0.4/green?label=version)](https://badgen.net/badge/color/1.0.4/green?label=version)

The Aliyun MNS Go SDK is the official SDK for MNS in the Go programming language

## [Readme in Simplified Chinese](README-CN.md)

## [Change Log](CHANGELOG.md)

## About

- This Go SDK is built on the official API
  of [Alibaba Cloud Message Service (MNS)](https://www.aliyun.com/product/mns/)
- Alibaba Cloud Message Service (MNS) is an efficient, reliable, secure, convenient, and elastically scalable
  distributed messaging service
- MNS enables application developers to freely transmit data and notification messages across the distributed
  components of their applications, building loosely coupled systems
- Using this SDK, users can quickly build highly reliable and concurrent one-to-one consumption models as well as
  one-to-many publish-subscribe models

## Running Environment

- Go 1.18 or above

## Installing

- Run the`go get github.com/aliyun/aliyun-mns-go-sdk` command to get the remote code package.
- Use `import "github.com/aliyun/aliyun-mns-go-sdk"` in your code to introduce MNS Go SDK package

## Getting Start

- Download the latest version of the Go SDK and enter the example directory
- Modify 'endpoint' to your own access point, which can be viewed by logging into
  the [MNS console](https://mns.console.aliyun.com/). For more detail,
  see [How to get endpoint](https://help.aliyun.com/zh/mns/user-guide/manage-queues-in-the-console?spm=a2c4g.11186623.0.i25#section-yhc-ix5-300)
- Set your `ALIBABA_CLOUD_ACCESS_KEY_ID` and
  `ALIBABA_CLOUD_ACCESS_KEY_SECRET` in the environment variables, Alibaba Cloud authentication information can be
  created in the [RAM console](https://ram.console.aliyun.com/).
  For more detail,
  see [How to get AccessKey](https://help.aliyun.com/document_detail/53045.html?spm=a2c4g.11186623.0.i29#task-354412)
- Follow the Alibaba Cloud standards and set the AK (Access Key) and SK (Secret Key) in the environment variables. For
  more details,
  see [How to set environment variables](https://help.aliyun.com/zh/sdk/developer-reference/configure-the-alibaba-cloud-accesskey-environment-variable-on-linux-macos-and-windows-systems)

## License

- Apache-2.0 see [license file](LICENSE)


