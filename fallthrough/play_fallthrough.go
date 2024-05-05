package main

import "fmt"

func main() {
	num := 1

	switch num {
	case 1:
		fmt.Println("One")
		fallthrough // 执行此处的 fallthrough，会穿透到下一个 case 分支
	case 2:
		fmt.Println("Two")
	case 3:
		fmt.Println("Three")
	default:
		fmt.Println("Unknown")
	}
}

// output
// Franks-Mac:playGo frank$ go run type-switch/main/play_fallthrough.go
// One
// Two
