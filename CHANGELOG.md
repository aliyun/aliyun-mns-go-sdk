# Change log

## 2.0.1
- Support MNS endpoint domains with a variable number of segments. `NewAliMNSClientWithConfig` no longer requires the endpoint host to have a fixed number of `.`-separated segments (previously exactly 5). The host is now derived from `net/url` `Hostname()` (with a bare-domain fallback for backward compatibility), and `accountId` is extracted from the first host label. This aligns with the Java SDK, which only validates the URI scheme and host. Backward compatible with existing endpoints.

## 2.0.0 
This is a major version upgrade that introduces breaking changes. Please read carefully before upgrading.

### ⚠️ Breaking Changes
- **Region Configuration is Now Required**: The `region` field in `AliMNSClientConfig` is now mandatory. Client initialization will fail with error "ali-mns: region is empty" if not provided.
- **Region-Dependent Subscription Endpoints**: Subscription endpoints now use the explicitly configured region instead of attempting to parse from endpoint URL.

### Other Changes
- Fixed [issue#28](https://github.com/aliyun/aliyun-mns-go-sdk/issues/28): Remove panic from client initialization, return errors instead.
- Remove deprecated client creation methods.
- Simplify client configuration with `AliMNSClientConfig`.
- Update examples to use unified `AliMNSClientConfig`.
- Fixed some spelling errors.

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