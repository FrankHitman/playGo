package main

import (
	"github.com/petermattis/goid"
	"log"
	"math/rand"
	"strconv"
	"time"
)

// multiple sender and multiple receiver
func main() {
	rand.Seed(time.Now().UnixNano())

	const Max = 100000
	const NumReceivers = 10
	const NumSenders = 10

	dataCh := make(chan int, 100)
	stopCh := make(chan struct{})

	// It must be a buffered channel.
	// toStop := make(chan string)
	toStop := make(chan string, NumReceivers + NumSenders)

	var stoppedBy string

	// moderator
	go func() {
		stoppedBy = <-toStop
		log.Println("stop by: ",stoppedBy)
		close(stopCh)
	}()

	// senders
	for i := 0; i < NumSenders; i++ {
		go func(id string) {
			for {
				value := rand.Intn(Max)
				if value == 0 {
					toStop <- "sender#" +id
					log.Println("sender will to stop: ", goid.Get())
					return
					// select {
					// case toStop <- "sender#" + id:
					// 	log.Println("sender will to stop: ", goid.Get())
					// default:
					// }
					// return
				}

				select {
				case <- stopCh:
					log.Println("sender get stop channel, goroutine id: ",goid.Get())
					return
				case dataCh <- value:
				}
			}
		}(strconv.Itoa(i))
	}

	// receivers
	for i := 0; i < NumReceivers; i++ {
		go func(id string) {
			for {
				select {
				case <- stopCh:
					log.Println("receiver get stop channel, goroutine id: ",goid.Get())
					return
				case value := <-dataCh:
					if value == Max-1 {
						toStop <- "receiver#" + id
						log.Println("receiver will to stop: ", goid.Get())
						return
						// select {
						// case toStop <- "receiver#" + id:
						// 	log.Println("receiver will to stop: ", goid.Get())
						// default:
						// }
						// return
					}

					log.Println(value)
				}
			}
		}(strconv.Itoa(i))
	}

	select {
	case <- time.After(time.Hour):
	}

}

// 代码里 toStop 就是中间人的角色，使用它来接收 senders 和 receivers 发送过来的关闭 dataCh 请求。
// 这里将 toStop 声明成了一个 缓冲型的 channel。假设 toStop 声明的是一个非缓冲型的 channel，那么第一个发送的关闭 dataCh 请求可能会丢失。
// 因为无论是 sender 还是 receiver 都是通过 select 语句来发送请求，如果中间人所在的 goroutine 没有准备好，那 select 语句就不会选中，直接走 default 选项，什么也不做。
// 这样，第一个关闭 dataCh 的请求就会丢失。
// 链接：https://juejin.im/post/5d350e70f265da1b897b0cbe