# 修复前故障复现（Docker）

## 项目与标准命令

项目使用 Go 1.23 构建。本机平台为 linux/arm64。

```bash
docker build -f benzhi.Dockerfile -t go-shiftboard__005-bug:20260817 .
docker run --rm go-shiftboard__005-bug:20260817 go build ./...
docker run --rm go-shiftboard__005-bug:20260817 go test ./...
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
--- FAIL: TestDispatchReceiptIncludesReadyStatus (0.00s)
    receipt_test.go:8: receipt = "dispatched work orders: 2\n", want "dispatched work orders: 2\nstatus: ready\n"
FAIL
```

## 期望行为

晨班交接单应同时输出已派发工单数和最终的已就绪状态行。
