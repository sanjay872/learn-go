package examples

import "fmt"

func Add(a int, b int) int {
	return a + b
}

func Vals() (int, int) {
	return 1, 2
}

func N_Args(nums ...int) {
	var sum int = 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println(sum)
}

func DemoFunctions() {
	fmt.Println(Add(1, 4))
	nums := []int{1, 3, 4, 3}
	N_Args(nums...)
}
