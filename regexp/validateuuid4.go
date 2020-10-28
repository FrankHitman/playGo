package main

import (
	"fmt"
	"regexp"
)

func main() {
	str := "ffd5ffff-e46b-4fff-bf8e-50ffd0ffffff"
	// this regexp string in gopkg.in/go-playground/validator.v9/regexes.go which used in baked_in.go
	uUID4RegexString := "^[0-9a-f]{8}-[0-9a-f]{4}-4[0-9a-f]{3}-[89ab][0-9a-f]{3}-[0-9a-f]{12}$"
	uUID4Regex := regexp.MustCompile(uUID4RegexString)
	fmt.Println(uUID4Regex.MatchString(str))
}
