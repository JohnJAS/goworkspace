# Worker Pool示例

这个示例展示了Go语言中worker pool模式的用法：

1. 创建jobs和results两个channel
2. 启动3个worker goroutine
3. 发送5个job到jobs channel
4. 每个worker从jobs channel接收job并处理
5. 将结果发送到results channel
6. 主goroutine收集所有结果

## 运行

```bash
go run main.go
```

## 输出示例
```
worker 1 started job 1
worker 2 started job 2
worker 3 started job 3
worker 1 finished job 1
worker 1 started job 4
worker 2 finished job 2
worker 2 started job 5
worker 3 finished job 3
worker 1 finished job 4
worker 2 finished job 5
```

## 注意
- 使用buffered channel可以控制并发数量
- 关闭jobs channel表示所有任务已完成
- 需要等待所有结果返回
- 这种模式适合处理大量并行任务
