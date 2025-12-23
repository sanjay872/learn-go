package examples

import "fmt"

type rect struct {
	w int
	h int
}

func (r *rect) area() int {
	return r.w * r.h
}

func (r rect) peri() int {
	return 2*r.w + 2*r.h
}

func MethodStruct() {
	r := rect{4, 4}

	fmt.Println("Area: ", r.area())
	fmt.Println("Perimeter: ", r.peri())

	s := &r

	fmt.Println("Area: ", s.area())
	fmt.Println("Preimeter: ", s.peri())
}
