// Demonstrates that you can dynamically convert from an interface to a struct
// which is insteresting because where is the type of the struct being kept?

package main

import "fmt"

type Fooer interface {
	Foo()
}

type Goo struct {
	goo int
}

func (*Goo) Foo() {}

func faa(foo Fooer) {
	_, ok := foo.(*Goo)
	fmt.Println(ok)
}

func main() {
	g := Goo{1}
	faa(&g)
}
