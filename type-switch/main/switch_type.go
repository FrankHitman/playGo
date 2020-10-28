package main

import "fmt"

func classifier(items ...interface{}) {
	for i, x := range items {
		switch x.(type) {
		case bool:
			fmt.Printf("Param #%d is a bool\n", i)
		case float64:
			fmt.Printf("Param #%d is a float64\n", i)
		case int, int64:
			fmt.Printf("Param #%d is a int\n", i)
		case nil:
			fmt.Printf("Param #%d is a nil\n", i)
		case string:
			fmt.Printf("Param #%d is a string\n", i)
		default:
			fmt.Printf("Param #%d is unknown\n", i)
		}
	}
}

func main() {
	classifier(13, -14.3, "BELGIUM", complex(1, 2), nil, false)
}

// -------output------
//Param #0 is a int
//Param #1 is a float64
//Param #2 is a string
//Param #3 is unknown
//Param #4 is a nil
//Param #5 is a bool
