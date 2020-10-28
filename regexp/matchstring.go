package main

import (
	"fmt"
	"regexp"
)

func main() {
	// "(GET)|(POST)" "GET|POST"
	res, err := regexp.MatchString("^*$", "GET")
	if err != nil {
		panic(err)
	}
	fmt.Println(res)
}

// output
// true