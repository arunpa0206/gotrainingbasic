package main

import (
	"fmt"
)

// Geometry is an interface that defines Geometrical Calculation
type shape interface {
	draw() string
}

// Pentagon defines a geometrical object
type Rectangle struct{}

// Hexagon defines a geometrical object
type Circle struct{}

// Edges implements the Geometry interface
func (r Rectangle) draw() string {
	return "Drawing Rectangle"
}

// Edges implements the Geometry interface
func (c Circle) draw() string {
	return "Drawing Circle"
}

func sketch(s shape) {
	fmt.Println(s)
	fmt.Println(s.draw())
}

// main is the entry point for the application.

func main() {
	r := new(Rectangle)
	c := new(Circle)
	sketch(r)
	sketch(c)

}
