package main

import "fmt"

type innerS struct {
	in1 int
	in2 int
}

type outerS struct {
	b    int
	c    float32
	int  // anonymous  field
	innerS //anonymous embedded field    匿名字段和内嵌结构体
}

func main() {
	outer := new(outerS)
	outer.b = 6
	outer.c = 7.5
	outer.int = 60
	outer.in1 = 5
	outer.in2 = 10

	fmt.Printf("outer.b is: %d\n", outer.b)
	fmt.Printf("outer.c is: %f\n", outer.c)
	fmt.Printf("outer.int is: %d\n", outer.int)
	fmt.Printf("outer.in1 is: %d\n", outer.in1)
	fmt.Printf("outer.in2 is: %d\n", outer.in2)
	fmt.Println("-------------")
	fmt.Println("outer.innerS.in1", outer.innerS.in1)
	fmt.Println("outer is ", outer)
	fmt.Println("outer.innerS is ", outer.innerS)
	// 使用结构体字面量
	outer2 := outerS{6, 7.5, 60, innerS{5, 10}}
	fmt.Println("outer2 is:", outer2)
}

// -----output-----
//outer.b is: 6
//outer.c is: 7.500000
//outer.int is: 60
//outer.in1 is: 5
//outer.in2 is: 10
//-------------
//outer.innerS.in1 5
//outer is  &{6 7.5 60 {5 10}}
//outer.innerS is  {5 10}
//outer2 is: {6 7.5 60 {5 10}}

//通过类型 outer.int 的名字来获取存储在匿名字段中的数据，于是可以得出一个结论：在一个结构体中对于每一种数据类型只能有一个匿名字段。

//当两个字段拥有相同的名字（可能是继承来的名字）时该怎么办呢？
//外层名字会覆盖内层名字（但是两者的内存空间都保留），这提供了一种重载字段或方法的方式；
//如果相同的名字在同一级别出现了两次，如果这个名字被程序使用了，将会引发一个错误（不使用没关系）。没有办法来解决这种问题引起的二义性，必须由程序员自己修正。