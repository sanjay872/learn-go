package examples

import (
	"fmt"
	"math"
)

type geometry interface {
	area() float64
	perim() float64
}

type rec struct {
	width, height float64
}

type circle struct {
	radius float64
}

func (r rec) area() float64 {
	return r.width * r.height
}

func (r rec) perim() float64 {
	return 2*r.width + 2*r.height
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

func measurment(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func detectShape(g geometry) {
	if c, ok := g.(circle); ok {
		fmt.Println("Circle with radius: ", c.radius)
	}

	if r, ok := g.(rec); ok {
		fmt.Println("Rectangle with width: ", r.width, " and height: ", r.height)
	}
}

func Interfaces() {
	r := rec{4, 4}
	c := circle{2}

	measurment(r)
	measurment(c)

	detectShape(r)
	detectShape(c)
}
