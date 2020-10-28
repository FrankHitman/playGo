package main

import (
	"fmt"
)

type Power struct {
	age  int
	high int
	name string
}

func main() {

	var i Power = Power{age: 10, high: 178, name: "NewMan"}

	fmt.Printf("type:%T\n", i)
	fmt.Printf("value:%v\n", i)
	fmt.Printf("value+:%+v\n", i)
	fmt.Printf("value#:%#v\n", i)

	fmt.Println("========interface========")
	var inter interface{} = i
	fmt.Printf("type:%T\n", inter)
	fmt.Printf("interface+:%+v\n", inter)
	fmt.Printf("interface#:%#v\n", inter)
}

// output
// type:main.Power										%T打印类型
// value:{10 178 NewMan}								%v打印struct中的value
// value+:{age:10 high:178 name:NewMan} 				%+v打印struct详情key/value键值对
// value#:main.Power{age:10, high:178, name:"NewMan"}	%#v打印带类型的struct详情
// ========interface========
// type:main.Power										struct赋值给interface{}类型后，保持自身的类型
// interface+:{age:10 high:178 name:NewMan}
// interface#:main.Power{age:10, high:178, name:"NewMan"}
