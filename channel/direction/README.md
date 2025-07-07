# Channel方向控制示例

这个示例展示了Go语言中channel方向控制的用法：

1. 创建两个channel
2. 定义pingPong函数，指定channel方向
3. 主goroutine发送数据
4. pingPong函数接收并处理数据
5. 主goroutine接收处理结果

## 运行

```bash
go run main.go
```

## 输出示例
```
Directional channel: ping pong
```

## 注意
- 可以在函数参数中指定channel方向
- <-chan表示只读channel
- chan<-表示只写channel
- 这有助于提高代码可读性和安全性
