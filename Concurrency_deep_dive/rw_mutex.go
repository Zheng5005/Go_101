package mac

import (
	"fmt"
	"sync"
	"time"
)

var (
	rwmu    sync.RWMutex
	counter_rwMutex int
)

func readCounter(wg *sync.WaitGroup) {
	defer wg.Done()
	rwmu.RLock()
	fmt.Println("Read Counter:", counter_rwMutex)
	rwmu.RUnlock()
}

func writeCounter(wg *sync.WaitGroup, value int) {
	defer wg.Done()
	rwmu.Lock()
	counter_rwMutex = value
	fmt.Printf("Written value %d for counter.\n", value)
	rwmu.Unlock()
}

func RWMutex() {
	var wg sync.WaitGroup
	for range 5 {
		wg.Add(1)
		go readCounter(&wg)
	}

	wg.Add(1)
	time.Sleep(time.Second)
	go writeCounter(&wg, 18)

	wg.Wait()
}
