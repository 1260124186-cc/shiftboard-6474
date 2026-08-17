# 修复前故障复现（Docker）

## 项目与标准命令

项目使用 Go 1.23 构建。本机平台为 linux/arm64。

```bash
docker build -f benzhi.Dockerfile -t go-shiftboard__004-bug:20260817 .
docker run --rm go-shiftboard__004-bug:20260817 go build ./...
docker run --rm go-shiftboard__004-bug:20260817 go test ./...
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
--- FAIL: TestZeroValueBoardCanImportAndListWorkOrders (0.00s)
    zero_value_test.go:15: zero-value board panicked: runtime error: invalid memory address or nil pointer dereference
FAIL
```

## 期望行为

即使排班服务按零值方式创建，首次导入工单和读取工单也应安全完成，不应触发 panic。
