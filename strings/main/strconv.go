package main

import (
	"fmt"
	"strconv"
)

func main() {
	var orig string = "666e"
	var an int
	var newS string
	var err error

	fmt.Printf("The size of ints is: %d\n", strconv.IntSize)

	an, err = strconv.Atoi(orig)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("The integer is: %d\n", an)
	an = an + 5
	newS = strconv.Itoa(an)
	fmt.Printf("The new string is: %s\n", newS)
}


// -----output-----
//The size of ints is: 64
//strconv.Atoi: parsing "666e": invalid syntax
//The integer is: 0
//The new string is: 5

