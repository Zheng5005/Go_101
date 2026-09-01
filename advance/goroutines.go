package advance

import (
	"fmt"
	"time"
)

// Goroutines are just functions that leave the main thread and run in the background and come back to join the main thread once the functions are finished/ready to return any value
// Goroutine don't stop the program flow and are non blocking
// Similar to async function in JS

func main() {
	var err error

	fmt.Println("Beginning program.")
	go sayHello()
	fmt.Println("After sayHello funtion.")

	go func() {
		err = doWork()
	}()

	// err = go doWork() // This is not accepted
	go printNumbers()
	go printLetters()

	time.Sleep(2 * time.Second)

	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Work completed successfully")
	}
}

func sayHello() {
	time.Sleep(1 * time.Second)
	fmt.Println("Hello from Goroutine")
}

func printNumbers() {
	for i := 0; i < 5; i++ {
		fmt.Println("Number: ", i, time.Now())
		time.Sleep(100 * time.Millisecond)
	}
}

func printLetters() {
	for _, letter := range "abcde" {
		fmt.Println(string(letter), time.Now())
		time.Sleep(200 * time.Millisecond)
	}
}

func doWork() error {
	// Simulate work
	time.Sleep(1 * time.Second)

	return fmt.Errorf("an error occured in doWork.")
}
