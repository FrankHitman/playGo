# channel

## channel、生产者、消费者之间的关系

1. 第 n 个 send 一定 happened before 第 n 个 receive finished，无论是缓冲型还是非缓冲型的 channel。
2. 对于容量为 m 的缓冲型 channel，第 n 个 receive 一定 happened before 第 n+m 个 send finished。
3. 对于非缓冲型的 channel，第 n 个 receive 一定 happened before 第 n 个 send finished。
4. channel close 一定 happened before receiver 得到通知。

## Channel 可能会引发 goroutine 泄漏。

泄漏的原因是 goroutine 操作 channel 后，处于发送或接收阻塞状态，而 channel 处于满或空的状态，一直得不到改变。
同时，垃圾回收器也不会回收此类资源，进而导致 gouroutine 会一直处于等待队列中，不见天日。

## Channel 发送和接收元素的本质是什么？

Remember all transfer of value on the go channels happens with the copy of value.

就是说 channel 的发送和接收操作本质上都是 “值的拷贝”，无论是从 sender goroutine 的栈到 chan buf，
还是从 chan buf 到 receiver goroutine，或者是直接从 sender goroutine 到 receiver goroutine。

## 操作channel的结果

| 操作      | nil channel | closed channel | not nil, not closed channel                           |
|---------|-------------|----------------|-------------------------------------------------------|
| close   | panic       | panic          | 正常关闭                                                  |
| 读 <- ch | 阻塞          | 读到对应类型的零值      | 阻塞或正常读取数据。缓冲型 channel 为空或非缓冲型 channel 没有等待发送者时会阻塞     |
| 写 ch <- | 阻塞          | panic          | 阻塞或正常写入数据。非缓冲型 channel 没有等待接收者或缓冲型 channel buf 满时会被阻塞 |

## 如何优雅的关闭channel

don't close a channel from the receiver side and don't close a channel if the channel has multiple concurrent senders.

don't close (or send values to) closed channels.

根据 sender 和 receiver 的个数，分下面几种情况：

- 一个 sender，一个 receiver
- 一个 sender， M 个 receiver
- N 个 sender，一个 receiver
- N 个 sender， M 个 receiver

对于 1，2，只有一个 sender 的情况就不用说了，直接从 sender 端关闭就好了，没有问题。重点关注第 3，4 种情况。

第 3 种情形下，优雅关闭 channel
的方法是：```the only receiver says "please stop sending more" by closing an additional signal channel。```
解决方案就是增加一个传递关闭信号的 channel (stopCh)，receiver 通过信号 channel 下达关闭数据 channel 指令。
senders 监听到关闭信号后，停止发送数据。 stopCh 的 发送方是 1个 receiver，接收方是 M个 sender。

最后一种情况，优雅关闭 channel 的方法是：any one of them says "let's end the game" by notifying a moderator to close an
additional signal channel。
和第 3 种情况不同，这里有 M 个 receiver，如果直接还是采取第 3 种解决方案，由 receiver 直接关闭 stopCh 的话，就会重复关闭一个
channel，导致 panic。
因此需要增加一个中间人，M 个 receiver 都向它发送 "关闭 dataCh 的请求”，中间人收到第一个请求后，就会直接通过 stopCh 向
senders
下达 "关闭 dataCh" 的指令，发送完中间人就关闭 stopCh （通过关闭 stopCh，这时就不会发生重复关闭的情况，因为 stopCh
的发送方只有中间人一个）。
另外，这里的 N 个 sender 也可以向中间人发送关闭 dataCh 的请求。

## 关闭的channel仍能读出数据

从一个有缓冲的 channel 里读数据，当 channel 被关闭，依然能读出有效值。只有当返回的 ok 为 false 时，读出的数据才是无效的。

## channel的应用

- 停止信号
- 超时任务与定时任务

```go
select {
	case <-time.After(100 * time.Millisecond):
	case <-s.stopc:
		return false
}

func worker() {
	ticker := time.Tick(1 * time.Second)
	for {
		select {
		case <- ticker:
			// 执行定时任务
			fmt.Println("执行 1s 定时任务")
		}
	}
}

```

- 解耦生产方和消费方

```go
func main() {
	taskCh := make(chan int, 100)
	go worker(taskCh)

    // 塞任务
	for i := 0; i < 10; i++ {
		taskCh <- i
	}

    // 等待 1 小时 
	select {
	case <-time.After(time.Hour):
	}
}

func worker(taskCh <-chan int) {
	const N = 5
	// 启动 5 个工作协程
	for i := 0; i < N; i++ {
		go func(id int) {
			for {
				task := <- taskCh
				fmt.Printf("finish task: %d by worker %d\n", task, id)
				time.Sleep(time.Second)
			}
		}(i)
	}
}
```

- 控制并发数

```go
var limit = make(chan int, 3)

func main() {
    // …………
    for _, w := range work {
        go func() {
            limit <- 1
            w()
            <-limit
        }()
    }
    // …………
}

```

构建一个缓冲型的 channel，容量为 3。接着遍历任务列表，每个任务启动一个 goroutine 去完成。
真正执行任务，访问第三方的动作在 w() 中完成，在执行 w() 之前，先要从 limit 中拿“许可证”，拿到许可证之后，才能执行 w()，
并且在执行完任务，要将“许可证”归还。这样就可以控制同时运行的 goroutine 数。
这里，limit <- 1 放在 func 内部而不是外部，书籍作者柴大在读者群里的解释是：

如果在外层，就是控制系统 goroutine 的数量，可能会阻塞 for 循环，影响业务逻辑。

limit 其实和逻辑无关，只是性能调优，放在内层和外层的语义不太一样。

还有一点要注意的是，如果 w() 发生 panic，那“许可证”可能就还不回去了，因此需要使用 defer 来保证。

## references

参考链接：https://juejin.im/post/5d350e70f265da1b897b0cbe
