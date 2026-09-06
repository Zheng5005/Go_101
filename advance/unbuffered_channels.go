package channels

import (
	"fmt"
	"time"
)

func unbuffered() {
	ch := make(chan int) //by default is an unbuffered channel, meaning a channel without an associated store

	go func()  {
		time.Sleep(3 * time.Second)

		fmt.Println(<-ch)
		fmt.Println("3 second Goroutine finished")
	}()

	ch <- 1

	// go func() {
	// 	time.Sleep(2 * time.Second)
	// 	fmt.Println("2 second Goroutine finished")
	// }()
	// go func() {
	// 	// ch <- 1
	// 	time.Sleep(3 * time.Second)
	// 	fmt.Println("3 second Goroutine finished")
	// }()
	// receiver := <-ch
	// fmt.Println(receiver)
	fmt.Println("End of program")
}
