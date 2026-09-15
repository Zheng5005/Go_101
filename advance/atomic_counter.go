package advanced

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type AtomicCounter struct {
	count int64
}

func (ac *AtomicCounter) increment() {
	atomic.AddInt64(&ac.count, 1)
}

func (ac *AtomicCounter) getValue() int64 {
	return atomic.LoadInt64(&ac.count)
}

func atomic_counter() {
	var wg sync.WaitGroup
	numGoroutines := 10
	counter := &AtomicCounter{}
	// value := 0

	for range numGoroutines {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for range 1000 {
				counter.increment()
				// value++
			}
		}()
	}

	wg.Wait()
	fmt.Printf("Final counter value: %d\n", counter.getValue())
	// fmt.Printf("Final counter value: %d\n", value)
}
