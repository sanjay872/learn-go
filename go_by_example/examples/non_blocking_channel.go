package examples

import "fmt"

func NonBlockingChannel() {

	message := make(chan string, 1)
	signal := make(chan bool)

	select {
	case res := <-message:
		fmt.Println(res)
	default:
		fmt.Println("No message received!")
	}

	msg := "hello"

	select {
	case message <- msg: // message is sent but no variable to check it
		fmt.Println("message received")
	default:
		fmt.Println("No message received")
	}

	select {
	case res := <-message: // the sent message is seen
		fmt.Println(res)
	case sig := <-signal:
		fmt.Println(sig)
	default:
		fmt.Println("Nothing received!")
	}

}
