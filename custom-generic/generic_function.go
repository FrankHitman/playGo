package main

type Ordered interface {
	~int | ~uint | ~int8 | ~uint8 | ~int16 | ~uint16 |
		~int32 | ~uint32 | ~int64 | ~uint64 | ~uintptr |
		~float32 | ~float64 | ~string
}

func Max[S ~[]E, E Ordered](vs S) E {
	if len(vs) == 0 {
		panic("no elements")
	}

	var r = vs[0]
	for i := range vs[1:] {
		if vs[i] > r {
			r = vs[i]
		}
	}
	return r
}

type Age int

var ages = []Age{99, 12, 55, 67, 32, 3}

var langs = []string{"C", "Go", "C++"}

func main() {
	// var maxAge = Max[[]Age, Age] // 实例化范型函数，并给这个函数指定别名。
	// println(maxAge(ages))        // 99
	//
	// var maxStr = Max[[]string, string]
	// println(maxStr(langs)) // Go

	// var maxAge = Max[[]Age] // partial argument list 只指定部分类型参数，因为第二个参数可以推导出来。
	// println(maxAge(ages))   // 99
	//
	// var maxStr = Max[[]string] // partial argument list
	// println(maxStr(langs))     // Go

	println(Max(ages))  // 99 甚至可以从传值的参数类型进行推导，从而完全省略
	println(Max(langs)) // Go
}

// output
// Franks-Mac:playGo frank$ go run custom-generic/generic_function.go
// 99
// Go
