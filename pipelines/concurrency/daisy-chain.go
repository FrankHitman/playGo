package main

import "fmt"

func f(left, right chan int) {
	a := <-right
	// fmt.Println("right is ", a)
	left <- 1 + a
}

func main() {
	const n = 10000
	leftmost := make(chan int)
	right := leftmost
	left := leftmost
	for i := 0; i < n; i++ {
		right = make(chan int)
		go f(left, right)
		// fmt.Println("left ",left)
		// fmt.Println("right ",right)
		// right = left
		left = right
	}
	go func(c chan int) { c <- 1 }(right)
	fmt.Println(<-leftmost)
}
