package examples

import "fmt"

func ping(pings chan<- string, msg string) { // in args, I have given message will be received
	pings <- msg
}

func pong(pings <-chan string, pongs chan<- string) { // in args, I will be getting the string out of pings and I will be giving string to pongs
	msg := <-pings
	pongs <- msg
}

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

	pings := make(chan string, 1)
	pongs := make(chan string, 1)

	ping(pings, "passed message")
	pong(pings, pongs)

	fmt.Println(<-pongs)
}
