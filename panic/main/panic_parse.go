// panic_package.go
package main

import (
	"../parse"
	"fmt"
)

func main() {
	var examples = []string{
		"1 2 3 4 5",
		"100 50 25 12.5 6.25",
		"2 + 2 = 4",
		"1st class",
		"",
	}

	for _, ex := range examples {
		fmt.Printf("Parsing %q:\n  ", ex)
		nums, err := parse.Parse(ex)
		if err != nil {
			fmt.Println("get err in main ", err.Error()) // here String() method from ParseError is used
			continue
		}
		fmt.Println(nums)
	}
}

//----output----
//Parsing "1 2 3 4 5":
//[1 2 3 4 5]
//Parsing "100 50 25 12.5 6.25":
//pkg: pkg parse: error parsing "12.5" as int
//Parsing "2 + 2 = 4":
//pkg: pkg parse: error parsing "+" as int
//Parsing "1st class":
//pkg: pkg parse: error parsing "1st" as int
//Parsing "":
//pkg: no words to parse

//Parsing "1 2 3 4 5":
//[1 2 3 4 5]
//Parsing "100 50 25 12.5 6.25":
//pkg parse: error parsing "12.5" as int
//false
//<nil>
//pkg: pkg parse: error parsing "12.5" as int
//Parsing "2 + 2 = 4":
//pkg parse: error parsing "+" as int
//false
//<nil>
//pkg: pkg parse: error parsing "+" as int
//Parsing "1st class":
//pkg parse: error parsing "1st" as int
//false
//<nil>
//pkg: pkg parse: error parsing "1st" as int
//Parsing "":
//no words to parse
//false
//<nil>
//pkg: no words to parse
