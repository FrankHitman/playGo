package _map

import (
	"fmt"
	"testing"
)

const N = 1000

func initMap() map[int]string {
	m := make(map[int]string, N)
	for i := 0; i < N; i++ {
		m[i] = fmt.Sprint("www.baidu.com:", i)
	}
	return m
}

func BenchmarkRangeForMap1(b *testing.B) {
	m:= initMap()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		RangeForMap1(m)
	}
}


func BenchmarkRangeForMap2(b *testing.B) {
	m := initMap()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		RangeForMap2(m)
	}
}



func BenchmarkRangeForMap3(b *testing.B) {
	m := initMap()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		RangeForMap3(m)
	}
}

// ---output-----
//Franks-Mac:map frank$ go test -bench=. -run=NONE
//goos: darwin
//goarch: amd64
//pkg: github.com/FrankHitman/playGo/map
//BenchmarkRangeForMap1-8           100000             12571 ns/op
//BenchmarkRangeForMap2-8           100000             20227 ns/op
//BenchmarkRangeForMap3-8           100000             12603 ns/op
//PASS
//ok      github.com/FrankHitman/playGo/map       5.008s
