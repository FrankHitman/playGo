package main

import "sync"

type Lockable[T any] struct {
	sync.Mutex
	Data T
}

func main() {
	var n Lockable[uint32]
	n.Lock()
	n.Data++
	n.Unlock()

	var f Lockable[float64]
	f.Lock()
	f.Data += 1.23
	f.Unlock()

	var b Lockable[bool]
	b.Lock()
	b.Data = !b.Data
	b.Unlock()

	var bs Lockable[[]byte]
	bs.Lock()
	bs.Data = append(bs.Data, "Go"...)
	bs.Unlock()
}

// // any is an alias for interface{} and is equivalent to interface{} in all ways.
// type any = interface{}

// refer to https://go101.org/generics/444-first-look-of-custom-generics.html
