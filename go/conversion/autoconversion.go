/*
This is an example of one way container specific templates could be done. I'm calling them autocasts because that is really all that is happening.

StructType     = "struct" [AutoCastDecl] "{" { FieldDecl ";" } "}" .
AutoCastDecl   = TypeName "(" TypeName ")"

* An autocast (i.e. template) can only be declared for: 
     - base types 
     - interface types 
     - pointers to any type 
* A autocast can be declared for any number of types 
* Autocasts can be limited to instantiate only on types that have the 
same base type as the autocast or conform to the interface type 
i.e. if an explict cast is legal then the autocast can be declared
* Autocasts may conform to a range operator interface 

type interface range {
	New(interface{})
	Next() interface{}
}

Advantages:
* You don't need to do anything special to create a container that can be autocasted
* Same mechanism works for all base types
* Compile time casts -- no runtime overhead
* No binary bloat
* No new keywords
* Minimal syntactical addition

Non-problems:
* You can't write compile time metaprograms

Example:
*/

// Plain everyday go package -- no autocast declarations
package tree

type TreeElement interface {
	compare(t *TreeElement) int
}

type Tree struct { }
func (t *Tree) New() *Tree { }
func (t *Tree) Add(TreeElement{}) { }
func (t *Tree) Remove(TreeElement{}) { }
func (t *Tree) Find(TreeElement{}) (*TreeElement) { }

type TreeRange struct {
	tree *Tree
	next *Next
}

func (tr *TreeRange) New(t* Tree) TreeElement {}
func (tr *TreeRange) Next() TreeElement {}



// Package makes use of an autocast
package main

import "tree"

// Conform to the TreeElement interface
type Foo struct { }
func (f *Foo) compare(t *tree.TreeElement) int {}
func (f *Foo) fud() {}

// Declare an autocast type
type FooTree struct tree.Tree(Foo) { }

func main() {
	f := Foo{}
	fooTree := FooTree.New()
	fooTree.Add(f)
	
	for f1 := range fooTree {
		f1.fud()		// No cast necessary to turn into a Foo
	}
}
