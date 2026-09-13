package advanced

import (
	"context"
	"fmt"
	"log"
	"time"
)

func doWork(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Work cancelled:", ctx.Err())
			return
		default:
			fmt.Println("Working...")
		}
		time.Sleep(500 * time.Millisecond)
	}
}

func main() {
	ctx := context.Background()
	// ctx, cancel := context.WithTimeout(ctx, 2*time.Second) // timer of the context starts here. You have this specified time duration to use this context. After this time duration, the context will send a cancelation signal.

	ctx, cancel := context.WithCancel(ctx) // timer of the context starts here. You have this specified time duration to use this context. After this time duration, the context will send a cancelation signal.

	go func() {
		time.Sleep(2 * time.Second) // simulating a heavy task. Consider this a heavy time consuming operation
		cancel()                    // manually canceling only after the task is finished
	}()

	ctx = context.WithValue(ctx, "requestID", "hdsjf3234324234")
	ctx = context.WithValue(ctx, "name", "John")
	ctx = context.WithValue(ctx, "IP", "dsd.34.4332.34")
	ctx = context.WithValue(ctx, "OS", "Operating System")

	go doWork(ctx)

	time.Sleep(3 * time.Second)

	requestID := ctx.Value("requestID")
	if requestID != nil {
		fmt.Println("Request ID:", requestID)
	} else {
		fmt.Println("No request ID found.")
	}

	logWithContext(ctx, "This is a test log message")
}

func logWithContext(ctx context.Context, message string) {
	requestIDVal := ctx.Value("requestID")
	log.Printf("RequestID: %v - %v", requestIDVal, message)
}

// func checkEvenOdd(ctx context.Context, num int) string {
// 	select {
// 	case <-ctx.Done():
// 		return "Operation canceled"
// 	default:
// 		if num%2 == 0 {
// 			return fmt.Sprintf("%d is even", num)
// 		} else {
// 			return fmt.Sprintf("%d is odd", num)
// 		}
// 	}
// }

// func main() {
// 	ctx := context.TODO()

// 	result := checkEvenOdd(ctx, 5)
// 	fmt.Println("Result with context.TODO():", result)

// 	ctx = context.Background()
// 	ctx, cancel := context.WithTimeout(ctx, 1*time.Second)
// 	defer cancel()

// 	result = checkEvenOdd(ctx, 10)
// 	fmt.Println("Result from timeout context:", result)

// 	time.Sleep(3 * time.Second)
// 	result = checkEvenOdd(ctx, 15)
// 	fmt.Println("Result after timeout:", result)
// }

// ======= DIFFERENCE BETWEEN CONTEXT.TODO AND CONTEXT.BACKGROUND
// func main() {

// 	todoContext := context.TODO()
// 	contextBkg := context.Background()

// 	ctx := context.WithValue(todoContext, "name", "John")
// 	fmt.Println(ctx)
// 	fmt.Println(ctx.Value("name"))

// 	ctx1 := context.WithValue(contextBkg, "city", "New York")
// 	fmt.Println(ctx1)
// 	fmt.Println(ctx1.Value("city"))
// }
