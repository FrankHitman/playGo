package main

import "fmt"

type I interface {
	m(int) bool
}

type T string

func (t T) m(n int) bool {
	return len(t) > n
}

func main() {
	var i I = T("gopher")
	fmt.Println(i.m(5))                           // true
	fmt.Println(I.m(i, 5))                        // true
	fmt.Println(interface{ m(int) bool }.m(i, 5)) // true

	// The following lines compile okay,
	// but will panic at run time.
	I(nil).m(5)
	I.m(nil, 5)
	interface{ m(int) bool }.m(nil, 5)
}

// output
// Franks-Mac:playGo frank$ go build interface/main/implicit_function.go
// Franks-Mac:playGo frank$ ./implicit_function
// true
// true
// true
// panic: runtime error: invalid memory address or nil pointer dereference
// [signal SIGSEGV: segmentation violation code=0x1 addr=0x0 pc=0x2699b0c]
//
// goroutine 1 [running]:
// main.main()
//        /Users/frank/go/src/github.com/FrankHitman/playGo/interface/main/implicit_function.go:23 +0x16c
// Franks-Mac:playGo frank$

