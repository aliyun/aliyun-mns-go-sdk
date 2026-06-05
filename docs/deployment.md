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
