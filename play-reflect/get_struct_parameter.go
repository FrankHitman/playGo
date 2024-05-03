package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	Name string
	Age  int
}

func main() {
	p := Person{Name: "Alice", Age: 30}
	t := reflect.TypeOf(p)

	for i := 0; i < t.NumField(); i++ { // NumField returns a struct type's field count.
		field := t.Field(i) // Field returns a struct type's i'th field.
		fmt.Printf("Field %d: %s (%s)\n", i+1, field.Name, field.Type)
	}
}

// output
// Franks-Mac:play-reflect frank$ go run get_struct_parameter.go
// Field 1: Name (string)
// Field 2: Age (int)
