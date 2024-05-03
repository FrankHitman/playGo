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
	t := reflect.TypeOf(Person{})
	p := reflect.New(t).Elem()
	p.FieldByName("Name").SetString("Bob")
	p.FieldByName("Age").SetInt(25)

	person := p.Interface().(Person) // type assertion
	fmt.Println("Dynamic object:", person)
}

// Franks-Mac:play-reflect frank$ go run create_object.go
// Dynamic object: {Bob 25}
