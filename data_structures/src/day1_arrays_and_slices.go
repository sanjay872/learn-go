package src

import "fmt"

func ArraysAndSlices() {
	// arrays
	var ar [10]int

	// slices
	var sl []int

	for i := range 10 {
		ar[i] = i
		sl = append(sl, i)
	}

	// slice with length and capacity
	s := make([]int, 5, 10)
	fmt.Println(s)

	s = append(s, 10)
	fmt.Println(s)

	// copy and references

	a := make([]int, 10)
	b := a                     // b will be refering to the same data
	c := append([]int{}, a...) // deep copy

	fmt.Println(b)
	fmt.Println(c)

}

func ReverseElements() {

}

func RotateArray() {

}
