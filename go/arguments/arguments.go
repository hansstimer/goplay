package main

import (
	"fmt"
)

var aMap = make(map[string]int, 4)

func maps(bMap map[string]int) {
	aMap["a"] = 0
	aMap["b"] = 1

	fmt.Printf("aMap: %#v\n", aMap)
	fmt.Printf("&aMap: %#v\n", &aMap)
	fmt.Printf("bMap: %#v\n", bMap)
	fmt.Printf("&bMap: %#v\n", &bMap)

	bMap["a"] = 100
	bMap["b"] = 102

	fmt.Printf("aMap: %#v\n", aMap)
	fmt.Printf("&aMap: %#v\n", &aMap)
	fmt.Printf("bMap: %#v\n", bMap)
	fmt.Printf("&bMap: %#v\n", &bMap)

	bMap["c"] = 103
	bMap["d"] = 104

	fmt.Printf("aMap: %#v\n", aMap)
	fmt.Printf("&aMap: %#v\n", &aMap)
	fmt.Printf("bMap: %#v\n", bMap)
	fmt.Printf("&bMap: %#v\n", &bMap)

	bMap["e"] = 105
	bMap["f"] = 106

	fmt.Printf("aMap: %#v\n", aMap)
	fmt.Printf("&aMap: %#v\n", &aMap)
	fmt.Printf("bMap: %#v\n", bMap)
	fmt.Printf("&bMap: %#v\n", &bMap)

}

var aArray = [2]int{0, 2}

func arrays(bArray [2]int) {
	fmt.Printf("aArray: %#v\n", aArray)
	fmt.Printf("&aArray: %p\n", &aArray)
	fmt.Printf("bArray: %#v\n", bArray)
	fmt.Printf("&bArray: %p\n", &bArray)

	bArray[0] = 100
	bArray[1] = 102

	fmt.Printf("aArray: %#v\n", aArray)
	fmt.Printf("&aArray: %p\n", &aArray)
	fmt.Printf("bArray: %#v\n", bArray)
	fmt.Printf("&bArray: %p\n", &bArray)
}

var aSlice = make([]int, 0, 4)

func slices(bSlice []int) {
	aSlice = append(aSlice, 100)
	aSlice = append(aSlice, 101)

	fmt.Printf("aSlice: %#v\n", aSlice)
	fmt.Printf("&aSlice: %p\n", &aSlice)
	fmt.Printf("bSlice: %#v\n", bSlice)
	fmt.Printf("&bSlice: %p\n", &bSlice)

	// bSlice[0] = 100
	// bSlice[1] = 101

	fmt.Printf("aSlice: %#v\n", aSlice)
	fmt.Printf("&aSlice: %p\n", &aSlice)
	fmt.Printf("bSlice: %#v\n", bSlice)
	fmt.Printf("&bSlice: %p\n", &bSlice)
}

var aString = "hello"

func strings(bString string) {
	fmt.Printf("aString: %p\n", &aString)
	fmt.Printf("bString: %p\n", &bString)
}

func main() {
	maps(aMap)
	arrays(aArray)
	slices(aSlice)
	strings(aString)
}
