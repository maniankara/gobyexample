package main

import "fmt"

// rectangle struct declaration
type rect struct {
	height, width int
}
// rectangle area implementation
func (r rect) area() int {
	return r.width * r.height
}
// rectangle permeter implementation
func (r rect) perm() int {
	return 2*r.width + 2*r.height
}

func main() {
	r1 := rect{5,4}
	fmt.Println(r1.area())
	fmt.Println(r1.perm())

}
