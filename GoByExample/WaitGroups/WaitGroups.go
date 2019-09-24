package main

import (
	"fmt"
	"sync"
	"time"
)

// a WaitGroup must be passed to functions by pointer

func worker(id int, wg *sync.WaitGroup) {
	fmt.Printf("Worker %d started\n", id)

	time.Sleep(time.Second)
	fmt.Printf("Worker %d done\n", id)

	wg.Done()
}

func main() {
	// wait for all the goroutines launced here to finish
	var wg sync.WaitGroup

	// launch several goroutines and increment the WaitGroup counter for each
	for i:= 1; i<=5; i++ {
		wg.Add(1)
		go worker(i, &wg)
	}

	// clock until counter goes back to 0
	wg.Wait()
}