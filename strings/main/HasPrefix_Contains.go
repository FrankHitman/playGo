package main

import (
"fmt"
"strings"
)

func main() {
	var str string = "This is an example of a string"
	fmt.Printf("T/F? Does the string \"%s\" have prefix %s? ", str, "Th")
	fmt.Printf("%t\n", strings.HasPrefix(str, "Th"))
	fmt.Println(strings.Contains(str, "an"))
	strings.NewReader(str)
}

//// Contains reports whether substr is within s.
//func Contains(s, substr string) bool {
//	return Index(s, substr) >= 0
//}