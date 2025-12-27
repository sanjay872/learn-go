package examples

import "fmt"

func Channel() {
	message := make(chan string, 2) // args -> datatype of message, capacity of buffer

	go func() {
		message <- "ping"
		message <- "ping 2"
		message <- "ping 3"
	}()

	fmt.Println(<-message)
	fmt.Println(<-message)
	fmt.Println(<-message)
}
