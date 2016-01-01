package main

import (
	"fmt"
	"reflect"
)

var _ = fmt.Println

func grow(a interface{}, pos ...int) interface{} {
	v := reflect.ValueOf(a).Elem()
	p := pos[0]
	if p > v.Len() {
		cap := v.Cap()
		if p >= cap {
			vv := reflect.MakeSlice(v.Type(), p, p*2)
			reflect.Copy(vv, v)
			v.Set(vv)
		}
		v.SetLen(p)
	}

	if len(pos) > 1 {
		vi := v.Index(p - 1)
		inf := grow(vi.Addr().Interface(), pos[1:]...)
		v.Index(p - 1).Set(reflect.ValueOf(inf))
	}

	return v.Interface()
}

func main() {
	// a := []int{10, 11}
	// grow(&a, 4)
	// a := [][]int{{10, 11}, {20, 21, 22}}
	a := [][]int{{}}
	grow(&a, 4, 3)
	fmt.Printf("a: %v\n", a)
}
