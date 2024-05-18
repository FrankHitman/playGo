package main

import "fmt"

func main() {
	// 以下展示了在数组中塞入不同类型的值。interface{} 就有这种动态属性
	values := []interface{}{
		456, "abc", true, 0.33, int32(789),
		[]int{1, 2, 3}, map[int]bool{}, nil,
	}
	for _, x := range values {
		// Here, v is declared once, but it denotes
		// different variables in different branches.
		switch v := x.(type) {
		case []int: // a type literal
			// The type of v is "[]int" in this branch.
			fmt.Println("int slice:", v)
		case string: // one type name
			// The type of v is "string" in this branch.
			fmt.Println("string:", v)
		case int, float64, int32: // multiple type names
			// The type of v is "interface{}",
			// the same as x in this branch.
			fmt.Println("number:", v)
		case nil:
			// The type of v is "interface{}",
			// the same as x in this branch.
			fmt.Println(v)
		default:
			// The type of v is "interface{}",
			// the same as x in this branch.
			fmt.Println("others:", v)
		}
		// Note, each variable denoted by v in the
		// last three branches is a copy of x.
	}

	// type-switch is equivalent to code below
	// for _, x := range values {
	//		if v, ok := x.([]int); ok {
	//			fmt.Println("int slice:", v)
	//		} else if v, ok := x.(string); ok {
	//			fmt.Println("string:", v)
	//		} else if x == nil {
	//			v := x
	//			fmt.Println(v)
	//		} else {
	//			_, isInt := x.(int)
	//			_, isFloat64 := x.(float64)
	//			_, isInt32 := x.(int32)
	//			if isInt || isFloat64 || isInt32 {
	//				v := x
	//				fmt.Println("number:", v)
	//			} else {
	//				v := x
	//				fmt.Println("others:", v)
	//			}
	//		}
	//	}
}

// output
// Franks-Mac:playGo frank$ go run interface/main/type_switch.go
// number: 456
// string: abc
// others: true
// number: 0.33
// number: 789
// int slice: [1 2 3]
// others: map[]
// <nil>

// refer to http://go101.org/article/interface.html
