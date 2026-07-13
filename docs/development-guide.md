# Aliyun Mns Go Sdk — 开发指南

<!-- 范围: 多步开发工作流 + 本地冒烟 + 故障排查 + FAQ;cheatsheet 命令在 ../AGENTS.md「常用命令」段 -->

## 开发工作流
<!-- 形态: 编号步骤 + 命令 + 上下文;每步 1-3 行 -->

1. **环境准备**:<env-setup-detail>
2. **克隆代码**:`git clone https://code.alibaba-inc.com/messaging/aliyun-mns-go-sdk` && `git checkout master`
3. **编译项目**:`<build-command>`
4. **配置修改**:<config-modify-detail>
5. **本地运行**:`<local-run-command>`
6. **验证结果**:<verify-result-detail>

## 本地冒烟测试
<!-- 形态: 编号步骤 + 命令 + 前置 + 预期输出 + 失败标志 -->

本项目的本地冒烟测试旨在验证代码改动后能否正常编译、启动,并连接<smoke-test-target>,验证程序的完整可用性。

> **注意**:<smoke-test-prerequisite>

### 前置准备

<smoke-test-prep>

### 验证步骤

1. **<smoke-step-1-title>**:
   ```bash
   <smoke-step-1-command>
   ```
   <smoke-step-1-expected>

2. **<smoke-step-2-title>**:
   ```bash
   <smoke-step-2-command>
   ```
   <smoke-step-2-expected>

3. **观察日志**:
   <smoke-log-observe>

   **启动成功标志**:
   - <success-marker-1>
   - <success-marker-2>

   **失败标志**:
   - <failure-marker-1>
   - <failure-marker-2>

### 关键配置说明

- <key-config-1>
- <key-config-2>

## 常见问题

### <faq-1-title>
**原因**:<faq-1-cause>
**解决**:<faq-1-solution>

### <faq-2-title>
**原因**:<faq-2-cause>
**解决**:<faq-2-solution>

### <faq-3-title>
**原因**:<faq-3-cause>
**解决**:<faq-3-solution>

## 调试技巧
<!-- 形态: 1-3 段或列表,列出常用调试方法(如远程 debug 端口、日志 level 调高、覆盖率报告) -->

<debug-techniques>
