# Aliyun MNS Go SDK

[![Github version](https://badgen.net/badge/color/1.0.9/green?label=version)](https://badgen.net/badge/color/1.0.9/green?label=version)

Aliyun MNS Go SDK 是 MNS 在 Go 编译语言的官方 SDK

## [Readme in English](README.md)

## [Change Log](CHANGELOG.md)

## 关于

- 此 Go SDK 基于[阿里云消息服务 MNS](https://www.aliyun.com/product/mns/) 官方 API 构建
- 阿里云消息服务（Message Service，简称 MNS）是一种高效、可靠、安全、便捷、可弹性扩展的分布式消息服务
- MNS 能够帮助应用开发者在他们应用的分布式组件上自由的传递数据、通知消息，构建松耦合系统
- 使用此 SDK，用户可以快速构建高可靠、高并发的一对一消费模型和一对多的发布订阅模型

## 运行环境

- Go 1.20 及以上

## 安装方法

- 执行命令 `go get github.com/aliyun/aliyun-mns-go-sdk` 获取远程代码包
- 在您的代码中通过 `import "github.com/aliyun/aliyun-mns-go-sdk"` 引入 MNS Go SDK

## 快速使用

- 下载最新版 Go SDK，进入 example 目录
- 修改 endpoint 为您自己的接入点，可登录 [MNS 控制台](https://mns.console.aliyun.com/)
  查看，具体操作，请参考[获取接入点](https://help.aliyun.com/zh/mns/user-guide/manage-queues-in-the-console?spm=a2c4g.11186623.0.i25#section-yhc-ix5-300)
- 在环境变量中设置您的 `ALIBABA_CLOUD_ACCESS_KEY_ID` 和
  `ALIBABA_CLOUD_ACCESS_KEY_SECRET`，阿里云身份验证信息在 [RAM 控制台](https://ram.console.aliyun.com/)
  创建。获取方式请参考[获取 AccessKey](https://help.aliyun.com/document_detail/53045.html?spm=a2c4g.11186623.0.i29#task-354412)
- 根据阿里云规范，您应该将 AK SK
  信息设置为环境变量使用，请参考[设置环境变量](https://help.aliyun.com/zh/sdk/developer-reference/configure-the-alibaba-cloud-accesskey-environment-variable-on-linux-macos-and-windows-systems)

## 许可协议

- Apache-2.0 请参阅[许可文件](LICENSE)