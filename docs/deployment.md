# aliyun-mns-go-sdk — 版本与发布

<!-- 本仓是 Go 库(SDK),不是可部署服务:无镜像/端口/JVM/K8s。本文档描述版本管理与对外分发流程;cheatsheet 命令在 ../AGENTS.md「常用命令」段。 -->

## 分发形态

- **产物类型**:Go module 库(`github.com/aliyun/aliyun-mns-go-sdk`),供外部客户 `import` 使用,**不产出可执行服务 / 镜像 / 端口**。
- **消费方式**:
  ```bash
  go get github.com/aliyun/aliyun-mns-go-sdk@<version>
  ```
- **分发通道**:Go module proxy(goproxy)按 module path + version 提供下载。
- **开发源 vs 发布源**:本仓(内部 gitlab `messaging/aliyun-mns-go-sdk`)是开发主仓;对外发布通过公开镜像仓 `github.com/aliyun/aliyun-mns-go-sdk` 打版本 tag。**本仓当前不承载 release tag。**

## 版本一致性(硬约束)

以下三处版本号必须始终一致,由 [`../openspec/constitution.md`](../openspec/constitution.md) C-R1 强制并提供机器校验:

| 位置 | 作用 |
|------|------|
| `version.go` 的 `Version` 常量 | SDK 运行时自报版本(拼入 User-Agent) |
| `CHANGELOG.md` 顶部条目 | 面向用户的版本变更说明 |
| `AGENTS.md`「项目概览」表格 **版本** 单元格 | Agent 上下文中的版本标识 |

## 发布 checklist

1. 在功能 / 修复分支完成代码改动并通过 `go test ./test/`。
2. 递增 `version.go` 的 `Version`(遵循 SemVer:修复 → patch,兼容新增 → minor,breaking → major)。
3. 在 `CHANGELOG.md` 顶部追加与 `version.go` 一致的版本条目,列出对外可见的变更 / breaking / 兼容性说明。
4. 同步更新 `AGENTS.md`「项目概览」表格的 **版本** 单元格。
5. 上述 2-4 与代码改动放入**同一 MR / CR**,评审合并。
6. 在公开镜像仓按 `vX.Y.Z` 语义化版本打 tag 并发布 release;Go module proxy 收录后用户即可 `go get` 到新版本。

## 兼容性要点

- 公开 API(导出的类型 / 函数 / 方法签名)不得 breaking change(constitution C-A1);确需 breaking 时按 Go module 语义升 major 版本。
- endpoint 校验等初始化行为的**放宽**属兼容性增强(如 2.0.1 支持主权云 IDPT endpoint),须在 `CHANGELOG.md` 标注"向后兼容"。
- 回滚方式:代码版本回滚——消费方 `go get` 回退到上一版本即可,无运行时状态。
# Aliyun Mns Go Sdk — 部署与运行

<!-- 范围: K8s/Helm chart 名 + 镜像构建 + 端口 + 日志路径 + 滚动升级要点;cheatsheet 命令在 ../AGENTS.md「常用命令」 -->

## 部署形态

- **运行环境**:<deploy-env>
- **部署方式**:<deploy-method>
- **关联 Helm Chart**:[mns-helm/aliyun-mns-go-sdk](https://code.alibaba-inc.com/messaging/mns-helm/tree/main/aliyun-mns-go-sdk)(若适用)

## 镜像构建

```bash
<image-build-command>
```

镜像命名:`<image-naming-pattern>`(如 `mns/aliyun-mns-go-sdk:<git-tag>.<commits>-<hash>`)

镜像推送:
```bash
<image-push-command>
```

## 端口

| 端口 | 用途 |
|------|------|
| <port-1> | <port-1-purpose> |
| <port-2> | <port-2-purpose> |

## 日志路径

| 日志文件 | 内容 |
|------|------|
| `<log-file-1>` | <log-file-1-content> |
| `<log-file-2>` | <log-file-2-content> |

## JVM 参数(若适用)

```
<jvm-args>
```

## 启动 / 停止

```bash
<start-command>     # 启动
<stop-command>      # 停止
<status-command>    # 状态检查
```

## 滚动升级要点
<!-- 形态: 列表,列出升级时需注意的事项(如 graceful shutdown / 数据库迁移顺序 / Helm chart values 变更) -->

- <rolling-upgrade-note-1>
- <rolling-upgrade-note-2>

## 监控与告警

- **SLA 指标**:<sla-metric>
- **告警通道**:<alert-channel>
