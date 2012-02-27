package main

// type A struct {
// 	a int
// }
// 
// type B struct {
// 	as []*A
// }
// 
// var b = B { 
// 	as: []*A{
// 		 0: &{1},
// 		 },
// 		}

var as = []int{1, 2, 3}

func foo() *int {
	return &as[1]
}

func main() {
	
}
