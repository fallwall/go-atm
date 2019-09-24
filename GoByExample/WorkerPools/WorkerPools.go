package main

import (
	"fmt"
	"time"
)

// receive work on the jobs channel
// sleep 1 sec = or any other task
// send the corresponding results on results
func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("wroker", id, "started job", j)
		time.Sleep(time.Second)
		fmt.Println("worker", id, "finished job", j)
		results <- j * 2
	}
}

func main() {

	jobs := make(chan int, 100)
	results := make(chan int, 100)

	// 3 workers, initally blocked
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	// 5 jobs then close
	for j := 1; j <= 5; j++ {
		jobs <- j
	}
	close(jobs)

	// collect resuults
	for a := 1; a <= 5; a++ {
		<-results
	}
}
