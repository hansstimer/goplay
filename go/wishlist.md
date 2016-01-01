Go Language Wish List
===

Literals
---
* binary
	
		10101b

* hex floating point

		0x1.fp3 // Useful for creating various numerical constants
	
* unsigned

		211u // Slightly more handy than uint8(211) when doing bit operations

FIXED: Less verbose literal struct arrays 
---
Given

	type Foo struct { 
	 a int 
	 b []struct{c, d int} 
	} 

Something less verbose than this

	var foo = Foo{1, []struct{c, d int}{ 
	    struct{c, d int}{2, 3}, 
	    struct{c, d int}{4, 6}, 
	}} 

Function default values
---

	func foo(i int = 2) {}
	foo();

Function named arguments
---

	foo(i:1)

Abbreviated map literals where a label is a string
---

	var m = map[string]int = {size:1, count:3, } // aka {"size":1, "count":3, }

Reverse conditional statements
---

	return(err) if err != 0
	exit(err) if err != 0

Composite literal initialization with embedded fields
---

	type Point struct { x, y int; }
	type MyPoint struct { Point; z int }
	
	a := MyPoint { Point{1, 2}, 3 }  // awkward
	a := MyPoint {1,2,3} // natural

Composite literal without stutter
---

	type A struct { a int; } 
	type B struct { A; b string; } 
	var b = B{ 
	        A: A{                 // We know we have an A, so why the stutter? 
	                a: 1, 
	        }, 
	        b: "b",                // Look Ma, no stuttering here! 
	        } 


Automatic literal map initialization
---

	type foo struct{ f map[string]int; g int; }
	var f = foo{ g:2 }

	func main() {
		// Blows up here; would be nice if it didn't need explicit initialization
		f.f["hey"] = 1		
	}
	
Literal number catenation
---
	
	million := 1 000 000
	e := 2.718 281 828
	
Range operator that returns pointers
---

	var a = []int{1, 2, 3}
	for i, p := rangep a {
		p *= p
	}
	// a == {2, 4, 6}
	
Here documents aka extended quotes
---

	s := <<LONGQUOTE
	Apple
	Banana
	Peach
	LONGQUOTE

Conversion from bool to int
---

	i := int(true)
	
Composite literal with addressed literal elements
---
	
	var a = []*int{&1, &2, &3,}			// Equivalent to: i := new(int); i = 1d
	var pts = []*Point{&{1,2}, &{3,4},}
	
Simpler container iteration
---
This is very wordy

	el := list.Front()
	for el != nil {
		h, ok := el.Value.(RequestI)
		if !ok {
			log.Panicln("Bogus RequestI")
		}
		
		el = el.Next()
	}
	
Instead of trying to come up with a general template mechanism that is 
acceptable, we might consider more limited goals specific for 
container creation. 

* Range operator syntax 
* Only compile time checks, no runtime checks 
* Same binary for all types that the template is instantiated on 

If we are only interested in generalizable containers then we can have 
some restrictions on template use: 

* Templates can only be declared for: 
     - base types 
     - interface types 
     - pointers to any type 
* A template can be declared for any number of types 
* Templates can be limited to instantiate only on types that have the 
same base type as the template or conform to the interface type 
* Templates must conform to a range operator interface 

This implies, for instance, that you can't have a single template that 
is for both int32 and float because they are different base types and 
require different binaries. However, you can easily have a template 
that takes any interface or any pointer. 

Another way to think of this kind of approach is that when you 
instantiate a template you are predeclaring the safe conversions that 
the compiler can automatically make when you use the container. This 
is about as far from typical templates as you can get and still be 
useful. 
	
	
