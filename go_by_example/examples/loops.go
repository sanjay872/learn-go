package examples

import "fmt"

func Loops() {

	// while loop

	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	// for loop

	for i := 1; i <= 3; i++ {
		fmt.Println(i)
	}

	for i := range 3 {
		fmt.Println(i)
	}

	for {
		fmt.Println("Infinite loop!")
		break
	}

	for i := 1; i <= 10; i++ {
		if i%2 != 0 {
			continue
		}
		fmt.Println(i)
	}
}
