package main

import (
	"fmt"
	"sort"
)

// GuessingGame test what the function of sort.Search actually does.
func GuessingGame() {
	var s string
	fmt.Printf("Pick an integer from 0 to 100.\n")
	answer := sort.Search(100, func(i int) bool {
		fmt.Printf("Is your number <= %d? ", i)
		fmt.Scanf("%s", &s)
		return s != "" && s[0] == 'y'
	})
	fmt.Printf("Your number is %d.\n", answer)
}
func main() {
	GuessingGame()
}

// output
// Franks-Mac:system-design frank$ go run guess-game.go
// Pick an integer from 0 to 100.
// Is your number <= 50? n
// Is your number <= 75? n
// Is your number <= 88? n
// Is your number <= 94? n
// Is your number <= 97? n
// Is your number <= 99? y
// Is your number <= 98? y
// Your number is 98.
