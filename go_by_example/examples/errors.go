package examples

import (
	"errors"
	"fmt"
)

func f(arg int) (int, error) {
	if arg == 42 {
		return -1, errors.New("can't work with 42")
	}
	return arg + 3, nil
}

var Error1 = errors.New("this is my error 1")
var Error2 = errors.New("this is my error 2")

func makeError(arg int) error {
	switch arg {
	case 1:
		return Error1
	case 2:
		return Error2
	case 3:
		return fmt.Errorf("error 3: %w", Error2)
	}
	return nil
}

func Errors() {

	if i, e := f(42); e != nil {
		fmt.Println("failed for val with error: ", e)
	} else {
		fmt.Println("Sucess with value: ", i)
	}

	for i := range 4 {
		e := makeError(i)
		if errors.Is(e, Error1) {
			fmt.Println("It is error 1")
		} else if errors.Is(e, Error2) {
			fmt.Println("It is error 2")
		} else {
			fmt.Println(e)
		}
	}

}
