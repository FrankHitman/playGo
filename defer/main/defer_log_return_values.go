package main

import (
	"io"
	"log"
)

//使用 defer 语句来记录函数的参数与返回值

func func1(s string) (n int, err error) {
	defer func() {
		log.Printf("func1(%q) = %d, %v", s, n, err) // get return values
	}()
	return 7, io.EOF
}

func main() {
	func1("Go")
}