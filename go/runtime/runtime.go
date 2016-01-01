package main

import (
	"runtime"
	"fmt"
)

func main() {
	fmt.Println("GOMAXPROCS:" runtime.GOMAXPROCS(0) )
}
