# 修复前故障复现（Docker）

## 项目与标准命令

项目使用 Go 1.23 构建。本机平台为 linux/arm64。

```bash
docker build -f benzhi.Dockerfile -t go-shiftboard__001-bug:20260817 .
docker run --rm go-shiftboard__001-bug:20260817 go build ./...
docker run --rm go-shiftboard__001-bug:20260817 go test ./...
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
--- FAIL: TestWorkOrderCloneOwnsMutableFields (0.00s)
    workorder_test.go:18: tags were shared
--- FAIL: TestBoardOwnsImportedMutableData (0.00s)
    ownership_test.go:34: assign after caller mutations: work order WO-ownership is missing a required zone tag
FAIL
```

## 期望行为

导入后的排班面板应保留自己的工单和区域规则数据；调用方随后改变原始标签不应改变面板中的可分配工单。
