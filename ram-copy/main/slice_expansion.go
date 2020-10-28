package main

import "fmt"

//当前size不超过1024，按每次2倍增长，否则按当前大小的1/4增长。
// slice will inherit array's capacity and length.
//func main() {
//	data := [7]int{0, 1, 2, 3, 4, 5,}
//	fmt.Println(data, len(data), cap(data))
//	a := data[:3]
//	fmt.Println(a, len(a), cap(a))
//	a = append(a, a[:3]...) // output [0 1 2 100]
//	fmt.Println(a, len(a), cap(a))
//	a = append(a, 100)
//	fmt.Println(a, len(a), cap(a))
//
//}

func main() {
	a := []int{1} // after initialization len=1 cap=1
	test(a)       // call test to append slice, but a is [1], not [1 2]
	fmt.Println(a, len(a), cap(a))
}
func test(a []int) {
	a = append(a, 2)
	fmt.Println(a, len(a), cap(a))
}

//----output----
//[1 2] 2 2
//[1] 1 1