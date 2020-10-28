package main

import (
	"fmt"
	"reflect"
)

func main() {
	var value []string
	dest := reflect.Indirect(reflect.ValueOf(&value))
	fmt.Println(dest.Kind() == reflect.Slice)

	dest3 := reflect.Indirect(reflect.ValueOf(value))
	fmt.Println(dest3.Kind() == reflect.Slice)

	value2 := make([]string, 0)
	dest2 := reflect.Indirect(reflect.ValueOf(value2))
	fmt.Println(dest2.Kind() == reflect.Slice)

	elem := reflect.New(dest.Type().Elem()).Interface()
	dest.Set(reflect.Append(dest, reflect.ValueOf(elem).Elem()))
	fmt.Println(dest)

	elem3 := reflect.New(dest3.Type().Elem()).Interface()
	dest3.Set(reflect.Append(dest3, reflect.ValueOf(elem3).Elem()))
	fmt.Println(dest3)
}

// -----output-----
// true  	只声明数组类型的指针
// true		只声明数组类型
// true		初始化后的数组类型
// []		指针类型才可以添加值
// panic: reflect: reflect.Value.Set using unaddressable value  数组类型报错
//
// goroutine 1 [running]:
// reflect.flag.mustBeAssignable(0x97)
// /usr/local/go/src/reflect/value.go:234 +0x13a
// reflect.Value.Set(0x10a44c0, 0x118f8c0, 0x97, 0x10a44c0, 0xc000092020, 0x97)
// /usr/local/go/src/reflect/value.go:1467 +0x2f
// main.main()
// /Users/frank/go/src/github.com/FrankHitman/playGo/gorm/main/reflect.go:25 +0x883
//
