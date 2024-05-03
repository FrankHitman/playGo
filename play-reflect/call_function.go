package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	Name string
	Age  int
}

func (p Person) Greet() {
	fmt.Printf("Hello, my name is %s and I'm %d years old.\n", p.Name, p.Age)
}

func main() {
	person := Person{Name: "Alice", Age: 30}
	v := reflect.ValueOf(person)
	method := v.MethodByName("Greet")
	method.Call(nil)
}

// MethodByName returns a function value corresponding to the method
// of v with the given name.
// The arguments to a Call on the returned function should not include
// a receiver; the returned function will always use v as the receiver.
// It returns the zero Value if no method was found.

// output
// Franks-Mac:playGo frank$ go run play-reflect/call_function.go
// Hello, my name is Alice and I'm 30 years old.
