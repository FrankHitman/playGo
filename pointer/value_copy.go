package main

import "fmt"

func double(x *int) {
	*x += *x
	x = nil // the line is just for explanation purpose, copy of pointer.
	// the modification of the copy of the passed pointer argument itself
	// still can't be reflected on the passed pointer argument.
}

func main() {
	var a = 3
	double(&a)
	fmt.Println(a) // 6
	p := &a
	double(p)
	fmt.Println(a, p == nil) // 12 false
}

// output
// Franks-Mac:playGo frank$ go run pointer/value_copy.go
// 6
// 12 false
// refer to https://go101.org/article/pointer.html
