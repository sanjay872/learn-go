package examples

import "fmt"

func ClosingChannels() {
	jobs := make(chan int, 5)
	done := make(chan bool)

	go func() {
		for {
			j, running := <-jobs
			if running {
				fmt.Println("Receiving job:", j)
			} else {
				fmt.Println("Received all Jobs!")
				done <- true
				return
			}
		}
	}()

	for j := 1; j <= 3; j++ {
		jobs <- j
		fmt.Println("Sending Job: ", j)
	}
	close(jobs)

	<-done // waiting for the jobs to complete.

	_, isPendng := <-jobs

	fmt.Println("Any job pending:", isPendng)
}
