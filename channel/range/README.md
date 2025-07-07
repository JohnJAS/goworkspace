# Channel关闭和Range遍历示例

这个示例展示了Go语言中channel关闭和range遍历的用法：

1. 创建一个channel
2. 启动goroutine发送数据并关闭channel
3. 使用range遍历channel中的数据
4. 当channel关闭且数据读取完毕后，range循环结束

## 运行

```bash
go run main.go
```

## 输出示例
```
Range over channel: 1 2 3
```

## 注意
- 只有发送方可以关闭channel
- 关闭channel后不能再发送数据
- range会一直从channel读取数据直到channel关闭
- 如果不关闭channel，range会导致死锁
