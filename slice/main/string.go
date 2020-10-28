package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	s := "hello, world"

	s1 := "hello, world"[:5]
	s2 := "hello, world"[7:]

	fmt.Println("len(s):", (*reflect.StringHeader)(unsafe.Pointer(&s)).Len)   // 12
	fmt.Println("len(s1):", (*reflect.StringHeader)(unsafe.Pointer(&s1)).Len) // 5
	fmt.Println("len(s2):", (*reflect.StringHeader)(unsafe.Pointer(&s2)).Len) // 5

	fmt.Printf("%v , len is: %v \n", []byte("Hello, 世界"), len([]byte("Hello, 世界")))
	fmt.Printf("%#v \n", []byte("Hello, 世界"))

	fmt.Println("\xe4\xb8\x96\xe7\x95\x8c")

	// 破坏掉第二和第三个字节，会打印�界，错误编码不会向后扩散。
	fmt.Println("\xe4\x00\x00\xe7\x95\x8c")
	// 遍历含有损坏掉UTF8字符串，第一和第二个字节依然会被单独遍历到，不过遍历到到值是损坏后到0

	for i, c := range "\xe4\x00\x00\xe7\x95\x8cabc" {
		fmt.Println(i, c)
	}
	fmt.Println("-------")
	for i, c := range []byte("世界abc") {
		fmt.Println(i, c)
	}
	fmt.Println("-----")

	ss := "\xe4\x00\x00\xe7\x95\x8c"
	for i := 0; i < len(ss); i++ {
		fmt.Printf("%d %x\n", i, ss[i])
	}

	fmt.Printf("%#v \n", []rune("世界"))
	fmt.Printf("%#v \n", string([]rune("世界")))
}

// len(s): 12
// len(s1): 5
// len(s2): 5
// [72 101 108 108 111 44 32 228 184 150 231 149 140] , len is: 13
// []byte{0x48, 0x65, 0x6c, 0x6c, 0x6f, 0x2c, 0x20, 0xe4, 0xb8, 0x96, 0xe7, 0x95, 0x8c}
// utf8码，每个中文占三位，每个英文占用一位
//  世界
// �界
// 0 65533
// 1 0
// 2 0
// 3 30028
// 6 97
// 7 98
// 8 99
// -----
// 0 228
// 1 184
// 2 150
// 3 231
// 4 149
// 5 140
// 6 97
// 7 98
// 8 99
// -----
// 0 e4
// 1 0
// 2 0
// 3 e7
// 4 95
// 5 8c
// []int32{19990, 30028}
// "世界"
//
// “一个字符串是一个不可改变的字节序列，字符串通常是用来包含人类可读的文本数据。
// 和数组不同的是，字符串的元素不可修改，是一个只读的字节数组。
// 每个字符串的长度虽然也是固定的，但是字符串的长度并不是字符串类型的一部分。”
//
// 字符串底层结构是reflect.StringHeader
// StringHeader is the runtime representation of a string.
// It cannot be used safely or portably and its representation may
// change in a later release.
// Moreover, the Data field is not sufficient to guarantee the data
// it references will not be garbage collected, so programs must keep
// a separate, correctly typed pointer to the underlying data.
//
// type StringHeader struct {
// 	Data uintptr
// 	Len  int
// }
