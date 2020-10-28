package main

import (
	"fmt"
	"time"
)

func myroutine(routineid string) {
	fmt.Println("Entry routine:", routineid)
	time.Sleep(5 * time.Second)
	fmt.Println("Exit  routine:", routineid)
}

func main() {
	fmt.Println("Entry routine: main")

	// Launch goroutine 1
	go myroutine("myroutine1")

	// Launch goroutine 2
	go func(routineid string) {
		fmt.Println("Entry routine:", routineid)
		time.Sleep(5 * time.Second)
		fmt.Println("Exit routine:", routineid)
	}("myroutine2")

	// wait main routine
	fmt.Scanln()
	fmt.Println("Exit  routine: main")
}
