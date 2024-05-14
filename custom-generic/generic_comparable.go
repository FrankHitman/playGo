package main

// NoDiff checks whether or not a collection
// of values are all identical.
func NoDiff[V comparable](vs ...V) bool {
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
	var NoDiffString = NoDiff[string]
	println(NoDiff("Go", "Go", "Go")) // true
	println(NoDiffString("Go", "go")) // false

	println(NoDiff(123, 123, 123, 123)) // true
	println(NoDiff[int](123, 123, 789)) // false

	type A = [2]int
	println(NoDiff(A{}, A{}, A{}))     // true
	println(NoDiff(A{}, A{}, A{1, 2})) // false

	println(NoDiff(new(int)))           // true
	println(NoDiff(new(int), new(int))) // false

	println(NoDiff[bool]()) // true

	// _ = NoDiff() // error: cannot infer V

	// error: *** does not implement comparable
	// _ = NoDiff([]int{}, []int{})
	// _ = NoDiff(map[string]int{})
	// _ = NoDiff(any(1), any(1))
}

// // comparable is an interface that is implemented by all comparable types
// // (booleans, numbers, strings, pointers, channels, arrays of comparable types,
// // structs whose fields are all comparable types).
// // The comparable interface may only be used as a type parameter constraint,
// // not as the type of a variable.
// type comparable interface{ comparable }

// Please note that all of these type arguments implement the comparable interface.
// Incomparable types, such as []int and map[string]int may not be passed as type arguments of calls
// to the NoDiff generic function.
// And please note that, although any is a comparable (value) type, it doesn't implement comparable,
// so it is also not an eligible type argument.

// output
// Franks-Mac:playGo frank$ go run custom-generic/play_comparable.go
// true
// false
// true
// false
// true
// false
// true
// false
// true

// refer to https://go101.org/generics/444-first-look-of-custom-generics.html
