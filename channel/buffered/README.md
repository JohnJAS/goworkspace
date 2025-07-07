# 带缓冲Channel示例

这个示例展示了Go语言中带缓冲channel的用法：

1. 创建一个容量为2的缓冲channel
2. 向channel发送两个数据
3. 从channel接收数据

## 运行

```bash
go run main.go
```

## 输出示例
```
Buffered channel: 1 2
```

## 注意
- 缓冲channel允许在没有接收者的情况下发送数据，直到缓冲区满
- 当缓冲区满时，发送操作会阻塞
- 当缓冲区为空时，接收操作会阻塞
