//package main
//
//import "fmt"
//
//func a() {
//	fmt.Println("a")
//}
//
//func e() {
//	panic("error")
//}
//
//func b() {
//	fmt.Println("b")
//}
//
//func main() {
//	defer a()
//	defer e()
//	defer b()
//
//}

//----output----
//b
//a
//panic: error
//
//goroutine 1 [running]:
//main.e()
//	/home/sdy/go/src/github.com/FrankHitman/playGo/cmd/play-defer.go:10 +0x39
//main.main()
//	/home/sdy/go/src/github.com/FrankHitman/playGo/cmd/play-defer.go:24 +0x76

//defer的特点就是LIFO，即后进先出，所以如果在同一个函数下多个defer的话，会逆序执行。
//那调用者的defer会被执行吗？
//panic仅保证当前goroutine下的defer都会被调到，但不保证其他协程的defer也会调到。如果是在同一goroutine下的调用者的defer，那么可以一路回溯回去执行；
// 但如果是不同goroutine，那就不做保证了。

package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	defer fmt.Println("defer main") // will this be printed when panic?
	var user = os.Getenv("USER_")
	go func() {
		defer fmt.Println("defer caller")
		func() {
			defer func() {
				fmt.Println("defer here")
			}()

			if user == "" {
				panic("should set user env.")
			}
		}()
	}()

	time.Sleep(1 * time.Second)
	fmt.Println("get result ", )
}
