package examples

import (
	"fmt"
	"time"
)

func TimerAndTicker() {
	timer1 := time.NewTimer(2 * time.Second)

	<-timer1.C

	fmt.Println("Time 1 completed")

	timer2 := time.NewTimer(2 * time.Second)

	go func() {
		<-timer2.C
		fmt.Println("Time 2 completed")
	}()

	stopped := timer2.Stop()

	if stopped {
		fmt.Println("Timer 2 stopped!")
	}

	ticker1 := time.NewTicker(600 * time.Millisecond)
	done := make(chan bool)

	go func() {
		for {
			select {
			case <-done:
				return
			case t := <-ticker1.C:
				fmt.Println("Ticking at: ", t)
			}
		}
	}()

	time.Sleep(1600 * time.Millisecond)
	ticker1.Stop()
	done <- true
	fmt.Println("Ticker Completed!")
}
