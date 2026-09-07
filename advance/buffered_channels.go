package advance

import (
	"fmt"
	"time"
)
// Buffered channels are initialized with a fixed-size capacity that allows sender to transmit multiple values without blocking, as long the buffer is not full

// func main() {
// 	// =========== BLOCKING ON RECEIVE ONLY IF THE BUFFER IS EMPTY
// 	ch := make(chan int, 2)

// 	go func() {
// 		time.Sleep(2 * time.Second)
// 		ch <- 1
// 		ch <- 2
// 	}()
// 	fmt.Println("Value: ", <-ch)
// 	fmt.Println("Value: ", <-ch)
// 	fmt.Println("End of program.")
// }

func buffered_channel() {
	// ================== BLOCKING ON SEND ONLY IF THE BUFFER IS FULL
	// make(chan Type, capacity)
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	fmt.Println("Reciving from buffer")

	go func() {
		fmt.Println("Goroutine 2 second timer started")
		time.Sleep(2 * time.Second)
		fmt.Println("Received:", <-ch) //ends <- starts
	}()

	// fmt.Println("Blocking starts")
	ch <- 3 // Blocks because the buffer is full
	// fmt.Println("Blocking ends")
	// fmt.Println("Received:", <-ch)
	// fmt.Println("Received:", <-ch)
}
