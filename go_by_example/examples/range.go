package examples

import "fmt"

func Range() {
	nums := []int{1, 2, 3, 4}

	for i, num := range nums {
		fmt.Println("Index is ", i, " and value is ", num)
	}

	mymap := map[int]int{1: 1, 2: 2, 3: 3}

	for k, v := range mymap {
		fmt.Println("key: ", k, " and Value: ", v)
	}

	for i, c := range "sanjay" {
		fmt.Printf("index %d and character %c \n", i, c)
	}
}
