package main

import (
	"fmt"
	"time"
)

// done channel
// to notify another goroutine that this's done
func worker(done chan bool) {
	fmt.Println("working")
	time.Sleep(time.Second)
	fmt.Println("done")

	// send a value to notify when we are done
	done <- true
}

func main() {

	done := make(chan bool, 1)
	// start a worker goroutine, giving it the channel to notify on
	go worker(done)

	// block until we receive a notification from the worker on that channel
	<-done
}