package examples

import (
	"errors"
	"fmt"
)

type argError struct {
	val int
	msg string
}

func (arg *argError) Error() string {
	return fmt.Sprintf("%d - %s", arg.val, arg.msg)
}

func fv2(val int) (int, error) {
	if val == 42 {
		return -1, &argError{val, "can't work with it"}
	}
	return val + 3, nil
}

func CustomError() {
	_, e := fv2(42)
	var ae *argError
	if errors.As(e, &ae) {
		fmt.Println(ae.val)
		fmt.Println(ae.msg)
	} else {
		fmt.Println("Error doesn't match val")
	}
}
