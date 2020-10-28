package main

import (
	"fmt"
	"time"
)

func main() {
	var result uint64
	for i := 0; i <= 10; i++ {
		start := time.Now()
		result = fibonacci(i)
		end := time.Now()
		delta := end.Sub(start)
		fmt.Printf("fibonacci(%d) is: %d\n", i, result)
		fmt.Println("time delta is ", delta)
		fmt.Println()
	}
}

const LIM = 41
var fibs [LIM]uint64

func fibonacci(n int) (res uint64) {
	if fibs[n] != 0{
		return fibs[n]
	}

	if n <= 1 {
		res = 1
	} else {
		res = fibonacci(n-1) + fibonacci(n-2)
	}
	fibs[n] = res // cache in ram will run faster
	return
}

