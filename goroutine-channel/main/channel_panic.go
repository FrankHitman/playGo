package main

import (
	"fmt"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(10)
	int_chan := make(chan int, 1)    // 带缓存的 channel
	string_chan := make(chan string) // 不带缓存的 channel，阻塞的
	// int_chan <- 2
	int_chan <- 1
	string_chan <- "hello"
	select {
	case value := <-int_chan:
		fmt.Println(value)
	case value := <-string_chan:
		fmt.Println(value)
		// default:
		//	fmt.Println("bye")

	}
}

// output
// fatal error: all goroutines are asleep - deadlock!
//
// goroutine 1 [chan send]:
// main.main()
//        /Users/frank/go/src/github.com/FrankHitman/playGo/goroutine-channel/main/channel_panic.go:14 +0x70

// explain
// It will panic at ```string_chan <- "hello"```, because this channel with no buffer, when put something in it ,
// it will check the consumer is already or not, when consumer is not ready, it will panic.
//
// ```int_chan <- 1``` will panic if add ```int_chan <- 2``` before it.
