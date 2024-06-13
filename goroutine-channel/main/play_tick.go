package main

import (
	"fmt"
	"time"
)

func main() {
	c := time.Tick(5 * time.Second)
	for next := range c {
		// fmt.Printf("%v %s\n", next, statusUpdate())
		fmt.Printf("%v \n", next)

	}
}

// output
// 2024-06-13 12:18:24.601677 +0800 CST m=+5.000246145
// 2024-06-13 12:18:29.602131 +0800 CST m=+10.000659382
// 2024-06-13 12:18:34.601964 +0800 CST m=+15.000451611
