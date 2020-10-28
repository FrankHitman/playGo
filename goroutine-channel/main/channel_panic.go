package main

import (
	"fmt"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(10)
	int_chan := make(chan int, 1)
	string_chan := make(chan string)
	//int_chan <- 2
	int_chan <- 1
	string_chan <- "hello"
	select {
	case value := <-int_chan:
		fmt.Println(value)
	case value := <-string_chan:
		fmt.Println(value)
		//default:
		//	fmt.Println("bye")

	}
}
