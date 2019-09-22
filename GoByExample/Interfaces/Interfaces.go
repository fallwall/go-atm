// interfaces are named collections of method signatures.package Interfaces

package main

import (
	"fmt"
	"math"
)

// here's basic interface for geometri shapes
type geometry interface {
	area() float64
	perim() float64
}

// let's implements this interface on rect and cicle types
type rect struct {
	width, height float64
}

type circle struct {
	radius float64
}

//implement geometry on rects

func (r rect) area() float64 {
	return r.width * r.height
}

func (r rect) perim() float64 {
	return 2*r.width + 2*r.height
}

//implement geometry on circles

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

// if a variable has an interface type, we call methods are in the named interface:
func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func main() {
	r := rect{width: 3, height: 4}
	c := circle{radius: 5}

	measure(r)
	measure(c)
}