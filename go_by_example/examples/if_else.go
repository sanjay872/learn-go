package examples

import "fmt"

func IfElse() {
	i := 1

	if i == 1 {
		fmt.Println("I am one!")
	}

	if i >= 2 {
		fmt.Println("I am two")
	} else if i == 1 {
		fmt.Println("I am One")
	} else {
		fmt.Println("I am less than or equal to zero!")
	}
}
