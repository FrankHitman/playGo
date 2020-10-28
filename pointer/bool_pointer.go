package main

import "fmt"

func main() {
	var p *bool
	a := false
	p = &a
	where := fmt.Sprintf("is_installed = %d", p)

	fmt.Println(where)
	fmt.Printf("is_installed = %p", p)
	fmt.Println()
	fmt.Printf("is_installed = %v", *p)
	fmt.Println()
	fmt.Printf("is_installed = %v", p)
	fmt.Println()
	fmt.Printf("a is %v", a)

}
// output
// is_installed = 824634318848  %d打印10进制，打印指针的值
// is_installed = 0xc000092000  %p打印指针的值，此16进制转化为10进制即为824634318848
// is_installed = false			%v-*p打印指针指向地址的内存中存储的值
// is_installed = 0xc000092000  %v-p打印指针的值
// a is false
// 再一次执行，地址发生变化
// is_installed = 824633819226
// is_installed = 0xc00001805a
// is_installed = false
// is_installed = 0xc00001805a
// a is false