package main

import "fmt"

type Stringer interface {
	String() string
}

type Printer struct {
	name string
}

func (p *Printer) String() string {
	return p.name
}

func main() {
	var v = "hello"
	checkStringer(v)
	var t = new(Printer)
	t.name = "tom"
	checkStringer(t)

}

func checkStringer(v interface{})  {
	if sv, ok := v.(Stringer); ok {
		fmt.Printf("v implements String(): %s\n", sv.String()) // note: sv, not v
	}else {
		fmt.Println("v not implements String(): ", v, sv) // note: sv, not v
	}
}

// ----output----
//v not implements String():  hello <nil>
//v implements String(): tom
