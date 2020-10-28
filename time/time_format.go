package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()

	fmt.Println("20060102150405 format is: ", now.Format("20060102150405"))
	fmt.Println("20060102030405 format is: ", now.Format("20060102030405"))

	newT, err := time.Parse("20060102030405", "20191101121212")
	if err != nil {
		fmt.Println("20191101121212 parse error: ", err)
	} else {
		fmt.Println(newT.Hour())
		fmt.Println(newT)
	}

	newT, err = time.Parse("20060102030405", "20191101131212")
	if err != nil {
		fmt.Println("20191101121212 parse error: ", err)
	} else {
		fmt.Println(newT.Hour())
		fmt.Println(newT)
	}

	newT, err = time.Parse("20060102150405", "20191101131212")
	if err != nil {
		fmt.Println("20191101121212 parse error: ", err)
	} else {
		fmt.Println(newT.Hour())
		fmt.Println(newT)
	}
}

// output
// 20060102150405 format is:  20191101150142 		当前时间的24小时制
// 20060102030405 format is:  20191101030142		当前时间的12小时制
// 12
// 2019-11-01 12:12:12 +0000 UTC
// 20191101121212 parse error:  parsing time "20191101131212": hour out of range 	转化为12小时制时候注意小时是否超限
// 13
// 2019-11-01 13:12:12 +0000 UTC