package main

import (
	"fmt"
)

const (
	DEBUG = 1 << iota
	INFO
	WARN
	ERROR
	FATAL
)

func main() {
	var a uint8
	a = 4
	var b int8
	b = 4
	fmt.Println(DEBUG)
	fmt.Println(INFO)
	fmt.Println(WARN)
	fmt.Println(ERROR)

	//位左移 <<：
	//用法：bitP << n。
	//bitP 的位向左移动 n 位，右侧空白部分使用 0 填充；如果 bitP 等于 2，则结果是 2 的相应倍数，即 2 的 n 次方。例如：
	fmt.Println(1 << 3)

	//位右移 >>：
	//用法：bitP >> n。
	//bitP 的位向右移动 n 位，左侧空白部分使用 0 填充；如果 bitP 等于 2，则结果是当前值除以 2 的 n 次方。


	//按位补足 ^：
	//该运算符与异或运算符一同使用，即 m^x，对于无符号 x 使用“其他为0位全部设置为 1”，对于有符号 x 时使用 m=-1
	fmt.Println(^100)
	fmt.Println(^a)
	fmt.Println(^b)

	a <<= 2
	fmt.Println(a)

	var c uint8
	c= 4
	c ^= a & 0xff
	fmt.Println(0xff)
	fmt.Println(1<<8-1)
	fmt.Println(c)
}

// -----output------
//1
//2
//4
//8
//8
//-101 = -01^100 = -001^100 = -101 byte
//251 = ^00000100 = 11111011 = 251  uint8 transform to byte
//-5 = -01^100 = -101 = -5 int8
//16 = 4<<2 = 100<<2 = 10000
//255 = 0xff = 11111111
//255 = 000000001<<8 -1 = 100000000-1 = 011111111
//20 = 00000100^(00010000&11111111) = 00000100^00010000 = 10100 = 20
