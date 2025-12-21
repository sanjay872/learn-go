package examples

import "fmt"

func ZeroVal(ival int) {
	ival = 0
}

func ZeroPtr(iptr *int) {
	*iptr = 0
}

func Pointers() {
	i := 1

	fmt.Println("Intital value: ", i)

	ZeroVal(i)
	fmt.Println("Value after ZeroValue function: ", i)

	ZeroPtr(&i)
	fmt.Println("Value after ZeroPtr function: ", i)
}
