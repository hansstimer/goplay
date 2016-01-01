package main

import (
	"fmt"
	"unsafe"
)

type PointI interface {
	getx() int32
}

type Point struct {
	x, y int32
}

type P struct {
	s, y int32
}

func (p *Point) getx() int32 {
	return p.x
}

type MyPoint struct {
	x int32
	Point
}

var a = MyPoint{3, Point{1, 2}, }

func foo(p *Point) {
	p.getx()
}

var ps = []PointI {&a, &Point{4, 5}, &MyPoint{6, Point{7, 8}}, &Point{9, 10}}

func main() {
	foo(&a.Point)
	p := Point{0,1}
	pi := PointI(&p)
	ai := PointI(&a)
	
	fmt.Printf("sizeof p:Point: %d\n", unsafe.Sizeof(p))
	fmt.Printf("sizeof p:&Point: %d\n", unsafe.Sizeof(&p))
	fmt.Printf("sizeof p:PointI: %d\n", unsafe.Sizeof(pi))
	fmt.Printf("sizeof ai:PointI: %d\n", unsafe.Sizeof(ai))

	q := Point{}
	fmt.Printf("sizeof q:Point: %d\n", unsafe.Sizeof(q))

	r := P{}
	fmt.Printf("sizeof r:P: %d\n", unsafe.Sizeof(r))
	
	i := int32(1)
	fmt.Printf("sizeof i:int32: %d\n", unsafe.Sizeof(i))
	

	fmt.Printf("%v\n", p)
	fmt.Printf("%#v\n", p)
	fmt.Printf("%T\n", p)
	fmt.Printf("%v\n", a.x)
	fmt.Printf("%v\n", a.getx())
	
}
