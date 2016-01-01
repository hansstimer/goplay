package main

import (
	"fmt"
)

type A struct{ a, b int }

// Remove the * from in front of A, and this wont compile
var m = map[string]A{"a": {1, 10}, "b": {2, 20}}

func main() {
	for _, value := range m {
		fmt.Printf("%v\n", value)
	}

	for _, value := range m {
		value.a += 1
	}

	for _, value := range m {
		fmt.Printf("%v\n", value)
	}

	// for key := range m {
	// 	m[key].a += 1
	// }

	// for _, value := range m {
	// 	fmt.Printf("%v\n", value)
	// }

}
