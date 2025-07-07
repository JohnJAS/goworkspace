# 生产者消费者模式最佳实践

这个示例展示了Go语言中生产者消费者模式的channel创建和关闭规范：

## 最佳实践

1. **Channel创建**
   - 由主goroutine创建channel
   - 根据需求选择是否使用buffered channel
   - 明确channel方向（<-chan/chan<-）

2. **Channel关闭**
   - 由生产者负责关闭channel
   - 关闭channel表示数据流结束
   - 消费者通过range自动检测channel关闭

3. **同步机制**
   - 使用额外的channel（如doneCh）来同步消费者完成状态
   - 主goroutine负责等待所有消费者完成

## 运行

```bash
go run main.go
```

## 输出示例
```
Producing: 1
Consumer 1 received: 1
Producing: 2
Consumer 2 received: 2
Producing: 3
Consumer 1 received: 3
Producing: 4
Consumer 2 received: 4
Producing: 5
Producer finished
Consumer 1 received: 5
All consumers finished
```

## 注意
- 生产者关闭channel后，消费者会自动退出
- 使用done channel确保所有消费者完成
- 这种模式适合解耦生产者和消费者
