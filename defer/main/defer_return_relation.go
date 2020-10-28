package main

import "fmt"

func main() {
	fmt.Println("f1 result: ", f1())
	fmt.Println("f2 result: ", f2())
}

func f1() int {
	var i int
	defer func() {
		i++
		fmt.Println("f11: ", i)
	}()

	defer func() {
		i++
		fmt.Println("f12: ", i)
	}()

	i = 1000
	return i
}

func f2() (i int) {
	defer func() {
		i++
		fmt.Println("f21: ", i)
	}()

	defer func() {
		i++
		fmt.Println("f22: ", i)
	}()

	i = 1000
	return i
}

//f12:  1001
//f11:  1002
//f1 result:  1000
//f22:  1001
//f21:  1002
//f2 result:  1002

//问题的关键是为什么无名参数返回的值是1000，其并未收到defer函数对于i自增的影响；而有名函数在执行defer后，最后返回的i值为1002。
//网上找了一些原因，提到一个结论
//
//原因就是return会将返回值先保存起来，对于无名返回值来说，
//保存在一个临时对象中，defer是看不到这个临时对象的；
//而对于有名返回值来说，就保存在已命名的变量中。
//go tool compile -S defer/main/defer_return_relation.go generate assembly language code
//作者：JackieZheng
//链接：https://juejin.im/post/5d173ce16fb9a07ea803df75
//来源：掘金
//著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。