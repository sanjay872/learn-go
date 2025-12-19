package examples

import "fmt"

func intSeq() func() int {
	i := 0

	return func() int {
		i += 1
		return i
	}
}

func Closures() {
	nextInt := intSeq()

	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	newNextInt := intSeq()
	fmt.Println(newNextInt())
}
