package main

import (
	"os"
	"runtime"
	"runtime/trace"
)

func gcfinished() *int {
	p := 1
	runtime.SetFinalizer(&p, func(_ *int) {
		println("gc finished")
	})
	return &p
}
func allocate() {
	_ = make([]byte, int((1<<20)*0.25))
}
func main() {
	f, _ := os.Create("trace.out")
	defer f.Close()
	trace.Start(f)
	defer trace.Stop()
	gcfinished()
	// 当完成 GC 时停止分配
	for n := 1; n < 50; n++ {
		println("#allocate: ", n)
		allocate()
	}
	println("terminate")
}

// output
// Franks-Mac:main frank$ GODEBUG=gctrace=1 go run gc_example1.go
// gc 1 @0.037s 2%: 0.027+1.8+0.005 ms clock, 0.21+4.5/2.9/0.36+0.041 ms cpu, 4->4->0 MB, 5 MB goal, 8 P
// gc 2 @0.051s 2%: 0.065+1.1+0.033 ms clock, 0.52+1.2/1.2/0.002+0.26 ms cpu, 4->4->0 MB, 5 MB goal, 8 P
// gc 3 @0.097s 1%: 0.077+0.75+0.031 ms clock, 0.61+1.1/0.96/0+0.25 ms cpu, 4->5->1 MB, 5 MB goal, 8 P
// gc 4 @0.127s 1%: 0.021+0.49+0.023 ms clock, 0.17+0.10/0.55/0.37+0.18 ms cpu, 4->4->0 MB, 5 MB goal, 8 P
// gc 5 @0.165s 1%: 0.026+0.50+0.024 ms clock, 0.20+0.15/0.37/0.65+0.19 ms cpu, 4->4->0 MB, 5 MB goal, 8 P
// gc 6 @0.190s 1%: 0.29+2.5+0.16 ms clock, 2.3+2.4/1.6/0+1.3 ms cpu, 4->4->0 MB, 5 MB goal, 8 P
// gc 7 @0.336s 0%: 0.035+0.51+0.005 ms clock, 0.28+0.35/0.60/1.2+0.043 ms cpu, 4->4->0 MB, 5 MB goal, 8 P
// gc 8 @0.347s 0%: 0.018+0.63+0.006 ms clock, 0.14+0/0.44/1.7+0.050 ms cpu, 4->4->0 MB, 5 MB goal, 8 P
// gc 9 @0.359s 0%: 0.020+0.36+0.022 ms clock, 0.16+0.22/0.36/1.0+0.17 ms cpu, 4->4->0 MB, 5 MB goal, 8 P
// gc 10 @0.371s 0%: 0.024+0.31+0.055 ms clock, 0.19+0/0.38/1.1+0.44 ms cpu, 4->4->0 MB, 5 MB goal, 8 P
// # command-line-arguments
// gc 1 @0.023s 2%: 0.036+6.0+0.046 ms clock, 0.29+3.4/2.2/9.1+0.37 ms cpu, 4->4->2 MB, 5 MB goal, 8 P
// gc 2 @0.049s 2%: 0.003+1.3+0.003 ms clock, 0.031+0.92/2.1/3.0+0.031 ms cpu, 5->5->4 MB, 6 MB goal, 8 P
// # command-line-arguments
// gc 1 @0.001s 8%: 0.024+3.8+0.013 ms clock, 0.19+0.25/3.3/0.81+0.10 ms cpu, 4->5->5 MB, 5 MB goal, 8 P
// gc 2 @0.023s 3%: 0.004+3.3+0.019 ms clock, 0.038+0/4.4/1.9+0.15 ms cpu, 9->10->9 MB, 10 MB goal, 8 P
// gc 3 @0.067s 2%: 0.003+3.9+0.020 ms clock, 0.026+0.43/5.7/10+0.16 ms cpu, 16->17->15 MB, 18 MB goal, 8 P
// gc 4 @0.155s 2%: 0.006+6.0+0.019 ms clock, 0.051+0.22/10/14+0.15 ms cpu, 29->31->27 MB, 31 MB goal, 8 P
// #allocate:  1
// #allocate:  2
// #allocate:  3
// #allocate:  4
// #allocate:  5
// #allocate:  6
// #allocate:  7
// #allocate:  8
// #allocate:  9
// #allocate:  10
// #allocate:  11
// #allocate:  12
// #allocate:  13
// #allocate:  14
// #allocate:  15
// #allocate:  16
// #allocate:  17
// #allocate:  18
// #allocate:  19
// #allocate:  20
// #allocate:  21
// #allocate:  22
// #allocate:  23
// #allocate:  24
// #allocate:  25
// gc finished
// #allocate:  26
// gc 1 @0.001s 3%: 0.011+0.34+0.049 ms clock, 0.091+0.10/0.052/0.22+0.39 ms cpu, 4->6->2 MB, 5 MB goal, 8 P
// #allocate:  27
// #allocate:  28
// #allocate:  29
// #allocate:  30
// #allocate:  31
// #allocate:  32
// #allocate:  33
// #allocate:  34
// #allocate:  35
// gc 2 @0.003s 7%: 0.006+0.40+0.17 ms clock, 0.048+0.12/0.095/0.18+1.3 ms cpu, 4->4->0 MB, 5 MB goal, 8 P
// #allocate:  36
// #allocate:  37
// #allocate:  38
// #allocate:  39
// #allocate:  40
// #allocate:  41
// #allocate:  42
// #allocate:  43
// #allocate:  44
// #allocate:  45
// #allocate:  46
// #allocate:  47
// #allocate:  48
// #allocate:  49
// terminate
// gc 3 @0.005s 6%: 0.042+0.41+0.006 ms clock, 0.33+0.15/0.11/0.19+0.048 ms cpu, 4->4->0 MB, 5 MB goal, 8 P
// Franks-Mac:main frank$
