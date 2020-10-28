package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
)

func main() {
	// reflect.SliceHeader{}
	var (
		a []int               // nil切片，和nil相等，用来表示一个不存在的切片
		b = []int{}           // 空切片，和nil不相等
		c = []int{1, 2, 3}    // 有3个元素的切片，len=cap=3
		f = c[:0]             // 有0个元素的切片，len=0，cap=3
		h = make([]int, 2, 3) // 有2个元素的切片，len=2， cap=3
	)
	fmt.Println(a, b, c, f, h)

}

// 切片是一种简化版的动态数组，因为动态数组的长度不固定，切片的长度自然也就不能是类型的组成部分。
//
//	reflect.SliceHeader{} 切片底层结构
// SliceHeader is the runtime representation of a slice.
// It cannot be used safely or portably and its representation may
// change in a later release.
// Moreover, the Data field is not sufficient to guarantee the data
// it references will not be garbage collected, so programs must keep
// a separate, correctly typed pointer to the underlying data.
// type SliceHeader struct {
// 	Data uintptr
// 	Len  int
// 	Cap  int
// }
// “切片可以和nil进行比较，只有当切片底层数据指针为空时切片本身为nil，这时候切片的长度和容量信息将是无效的。
// 如果有切片的底层数据指针为空，但是长度和容量不为0的情况，那么说明切片本身已经被损坏了
// （比如直接通过reflect.SliceHeader或unsafe包对切片作了不正确的修改）。”
//
// “在对切片本身赋值或参数传递时，和数组指针的操作方式类似，只是复制切片头信息（reflect.SliceHeader），
// 并不会复制底层的数据。对于类型，和数组的最大不同是，切片的类型和长度信息无关，
// 只要是相同类型元素构成的切片均对应相同的切片类型。”
//
// “不过要注意的是，在容量不足的情况下，切片对append的操作会导致重新分配内存，可能导致巨大的内存分配和复制数据代价。
// 即使容量足够，依然需要用append函数的返回值来更新切片本身，因为新切片的长度已经发生了变化。
//
// “切片高效操作的要点是要降低内存分配的次数，尽量保证append操作不会超出cap的容量，降低触发内存分配的次数和每次分配内存大小。”
//
// 内存泄漏问题：
// “切片操作并不会复制底层的数据。底层的数组会被保存在内存中，直到它不再被引用。
// 但是有时候可能会因为一个小的内存引用而导致底层整个数组处于被使用的状态，这会延迟自动内存回收器对底层数组的回收。”

func FindPhoneNumber(filename string) []byte {
	b, _ := ioutil.ReadFile(filename)
	return regexp.MustCompile("[0-9]+").Find(b)
}
//  |
//  v
func FindPhoneNumber2(filename string) []byte {
	b, _ := ioutil.ReadFile(filename)
	b = regexp.MustCompile("[0-9]+").Find(b)
	return append([]byte{}, b...)
}

// “假设切片里存放的是指针对象，那么下面删除末尾的元素后，被删除的元素依然被切片底层数组引用，从而导致不能及时被自动垃圾回收器回收（这要依赖回收器的实现方式）：”
func removeLast() {
	var i1 int = 1
	var i2 int = 2
	var a = []*int{&i1, &i2}
	a = a[:len(a)-1] // 被删除的最后一个元素依然被引用, 可能导致GC操作被阻碍
}
//  |
//  v
func removeLast2() {
	var i1 int = 1
	var i2 int = 2
	var a = []*int{&i1, &i2}
	a[len(a)-1] = nil // GC回收最后一个元素内存
	a = a[:len(a)-1]  // 从切片删除最后一个元素
}
