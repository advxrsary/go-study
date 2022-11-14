// Golang basics tutorial

package main

import (
	"fmt"
)

func main() {
	// Variables

	var a int            // int
	var b float64        // float64
	var c string         // string
	var d bool           // bool
	var e []int          // slice
	var f map[string]int // map
	var g struct {       // struct
		x int // struct field
	}
	var h *int        // pointer
	var i interface{} // interface

	// assign values to variables
	a = 1
	b = 2.2
	c = "three"
	d = true
	e = []int{1, 2, 3}
	f = map[string]int{"one": 1, "two": 2, "three": 3}
	g = struct {
		x int
	}{x: 1}
	h = &a
	i = 1

	// print all variables values
	fmt.Printf("\nVariables\n")
	fmt.Printf("int: \t\t%v\n", a)
	fmt.Printf("float64: \t%v\n", b)
	fmt.Printf("string: \t%v\n", c)
	fmt.Printf("bool: \t\t%v\n", d)
	fmt.Printf("slice: \t\t%v\n", e)
	fmt.Printf("map: \t\t%v\n", f)
	fmt.Printf("struct: \t%v\n", g)
	fmt.Printf("pointer: \t%v\n", h)
	fmt.Printf("interface: \t%v\n", i)

}
