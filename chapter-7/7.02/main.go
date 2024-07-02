package main

import (
	"fmt"
)

type Shape interface {
	Area() float64
	Name() string
}

/*
all of these structs will satisfy the Shape interface since they have the
respective methods defined and implemented
*/
type triangle struct {
	base   float64
	height float64
}

type rectangle struct {
	height float64
	width  float64
}

type square struct {
	side float64
}

func (t triangle) Area() float64 {
	return (t.base * t.height) / 2
}

func (t triangle) Name() string {
	return "triangle"
}

func (r rectangle) Area() float64 {
	return r.width * r.height
}

func (r rectangle) Name() string {
	return "rectangle"
}

func (s square) Area() float64 {
	return s.side * s.side
}

func (s square) Name() string {
	return "square"
}

/*
demonstration of polymorphism. the ability to be in various forms.
ie, a shape can be a triangle, square, rectangle, etc.
*/

func main() {
	t := triangle{base: 15.5, height: 20.1}
	r := rectangle{height: 20, width: 10}
	s := square{side: 10}
	printShareDetails(t, r, s)
}

func printShareDetails(shapes ...Shape) {
	for _, item := range shapes {
		fmt.Printf("The area of %s is %.2f\n", item.Name(), item.Area())
	}
}
