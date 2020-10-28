package main

import (
	"fmt"
	"./s"
)

var _ int64 = s.S()


func main() {
	fmt.Println("main")
}

//-----output-----
//init in sandbox.go
//calling s() in sandbox.go
//main
//Package initialization is done only once even if package is imported many times.

//package main
//import "fmt"
//var _ int64 = s()
//func init() {
//	fmt.Println("init in sandbox.go")
//}
//func s() int64 {
//	fmt.Println("calling s() in sandbox.go")
//	return 1
//}
//func main() {
//	fmt.Println("main")
//}
// -----output-----
//calling s() in sandbox.go
//init in sandbox.go
//main