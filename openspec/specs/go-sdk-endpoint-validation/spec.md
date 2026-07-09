# Spec: go-sdk-endpoint-validation

> aliyun-mns-go-sdk Endpoint 初始化校验与 accountId 提取的权威行为契约（canonical）。
> 归档来源：openspec change `go-sdk-idpt-endpoint-format`。

## Requirements

### R1: Endpoint 校验不应限制域名段数

WHEN `NewAliMNSClientWithConfig` 接收合法 URI endpoint
AND endpoint 的 host 非空
THEN SDK 初始化不得因为 host 按 `.` 分割后的段数不是 5 而失败

#### Scenario: 主权云公网 Endpoint
- **GIVEN** endpoint 为 `https://123.ap-paris-idpt.mns.idptcloud06api.alibaba`
- **WHEN** 调用 `NewAliMNSClientWithConfig`
- **THEN** 返回非空 client 且无错误

#### Scenario: 主权云 OXS 内网 Endpoint
- **GIVEN** endpoint 为 `https://123.mns-intranet.ap-paris-idpt.mns.idptcloud06api.alibaba`
- **WHEN** 调用 `NewAliMNSClientWithConfig`
- **THEN** 返回非空 client 且无错误

#### Scenario: 主权云 VPC 内网 Endpoint
- **GIVEN** endpoint 为 `https://123.mns-bind-vpc.ap-paris-idpt.mns.idptcloud06api.alibaba`
- **WHEN** 调用 `NewAliMNSClientWithConfig`
- **THEN** 返回非空 client 且无错误

### R2: accountId 应从 URI host 第一段提取

WHEN endpoint 为合法 URI
THEN `GetAccountId()` 返回 host 中第一个 `.` 之前的 label
AND 不从完整 endpoint 字符串中解析 accountId

#### Scenario: 主权云 Endpoint accountId 提取
- **GIVEN** endpoint 为 `https://123.mns-bind-vpc.ap-paris-idpt.mns.idptcloud06api.alibaba`
- **WHEN** client 创建成功
- **THEN** `GetAccountId()` 返回 `123`

### R3: 裸域名 endpoint 继续兼容

WHEN endpoint 不带 `http://` 或 `https://` scheme
AND endpoint 第一段 label 非空
THEN `NewAliMNSClientWithConfig` 仍可创建 client
AND `GetAccountId()` 返回裸域名第一段 label

#### Scenario: 裸域名旧用法
- **GIVEN** endpoint 为 `xxx.mns.cn-hangzhou.aliyuncs.com`
- **WHEN** 调用 `NewAliMNSClientWithConfig`
- **THEN** 返回非空 client 且无错误
- **AND** `GetAccountId()` 返回 `xxx`

### R4: 带 scheme 但 host 为空时拒绝

WHEN endpoint 包含 URL scheme
AND URI host 为空
THEN `NewAliMNSClientWithConfig` 返回 `ali-mns: message queue url is invalid`

### R5: Region 仍以配置值为准

WHEN endpoint 中包含 region 信息
AND `AliMNSClientConfig.Region` 配置为另一个 region
THEN `GetRegion()` 返回配置中的 `Region`，不从 endpoint 解析覆盖
