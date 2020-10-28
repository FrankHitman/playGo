package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"siemens.com/wallbox/pkg/fastlog"
)

func doPrint(ch <-chan int, wg *sync.WaitGroup) {
	// must close channel first then can do wg.Done()
	defer wg.Done()
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Goroutine 1 run into error: ", err)
			wg.Add(1)
			go doPrint(ch, wg)
		}
	}()
	for {
		value, ok := <-ch
		if !ok {
			fmt.Println("channel 1 closed")
			break
		}
		if value == 6 {
			panic("get 6")
		}
		fmt.Println("Goroutine 1 read: ", value)
		time.Sleep(time.Second)
	}
}
func main() {
	// runtime.GOMAXPROCS(1)

	// sync with child goroutine to ensure all goroutine finished
	wg := sync.WaitGroup{}
	wg.Add(2)
	defer wg.Wait()

	// must close channel first then wg.Wait() to wait all goroutine wg.Done()
	var ch = make(chan int, 10)
	defer close(ch)

	go doPrint(ch, &wg)
	// go func() {
	//
	// 	defer wg.Done()
	// 	for {
	// 		value, ok := <-ch
	// 		if !ok {
	// 			fmt.Println("channel1 closed")
	// 			break
	// 		}
	// 		fmt.Println("Goroutine 1 read: ", value)
	//
	// 	}
	//
	// }()

	go func() {
		defer wg.Done()
		for {
			value, ok := <-ch
			if !ok {
				fmt.Println("channel 2 closed")
				break
			}
			fmt.Println("Goroutine 2 read: ", value)
			time.Sleep(time.Second)
		}
	}()

	time.Sleep(time.Second * 2)

	ch <- 3
	time.Sleep(time.Second)
	ch <- 6
	time.Sleep(time.Second)
	ch <- 9
	time.Sleep(time.Second)

	ch <- 3
	time.Sleep(time.Second)

	ch <- 6
	time.Sleep(time.Second)

	ch <- 9
	time.Sleep(time.Second)
	fmt.Println("-------")
	ch <- 3
	ch <- 6
	ch <- 9
	ch <- 3
	ch <- 6
	ch <- 9
	ch <- 3
	ch <- 6
	ch <- 9
	ch <- 3
	ch <- 6
	ch <- 9

	// Wait for interrupt signal to gracefully shutdown the server with a timeout of 3 seconds.
	quit := make(chan os.Signal)
	// kill (no param) default send syscanll.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall. SIGKILL but can"t be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	fastlog.Println("Shutdown Server ...")

}
