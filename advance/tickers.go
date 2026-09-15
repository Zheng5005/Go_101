package advance

import (
	"fmt"
	"time"
)

func tickers() {

	ticker := time.NewTicker(time.Second)
	stop := time.After(5 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case tick := <-ticker.C:
			fmt.Println("Tick at:", tick)
		case <-stop:
			fmt.Println("Stopping ticker.")
			return
		}
	}
}

// ========= SCHEDULING LOGGING, PERIODIC TASKS, POLLING FOR UPDATES
// func periodicTask() {
// 	fmt.Println("Performing periodic task at:", time.Now())
// }

// func main() {
// 	ticker := time.NewTicker(time.Second)
// 	defer ticker.Stop()

// 	for {
// 		select {
// 		case <-ticker.C:
// 			periodicTask()
// 		}
// 	}
// }

// func main() {

// 	ticker := time.NewTicker(2 * time.Second)
// 	defer ticker.Stop()
// 	// for tick := range ticker.C {
// 	// 	fmt.Println("Tick at:", tick)
// 	// }
// 	i := 1
// 	for range 5 {
// 		i *= 2
// 		fmt.Println(i)
// 	}

// 	// for tick := range ticker.C {
// 	// 	i *= 2
// 	// 	fmt.Println(tick)
// 	// }
// }
