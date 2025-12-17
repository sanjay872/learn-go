package examples

import (
	"fmt"
	"slices"
)

func Slices() {
	var s []string

	fmt.Println("un initialised: ", s, s == nil, len(s) == 0)

	s = make([]string, 3)

	fmt.Println("initialised: ", s, s == nil, len(s) == 0)

	// access
	s[0] = "1"
	s[1] = "2"
	s[2] = "3"

	fmt.Println("Set:", s)
	fmt.Println("get", s[1])

	fmt.Println("len:", len(s))

	// appending, it creates new space
	s = append(s, "4")
	s = append(s, "5")

	fmt.Println("appended ", s)

	// copy
	c := make([]string, len(s))
	copy(c, s)

	fmt.Println("Copied array:", c)

	// slicing
	l := s[2:5]
	fmt.Println("sliced array:", l)

	// initalise
	t1 := []string{"1", "2", "3"}
	t2 := make([]string, len(t1))
	copy(t1, t2)

	fmt.Println(t1)

	// equals
	if slices.Equal(t1, t2) {
		fmt.Println("Both are equal")
	}

	// two dimensional array with different inner size

	t3D := make([][]int, 3)

	for i := range 3 {
		innerLen := i + 1
		t3D[i] = make([]int, innerLen)
		for j := range innerLen {
			t3D[i][j] = i + j
		}
	}
	fmt.Println("2d: ", t3D)
}
