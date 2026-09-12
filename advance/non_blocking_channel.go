package advance

import (
	"fmt"
	"time"
)

func non_blocking_channel() {
	// ch := make(chan int)

	// === NON BLOCKING RECEIVE OPERATION
	// select {
	// case msg := <-ch:
	// 	fmt.Println("Received:", msg)
	// default:
	// 	fmt.Println("No messages available.")
	// }

	// // === NON BLOCKING SEND OPERATION
	// select {
	// case ch <- 1:
	// 	fmt.Println("Sent message.")
	// default:
	// 	fmt.Println("Channel is not ready to receive.")
	// }

	// === NON BLOCKING OPERATION IN REAL TIME SYSTEMS

	data := make(chan int)
	quit := make(chan bool)

	go func()  {
		for {
			select {
			case d := <-data:
				fmt.Println("Data received:", d)
			case <-quit:
				fmt.Println("Stopping...")
			default:
				fmt.Println("Waiting for data...")
				time.Sleep(500 * time.Millisecond)
			}
		}
	}()

	for i := range 5 {
		data <- i
		time.Sleep(time.Second)
	}

	quit <- true
}
