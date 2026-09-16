package mac

import (
	"fmt"
	"sync"
)

// go run -race <filename.go>
// The command above helps you to identify a race condition in your program

type counter struct {
	mu    sync.Mutex
	count int
}

func (c *counter) increment() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.count++
}

func (c *counter) getValue() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.count
}

func main() {

	var wg sync.WaitGroup
	counter := &counter{}
	numGoroutines := 10

	// wg.Add(numGoroutines)
	for range numGoroutines {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for range 1000 {
				counter.increment()
				// counter.count++
			}
		}()
	}

	wg.Wait()
	fmt.Printf("Final counter value: %d\n", counter.getValue())

}
