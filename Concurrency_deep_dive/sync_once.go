package mac

import (
	"fmt"
	"sync"
)

var once sync.Once

func initialize() {
	fmt.Println("This will not be repeated no matter how mahy times you call this function using once.Do.")
}

func main_sync_once() {

	var wg sync.WaitGroup
	for i := range 5 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Println("Goroutine # ", i)
			once.Do(initialize)
			// initialize()
		}()
	}
	wg.Wait()
}
