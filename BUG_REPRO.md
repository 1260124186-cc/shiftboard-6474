# 修复前故障复现（Docker）

## 项目与标准命令

项目使用 Go 1.23 构建。本机平台为 linux/arm64。

```bash
docker build -f benzhi.Dockerfile -t go-shiftboard__002-bug:20260817 .
docker run --rm go-shiftboard__002-bug:20260817 go build ./...
docker run --rm go-shiftboard__002-bug:20260817 go test ./...
```

## 环境构建与编译

镜像已在 linux/arm64 成功构建，容器内 `go version` 输出为 `go version go1.23.12 linux/arm64`，容器内 `go build ./...` 成功。

## 故障触发步骤

在容器中执行标准测试命令：

```bash
go test ./...
```

## 实际错误输出

```text
--- FAIL: TestMissingZonePolicyKeepsItsClassification (0.00s)
    policy_error_test.go:22: assignment error did not preserve the policy classification: build shift assignments: no policy for zone south
FAIL
```

## 期望行为

区域没有排班规则时，调用端应能稳定识别该业务错误，报表也应呈现对应的业务状态，而不是退化为普通失败。
