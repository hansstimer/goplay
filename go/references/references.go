package main

import (
	"fmt"
	"unsafe"
)

func foo(b []int, s string) {
	fmt.Printf("slice: %v\n", unsafe.Sizeof(b))
	fmt.Printf("string: %v\n", unsafe.Sizeof(s))
}

type A struct {
	a []int
	b map[string]int
}

func main() {
	var a = make([]int, 0, 10)
	a[0] = 1
	s := "hello"
	foo(a, s)

	var aa A
	fmt.Printf("%v\n", len(aa.a))
	fmt.Printf("%v\n", len(aa.b))
	// aa.b["1"] = 1 // error
}
