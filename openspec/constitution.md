# aliyun-mns-go-sdk Constitution — Hard Constraints

> 本文件是 MNS Go SDK 项目的约束层文档。所有规则均为**强制性**，不得以任何理由绕过。
> 每条规则尽量配机器验证手段（grep / go 命令），可在 CI 或代码评审中直接使用。
> 硬约束的自然语言描述另见 [`../AGENTS.md`](../AGENTS.md)「注意事项」；本文件是其可验证化的权威版本。

---

## C-A: 架构 / 兼容性约束

### C-A1 公开 API 不得 breaking change

导出的类型、函数、方法签名（首字母大写的标识符）不得删除或不兼容修改。确需 breaking 时，按 Go module 语义升 **major** 版本。

**验证**（人工 + review）：对比 CR diff 中导出符号的签名变化；`NewAliMNSClientWithConfig` / `AliMNSClientConfig` / `GetAccountId` / `GetRegion` 等公开入口签名必须稳定。

---

### C-A2 禁止 panic，错误一律返回 error

客户端初始化与运行路径**禁止** `panic`；错误使用 `error`（gogap/errors 或 `fmt.Errorf`）返回。

**验证**：
```bash
grep -rn "panic(" --include="*.go" . | grep -v "_test.go"
```
（预期：核心路径无 `panic`；如有须在 review 说明理由。）

---

### C-A3 Client 线程安全

`MNSClient` 实例应可被多个 goroutine 共享，内部状态初始化后只读。**属流程约束**，改动 Client 结构时须在 review 确认并发安全。

---

### C-A4 数据面 API 契约与 smq proxy 一致

SDK 封装的 HTTP 请求（path / 参数 / 签名 / header）必须与 smq proxy 数据面契约一致。改动请求构造逻辑时须对照 proxy 侧契约。**属流程约束。**

---

## C-R: 发布约束

### C-R1 版本号三处一致

`version.go` 的 `Version` 常量、`CHANGELOG.md` 顶部条目、`AGENTS.md`「项目概览」表格 **版本** 单元格必须始终一致。修改 `version.go` 版本号的 CR 必须同步更新另外两处。发布流程见 [`../docs/deployment.md`](../docs/deployment.md)。

**验证**：
```bash
# version.go Version 必须等于 CHANGELOG.md 最新条目版本
VER=$(grep -oE 'Version *= *"[0-9]+\.[0-9]+\.[0-9]+"' version.go | grep -oE '[0-9]+\.[0-9]+\.[0-9]+')
CL=$(grep -oE '^## [0-9]+\.[0-9]+\.[0-9]+' CHANGELOG.md | head -1 | grep -oE '[0-9]+\.[0-9]+\.[0-9]+')
AG=$(grep -oE '\*\*版本\*\* \| [0-9]+\.[0-9]+\.[0-9]+' AGENTS.md | grep -oE '[0-9]+\.[0-9]+\.[0-9]+')
[ "$VER" = "$CL" ] && [ "$VER" = "$AG" ] && echo "OK $VER" || echo "MISMATCH version.go=$VER changelog=$CL agents=$AG"
```

**Rationale**：升版本却漏改 CHANGELOG 是历史真实缺陷（go-sdk-idpt-endpoint-format change 首次暴露）。文档缺约束时 agent/人都会漏；本规则把它变成一条可机器校验的红线。

---

## C-S: 安全约束

### C-S1 禁止硬编码敏感信息

AccessKey / SecretKey / Token 等**禁止**硬编码在 `.go` 源码或配置中（测试用占位值如 `"ak"`/`"sk"` 除外）。

**验证**：
```bash
grep -rn "LTAI" --include="*.go" . | grep -v "_test.go"
```

---

## C-Q: 代码质量约束

### C-Q1 Go 版本与依赖管理

最低 Go 1.20，使用 `go.mod` 管理依赖；HTTP 传输使用 fasthttp（性能考量），不引入 `net/http` 作为传输层（`net/url` 仅用于解析不受此限）。

**验证**：
```bash
grep -E '^go 1\.' go.mod
```

### C-Q2 测试规范

使用标准库 `testing`，测试位于 `test/` 目录；新增 / 修改行为须有对应测试。

**验证**：
```bash
go test ./test/
```

---

## 快速验证清单

```bash
# 1. 无 panic（核心路径）
grep -rn "panic(" --include="*.go" . | grep -v "_test.go"
# 2. 无硬编码 AK
grep -rn "LTAI" --include="*.go" . | grep -v "_test.go"
# 3. 版本三处一致（见 C-R1 脚本）
# 4. 单元测试
go test ./test/
```
