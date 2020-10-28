package slice

import "testing"

const N = 1000

func initSlice() []string {
	s := make([]string, N)
	for i := 0; i < N; i++ {
		s[i] = "www.baidu.com"
	}
	return s
}

func BenchmarkForSlice(b *testing.B) {
	s := initSlice()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ForSlice(s)
	}
}

func BenchmarkRangeForSlice(b *testing.B) {
	s := initSlice()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		RangeForSlice(s)
	}
}

func BenchmarkRangeForSlice2(b *testing.B) {
	s := initSlice()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		RangeForSlice2(s)
	}
}

// ----output-----
//Franks-Mac:slice frank$ go test -bench=. -run=NONE
//goos: darwin
//goarch: amd64
//pkg: github.com/FrankHitman/playGo/slice
//BenchmarkForSlice-8              5000000               277 ns/op
//BenchmarkRangeForSlice-8         3000000               476 ns/op
//BenchmarkRangeForSlice2-8        5000000               278 ns/op
//PASS
//ok      github.com/FrankHitman/playGo/slice     5.265s


//从性能测试可以看到，常规的for循环，要比for range的性能高出近一倍，到这里相信大家已经知道了原因，
// 没错，因为for range每次是对循环元素的拷贝，所以集合内的运算越复杂，性能越差，
// 而反观常规的for循环，它获取集合内元素是通过s[i]，这种索引指针引用的方式，要比拷贝性能要高的多。

// 比较奇怪的是在mac平台测试与在Linux测试结果不同
//goos: linux
//goarch: amd64
//BenchmarkForSlice-4              2000000               703 ns/op
//BenchmarkRangeForSlice-4         2000000               622 ns/op
//BenchmarkRangeForSlice2-4        5000000               362 ns/op
//PASS
//ok      _/home/dongyue/playGo   6.224s


