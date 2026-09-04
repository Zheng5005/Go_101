package channels

import (
	"fmt"
	"time"
)

func channels() {
	//variable := make(chan type),
	//Operator: <-
	// Receiver: receiverchan <- value
	// Sender: <- senderchan

	// Channels don't work on a function, it has to be in a goroutine
	// channels are blocking
	greeting := make(chan string)
	greetString := "Hello"

	go func()  {
		greeting <- greetString // blocking because it is continuously trying to receive values, it is ready to receive continuous flow of data.
		greeting <- "World"
		for _, e := range "abcde" {
			greeting <- "Alphabet: " + string(e)
		}
	}()

	// go func() {
	// 	receiver := <-greeting
	// 	fmt.Println(receiver)
	// 	receiver = <-greeting
	// 	fmt.Println(receiver)
	// }()

	receiver := <-greeting
	fmt.Println(receiver)
	receiver = <-greeting
	fmt.Println(receiver)

	for range 5 {
		rcvr := <-greeting
		fmt.Println(rcvr)
	}

	time.Sleep(1 * time.Second)
	fmt.Println("End of program.")
}
