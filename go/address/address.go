package main

import (
	"fmt"
)

type A struct {
	a int
}

var as [10]*A
var bs [10]*A
var aa = [10]A{{0}, {1}, {2}, {3}, {4}, {5}, {6}, {7}, {8}, {9}}
var ab = [10]*A{&aa[0], &aa[1], &aa[2], &aa[3], &aa[4], &aa[5], &aa[6], &aa[7], &aa[8], &aa[9]}

func main() {
	for i, _ := range as {
		var a = A{i}
		as[i] = &a
		bs[i] = &a
	}
	
	for i, _ := range as {
		fmt.Printf("%p\n", as[i])
		fmt.Printf("%p\n", bs[i])
		fmt.Printf("%p\n", &aa[i])
		fmt.Printf("%p\n", ab[i])
		fmt.Println()
	}
}
