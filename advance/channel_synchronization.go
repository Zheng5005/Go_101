package advanced

import (
	"fmt"
	"time"
)

// func main() {

// 	done := make(chan int)

// 	go func() {
// 		fmt.Println("Working...")
// 		time.Sleep(2 * time.Second)
// 		done <- 0
// 	}()

// 	<-done
// 	fmt.Println("Finished.")

// }

// func main() {
// 	ch := make(chan int)

// 	go func() {
// 		ch <- 9 // Blocking until the value is received
// 		fmt.Println("Sent value")
// 	}()

// 	value := <-ch // Blocking until a value is sent
// 	fmt.Println(value)
// }

// // ========= SYNCHRONIZING MULTIPLE GOROUTINES AND ENSURING THAT ALL GOROUTINES ARE COMPLETE
// func main() {
// 	numGoroutines := 3
// 	done := make(chan int, 3)

// 	for i := range numGoroutines {
// 		go func(id int) {
// 			fmt.Printf("Goroutine %d working...\n", id)
// 			time.Sleep(time.Second)
// 			done <- id // SENDING SIGNAL OF COMPLETION
// 		}(i)
// 	}

// 	for range numGoroutines {
// 		<-done // Wait for each goroutine to finish, WAIT FOR ALL GOROUTINES TO SIGNAL COMPLETION
// 	}

// 	fmt.Println("All goroutines are complete")
// }

// ========== 	SYNCHRONIZING DATA EXCHANGE
func channel_sync() {

	data := make(chan string)

	go func() {
		for i := range 5 {
			data <- "hello " + string('0'+i)
			time.Sleep(100 * time.Millisecond)
		}
		close(data)
	}()
	// close(data) // Channel closed before Goroutine could send a value to the channel

	for value := range data {
		fmt.Println("Received value:", value, ":", time.Now())
	} // Loops over only on active channel, creates receiver each time and stops creating receiver (looping) once the channel is closed
}
