package examples

import "fmt"

func fib(n int) int {
	if n <= 1 {
		return n
	}

	return n * fib(n-1)
}

func Recursion() {
	fmt.Println(fib(5))
}
