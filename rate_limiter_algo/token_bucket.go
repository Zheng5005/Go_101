package ratelimiteralgo

import (
	"fmt"
	"time"
)

type RateLimiter struct {
	tokens     chan struct{}
	refillTime time.Duration
}

func NewRateLimiter(rateLimit int, refillTime time.Duration) *RateLimiter {
	rl := &RateLimiter{
		tokens:     make(chan struct{}, rateLimit),
		refillTime: refillTime,
	}
	for range rateLimit {
		rl.tokens <- struct{}{}
	}
	go rl.startRefill()
	return rl
}

func (rl *RateLimiter) startRefill() {
	ticker := time.NewTicker(rl.refillTime)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			select {
			case rl.tokens <- struct{}{}:
			default:
			}
		}
	}
}

func (rl *RateLimiter) allow() bool {
	select {
	case <-rl.tokens:
		return true
	default:
		return false
	}
}

func main() {

	rateLimiter := NewRateLimiter(5, time.Second)

	for range 10 {
		if rateLimiter.allow() {
			fmt.Println("Request allowed")
		} else {
			fmt.Println("Request denied")
		}
		time.Sleep(300 * time.Millisecond)
	}
}

//1 0 ms		First Request Allowed	5 tokens left
//2 200 ms		Second Request Allowed	4 tokens left
//3 400 ms		Third Request Allowed	3 tokens left
//4 600 ms		Fourth Request Allowed	2 tokens left
//5 800 ms		Fifth Request Allowed	1 tokens left
//6 1000 ms		Sixth Request Allowed	(not 0) 1 tokens left the startRefill function executes and adds one more token
//7 1200 ms		Seven Request Denied	tokens left
//8 1400 ms		Eight Request Denied	tokens left
//9 1600 ms		Ninth Request Denied	tokens left
//10 1800 ms		Tenth Request Denied
