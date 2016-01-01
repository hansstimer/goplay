package main

type A struct {
	a int
}

func (a *A) foo() (int) {
	return a.a
}

type B struct {
	b int
}

func (b *B) foo() (int) {
	return b.b
}

func main() {

}



