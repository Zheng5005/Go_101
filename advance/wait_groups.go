package advanced

import (
	"fmt"
	"sync"
	"time"
)

// CONSTRUCTION EXAMPLE
type Worker struct {
	ID   int
	Task string
}

// PerformTask simulates a worker performing a task
func (w *Worker) PerformTask(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("WorkerID %d started %s\n", w.ID, w.Task)
	time.Sleep(time.Second)
	fmt.Printf("WorkerID %d finished %s\n", w.ID, w.Task)
}

func wait_groups() {
	var wg sync.WaitGroup

	// Define tasks to be performed by workers
	tasks := []string{"digging", "laying bricks", "painting"}

	for i, task := range tasks {
		worker := Worker{ID: i + 1, Task: task}
		wg.Add(1)
		go worker.PerformTask(&wg)
	}

	// Wait for all workers to finish
	wg.Wait()

	// Construction is finished
	fmt.Println("Contruction finished")
}

// // EXAMPLE WITH CHANNELS
// func worker(id int, tasks <-chan int, results chan<- int, wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	fmt.Printf("WorkerID %d starting.\n", id)
// 	time.Sleep(time.Second) // simulate some work
// 	for task := range tasks {
// 		results <- task * 2
// 	}
// 	fmt.Printf("WorkerID %d finished\n", id)
// }

// func main() {
// 	var wg sync.WaitGroup
// 	numWorkers := 3
// 	numJobs := 5
// 	results := make(chan int, numJobs)
// 	tasks := make(chan int, numJobs)

// 	wg.Add(numWorkers)

// 	for i := range numWorkers {
// 		go worker(i+1, tasks, results, &wg)
// 	}

// 	for i := range numJobs {
// 		tasks <- i + 1
// 	}
// 	close(tasks)

// 	go func() {
// 		wg.Wait()
// 		close(results)
// 	}()

// 	for result := range results {
// 		fmt.Println("Result:", result)
// 	}
// }

// // ============ BASIC EXAMPLE WITHOUT USING CHANNELS
// func worker(id int, wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	// wg.Add(1)  /// WRONG PRACTICE
// 	fmt.Printf("Worker %d starting\n", id)
// 	time.Sleep(time.Second) // simulate some time spent on processing the task
// 	fmt.Printf("Worker %d finished\n", id)
// }

// func main() {
// 	var wg sync.WaitGroup
// 	numWorkers := 3

// 	wg.Add(numWorkers) // THIS IS THE CORRECT WAY OF ADDING COUNTER TO WAIT GROUPS

// 	// Launch workers
// 	for i := range numWorkers {
// 		go worker(i, &wg)
// 	}

// 	wg.Wait() // blocking mechanism
// 	fmt.Println("All workers finished")
// }
