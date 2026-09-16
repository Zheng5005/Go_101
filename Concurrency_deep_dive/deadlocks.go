package mac

import (
	"fmt"
	"sync"
	"time"
)

func deadlock() {

	var mu1, mu2 sync.Mutex

	go func() {
		mu1.Lock()
		fmt.Println("Goroutine 1 locked mu1")
		time.Sleep(time.Second)
		mu2.Lock()
		fmt.Println("Goroutine 1 locked mu2")
		mu1.Unlock()
		mu2.Unlock()
		fmt.Println("Goroutine1 finished")
	}()

	go func() {
		mu1.Lock()
		fmt.Println("Goroutine 2 locked mu2")
		time.Sleep(time.Second)
		mu2.Lock()
		fmt.Println("Goroutine 2 locked mu1")
		mu2.Unlock()
		mu1.Unlock()
		fmt.Println("Goroutine2 finished")
	}()

	time.Sleep(3 * time.Second)
	fmt.Println("Main function completed")
	// select {}
}
