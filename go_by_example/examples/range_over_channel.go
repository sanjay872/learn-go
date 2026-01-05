package examples

import "fmt"

func RangeOverChannel() {
	str := make(chan string, 3)

	str <- "hello"
	str <- "world"

	close(str)

	for s := range str {
		fmt.Println(s)
	}
}
