# 修复前故障复现（Docker）

## 项目与标准命令

项目使用 Go 1.23 构建。本机平台为 linux/arm64。

```bash
docker build -f benzhi.Dockerfile -t go-shiftboard__003-bug:20260817 .
docker run --rm go-shiftboard__003-bug:20260817 go build ./...
docker run --rm go-shiftboard__003-bug:20260817 go test ./...
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
--- FAIL: TestCanceledImportDoesNotPartiallyPersistWorkOrders (0.00s)
    import_context_test.go:23: import error did not preserve cancellation: import work orders: import interrupted: import canceled after 1 work orders
FAIL
```

## 期望行为

已经取消的批量导入应不改变排班面板中的工单，并让调用方保留可识别的取消结果。
