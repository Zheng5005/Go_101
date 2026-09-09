package advanced

import "fmt"

func main() {

	ch := make(chan int)
	producer(ch)
	consumer(ch)
}
// Send only channel
func producer(ch chan<- int)  {
	go func() {
		for i := range 5 {
			ch <- 1
		}
		close(ch)
	}()
}

// Receive only channel
func consumer(ch <-chan int)  {
	for value := range ch {
		fmt.Println("Received: ", value)
	}
}
