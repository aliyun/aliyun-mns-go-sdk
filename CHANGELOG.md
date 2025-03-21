# Change log

## 1.0.11
- Updated version number to 1.0.11 with no other changes.

## 1.0.10
- Resolved the issue where the region information check failed during resource creation for endpoints with the suffix `-control`.

## 1.0.9
- Fix the error [issue#26](https://github.com/aliyun/aliyun-mns-go-sdk/issues/26) where the StsTokenCredential component in the credentials package does not exist.

## 1.0.8
- Support configuring the logEnable parameter for queues during creation and attribute setting.

## 1.0.7
- Add an example of base64 encoding and decoding to queue_example.go.
- Add support for dynamic credentials.

## 1.0.6
- Added an example of HTTP endpoint subscription in `topic_example.go`.
- Added an example of HTTP authorization in `http_authorization.go`.
- Removed the check for message body size to allow for larger messages.

## 1.0.5
- update the minimum Go version declared in go.mod to fix build failures.

## 1.0.4
- add version and platform information to the user agent
- following Alibaba standards, provide new recommended methods for creating MNS client, and update the example code
- support custom maxConnsPerHost value for the client.

## 1.0.3

- support custom transport configuration

## 1.0.2

- support OpenService API

## 1.0.1

- support setting timeout
- add request id to response
