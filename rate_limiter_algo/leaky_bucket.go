package ratelimiteralgo

import (
	"fmt"
	"sync"
	"time"
)

type LeakyBucket struct {
	capacity int
	leakRate time.Duration
	tokens   int
	lastLeak time.Time
	mu       sync.Mutex
}

func NewLeakyBucket(capacity int, leakRate time.Duration) *LeakyBucket {
	return &LeakyBucket{
		capacity: capacity,
		leakRate: leakRate,
		tokens:   capacity,
		lastLeak: time.Now(),
	}
}

func (lb *LeakyBucket) Allow() bool {
	lb.mu.Lock()
	defer lb.mu.Unlock()

	now := time.Now()
	elapsedTime := now.Sub(lb.lastLeak)
	tokensToAdd := int(elapsedTime / lb.leakRate) // (0.2 / 0.5) result is 0.4 int value is 0
	lb.tokens += tokensToAdd

	if lb.tokens > lb.capacity {
		lb.tokens = lb.capacity
	}

	lb.lastLeak = lb.lastLeak.Add(time.Duration(tokensToAdd) * lb.leakRate)
	// lb.lastLeak = lb.lastLeak.Add(elapsedTime) //WRONG APPROACH
	// elapsedTime = 1.3 seconds
	// initial lastLeak = 0
	// tokensToAdd = int(1.3/.5) = int(2.6) = 2 tokens
	// lb.lastLeak = lb.lastLeak + (time.Duration(tokensToAdd) * lb.leakRate) 0 + (2 * 0.5) = 1 second
	// lb.lastLeak = lb.lastLeak + elapsed time 0 + 1.3

	fmt.Printf("Tokens added %d, Tokens subtracted %d, Total tokens: %d\n", tokensToAdd, 1, lb.tokens)
	fmt.Printf("Last leak time: %v\n", lb.lastLeak)
	if lb.tokens > 0 {
		lb.tokens--
		return true
	}
	return false
}

// ========= NON CONCURRENT REQUESTS
// func main() {
// 	leakyBucket := NewLeakyBucket(5, 500*time.Millisecond)

// 	for i := 0; i < 10; i++ {
// 		if leakyBucket.Allow() {
// 			fmt.Println("Current time", time.Now())
// 			fmt.Println("Request allowed")
// 		} else {
// 			fmt.Println("Current time", time.Now())
// 			fmt.Println("Request denied")
// 		}
// 		time.Sleep(200 * time.Millisecond)
// 	}
// }

// ========= CONCURRENT REQUESTS
func main3() {
	leakyBucketInst := NewLeakyBucket(5, 500*time.Millisecond)
	var wg sync.WaitGroup

	for range 10 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			if leakyBucketInst.Allow() {
				fmt.Println("Current time:", time.Now())
				fmt.Println("Request allowed.")
			} else {
				fmt.Println("Current time:", time.Now())
				fmt.Println("XXX----Request denied.")
			}
			time.Sleep(200 * time.Millisecond)
		}()
	}
	time.Sleep(500 * time.Millisecond)
	for range 10 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			if leakyBucketInst.Allow() {
				fmt.Println("Current time:", time.Now())
				fmt.Println("Request allowed.")
			} else {
				fmt.Println("Current time:", time.Now())
				fmt.Println("XXX----Request denied.")
			}
			time.Sleep(200 * time.Millisecond)
		}()
	}

	wg.Wait()
}
