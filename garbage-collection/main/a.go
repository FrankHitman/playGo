package main

import (
	"fmt"
	"runtime"
	"time"
)

const (
	windowSize = 200000
	msgCount   = 1000000
)

type (
	message []byte
	buffer  [windowSize]message
)

var worst time.Duration

func mkMessage(n int) message {
	m := make(message, 1024)
	for i := range m {
		m[i] = byte(n)
	}
	return m
}

func pushMsg(b *buffer, highID int) {
	start := time.Now()
	m := mkMessage(highID)
	(*b)[highID%windowSize] = m
	elapsed := time.Since(start)
	if elapsed > worst {
		worst = elapsed
	}
}

func main() {
	var b buffer
	for i := 0; i < msgCount; i++ {
		pushMsg(&b, i)
	}

	var m runtime.MemStats
	runtime.ReadMemStats(&m) // stop the world
	fmt.Printf("%d Kb\n", m.Alloc / 1024)
	//runtime.SetFinalizer(obj, func(obj *typeObj))

	fmt.Println("Worst push time: ", worst)
}