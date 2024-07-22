package main

import "math"

func main() {
	const PI float64 = 3.1415
	var radius = 3.2 // type inference for compiler

	// Reduced Form: Create a Variable
	area := PI * math.Pow(radius, 2)
	println("The area of the circle is", area)

	// Constants
	const (
		a = 1
		b = 2
	)

	// Variables
	var (
		c = 3
		d = 4
	)

	println(a, b, c, d)

	// Boolean: defined two or more variables in a single line
	var e, f bool = true, false
	println(e, f)

	g, h, i := 2, false, "hello!"
	println(g, h, i)
}
