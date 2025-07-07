# Select语句示例

这个示例展示了Go语言中select语句的用法：

1. 创建两个channel
2. 启动两个goroutine分别向channel发送数据
3. 使用select语句监听多个channel
4. 使用time.After实现超时控制

## 运行

```bash
go run main.go
```

## 输出示例
```
Select timeout
```

## 注意
- select语句用于同时监听多个channel操作
- 当多个case同时就绪时，会随机选择一个执行
- time.After可以用于实现超时控制
- 如果没有default分支，select会阻塞直到某个case就绪
