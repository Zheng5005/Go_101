package ratelimiteralgo

import (
	"fmt"
	"sync"
	"time"
)

type RateLimiterWindow struct {
	mu        sync.Mutex
	count     int
	limit     int
	window    time.Duration
	resetTime time.Time
}

func NewRateLimiterWindow(limit int, window time.Duration) *RateLimiterWindow {
	return &RateLimiterWindow{
		limit:  limit,
		window: window,
	}
}

func (rl *RateLimiterWindow) Allow() bool {
	rl.mu.Lock()
	defer rl.mu.Unlock()

	now := time.Now()
	if now.After(rl.resetTime) {
		rl.resetTime = now.Add(rl.window)
		rl.count = 0
	}

	if rl.count < rl.limit {
		rl.count++
		return true
	}
	return false
}

func main2() {
	var wg sync.WaitGroup
	rateLimiter := NewRateLimiterWindow(3, 1*time.Second)

	for range 10 {
		wg.Add(1)
		go func() {
			if rateLimiter.Allow() {
				fmt.Println("Request allowed.")
			} else {
				fmt.Println("Request denied.")
			}
			wg.Done()
		}()
		// time.Sleep(200 * time.Millisecond)
	}
	wg.Wait()

}
