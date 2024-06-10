package pkg

import (
	"fmt"
	"time"
)

func sendData(ch chan<- int) {
	for i := 1; i <= 5; i++ {
		fmt.Println("Sending:", i)
		ch <- i // Send value into the channel
		time.Sleep(500 * time.Millisecond) // Simulate some processing time
	}
	close(ch) // Close the channel to indicate that no more data will be sent
}

func receiveData(ch <-chan int) {
	for {
		value, ok := <-ch // Receive value from the channel
		if !ok {
			fmt.Println("Channel closed. Exiting receiveData goroutine.")
			return
		}
		fmt.Println("Received:", value)
	}
}

func RunChannel() {
	ch := make(chan int, 3)

	go sendData(ch)

	go receiveData(ch)

	time.Sleep(3 * time.Second)
}
