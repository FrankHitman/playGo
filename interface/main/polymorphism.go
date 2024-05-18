package main

import "fmt"

type Filter interface {
	About() string
	Process([]int) []int
}

// UniqueFilter is used to remove duplicate numbers.
type UniqueFilter struct{}

func (UniqueFilter) About() string {
	return "remove duplicate numbers"
}
func (UniqueFilter) Process(inputs []int) []int {
	outs := make([]int, 0, len(inputs))
	pusheds := make(map[int]bool) // hash 表用于判断是否重复
	for _, n := range inputs {
		if !pusheds[n] {
			pusheds[n] = true
			outs = append(outs, n)
		}
	}
	return outs
}

// MultipleFilter is used to keep only
// the numbers which are multiples of
// the MultipleFilter as an int value. 保留是 multiplefileter 数的倍数
type MultipleFilter int

func (mf MultipleFilter) About() string {
	return fmt.Sprintf("keep multiples of %v", mf)
}
func (mf MultipleFilter) Process(inputs []int) []int {
	var outs = make([]int, 0, len(inputs))
	for _, n := range inputs {
		if n%int(mf) == 0 {
			outs = append(outs, n)
		}
	}
	return outs
}

// With the help of polymorphism, only one
// "filterAndPrint" function is needed.
// 多态性使得不必为每种筛选器类型编写一个 filterAndPrint 函数。
func filterAndPrint(fltr Filter, unfiltered []int) []int {
	// Calling the methods of "fltr" will call the
	// methods of the value boxed in "fltr" actually.
	filtered := fltr.Process(unfiltered)
	fmt.Println(fltr.About()+":\n\t", filtered)
	return filtered
}

func main() {
	numbers := []int{12, 7, 21, 12, 12, 26, 25, 21, 30}
	fmt.Println("before filtering:\n\t", numbers)

	// Three non-interface values are boxed into
	// three Filter interface slice element values.
	filters := []Filter{
		UniqueFilter{},
		MultipleFilter(2), // 保留是 2 的倍数的数
		MultipleFilter(3), // 保留是 3 的倍数的数
	}

	// Each slice element will be assigned to the
	// local variable "fltr" (of interface type
	// Filter) one by one. The value boxed in each
	// element will also be copied into "fltr".
	for _, fltr := range filters {
		numbers = filterAndPrint(fltr, numbers)
	}
}

// output
// Franks-Mac:playGo frank$ go run interface/main/polymorphism.go
// before filtering:
//         [12 7 21 12 12 26 25 21 30]
// remove duplicate numbers:
//         [12 7 21 26 25 30]
// keep multiples of 2:
//         [12 26 30]
// keep multiples of 3:
//         [12 30]

// 多态性的另外一个优点是：
// 标准库声明公有的接口类型，声明接受这些类型作为参数的公有方法
// 用户侧开发者可以实现标准库中的接口类型，然后使用标准库的公有方法调用这个接口类型，
// 不必修改标准库中源码，进而扩展标准库的功能。

// refer to http://go101.org/article/interface.html
