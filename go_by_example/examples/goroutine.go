package examples

import (
	"fmt"
	"time"
)

func myf(from string) {
	for i := range 3 {
		fmt.Println(from, ":", i)
	}
}

func GoRoutine() {

	myf("direct")

	go myf("goroutine")

	go func(from string) {
		fmt.Println(from)
	}("going")

	time.Sleep(time.Second)
	fmt.Println("Done")
}
