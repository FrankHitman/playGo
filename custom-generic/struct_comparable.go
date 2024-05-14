package main

type NoDiff[V comparable] struct{}

func (nd NoDiff[V]) Do(vs ...V) bool {
	if len(vs) == 0 {
		return true
	}

	v := vs[0]
	for _, x := range vs[1:] {
		if v != x {
			return false
		}
	}
	return true
}

func main() {
	var NoDiffString = NoDiff[string]{}.Do
	println(NoDiffString("Go", "go")) // false

	println(NoDiff[int]{}.Do(123, 123, 789)) // false

	println(NoDiff[*int]{}.Do(new(int))) // true
}

// output
// Franks-Mac:playGo frank$ go run custom-generic/struct_comparable.go
// false
// false
// true

// Generic functions could be viewed as simplified forms of methods of generic types

// refer to https://go101.org/generics/444-first-look-of-custom-generics.html
