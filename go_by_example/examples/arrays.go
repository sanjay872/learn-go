package examples

import "fmt"

func Arrays() {
	var a [5]int // to store 5 values

	fmt.Println(a)

	a[4] = 100
	fmt.Println(len(a))

	// initialise

	b := [5]int{1, 2, 3, 4, 5}

	fmt.Println(b)

	var c [2][2]int

	for i := range 2 {
		for j := range 2 {
			fmt.Println(c[i][j])
		}
	}
}
