# aliyun-mns-go-sdk — AI 辅助开发上下文

MNS v1 官方 Go SDK,封装 smq proxy HTTP 数据面 API(队列收发消息、主题发布订阅等),供外部客户使用。

## 项目概览

| 属性 | 值 |
|------|-----|
| **名称** | aliyun-mns-go-sdk (MNS Go SDK) |
| **语言** | Go 1.20+ |
| **模块** | github.com/aliyun/aliyun-mns-go-sdk |
| **版本** | 2.0.0 |
| **HTTP 库** | fasthttp (valyala/fasthttp) |
| **认证** | aliyun/credentials-go |
| **API 文档** | https://help.aliyun.com/zh/mns/developer-reference/ |

## 常用命令

```bash
go build ./...                       # 编译
go test ./...                        # 运行全部测试
go test -v ./test/                   # 运行测试(详细输出)
go vet ./...                         # 静态检查
go mod tidy                          # 整理依赖
```

详细开发工作流 / 本地冒烟测试 / 故障排查见 [`docs/development-guide.md`](docs/development-guide.md)。

## 项目结构

```
aliyun-mns-go-sdk/
├── *.go                    # SDK 核心代码(根包)
├── example/                # 用户样例
├── test/                   # 集成测试
├── docs/                   # Agent 文档
└── go.mod                  # Go module 配置
```

**模块依赖**:单包项目,核心依赖 fasthttp / credentials-go / gogap/errors。

## 注意事项

1. **Go 版本**:最低 Go 1.20,使用 go.mod 管理依赖
2. **HTTP 库**:使用 fasthttp 而非标准 net/http(性能考量)
3. **认证方式**:支持 AK/SK 和 Credentials Provider(aliyun/credentials-go)
4. **API 契约**:SDK 封装的 HTTP API 必须与 smq proxy 数据面保持一致
5. **向后兼容**:公开 API 不得 breaking change,遵循 Go 模块版本语义
6. **并发安全**:Client 实例应线程安全,可被多个 goroutine 共享
7. **错误处理**:使用 gogap/errors,不要 panic

## 文档索引

| 文档 | 内容 |
|------|------|
| [`docs/architecture.md`](docs/architecture.md) | SDK 架构、包结构 |
| [`docs/dependencies.md`](docs/dependencies.md) | Go module 依赖 |
| [`docs/coding-standards.md`](docs/coding-standards.md) | Go 编码规范 |
| [`docs/deployment.md`](docs/deployment.md) | 发布流程(Go module proxy) |
| [`docs/development-guide.md`](docs/development-guide.md) | 开发工作流、本地测试 |

## 全局工作区

本仓库是 mns-workspace 的子项目之一(`mns-v1-dataplane-sdks/aliyun-mns-go-sdk`)。跨仓库上下文、依赖关系和工作流参见 [全局工作区导航](../../AGENTS.md)。
