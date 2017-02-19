package main

import (
	"math"
	"fmt"
)
/*
	Implementation of abstract factory pattern with structs interface
 */

// abstract interface with methods declared
type geometry interface {
	area() float64;
	perim() float64;
}

// rectangle struct declaration
type rect struct {
	height, width float64
}

	// rectangle area implementation. E.g. rect.area()
	func (r rect) area() float64  {
		return r.height * r.width
	}

	// rectangle perimeter. E.g. rect.perim()
	func (r rect) perim() float64 {
		return 2*r.height + 2*r.width
	}

// circle struct declaration
type circle struct {
	radius float64
}

	// circle area implementation. E.g. circle.area()
	func (c circle) area() float64 {
		return math.Pi*c.radius*c.radius
	}

	// circle perimeter. E.g. circle.perim()
	func (c circle) perim() float64 {
		return 2* math.Pi*c.radius
	}

// factory method
func measure(g geometry)  {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

/*
	geometry (interface)
		|area();
		|perim();
	/		\
      rect	      circle  => Both of these should implement area() and perim().
      				Otherwise, error: "cannot use c (type circle) as type geometry in argument to measure:
				circle does not implement geometry (missing perim method)"

 */
func main() {

	r := rect{height:20, width:30}
	c := circle{radius:35}

	measure(r)
	measure(c)
}
