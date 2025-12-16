package examples

import (
	"fmt"
	"time"
)

func Switches() {
	i := 1

	switch i {
	case 1:
		fmt.Println("I am one")
	case 2:
		fmt.Println("I am two")
	default:
		fmt.Println("I am not in switch!")
	}

	switch time.Now().Weekday() {
	case time.Monday:
		fmt.Print("Today is monday!")
	case time.Tuesday:
		fmt.Print("Today is Tuesday!")
	case time.Wednesday:
		fmt.Print("Today is Wednesday!")
	}

	t := time.Now()

	switch {
	case t.Hour() < 12:
		fmt.Println("It is before noon")
	default:
		fmt.Println("It is after noon")
	}

	whatIam := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("Its boolean")
		case int:
			fmt.Println("It is Integer")
		default:
			fmt.Println("Something else of type", t)
		}
	}

	whatIam(true)
	whatIam("me")
	whatIam(1)
}
