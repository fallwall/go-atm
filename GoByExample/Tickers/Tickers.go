// tickers are for when you want to do something repeatedly at regular intervals. 

package main

import (
	"fmt"
	"time"
)

func main() {

	// similar to Timer syntax
	ticker := time.NewTicker(500 * time.Millisecond)
	done := make(chan bool)

	go func() {
		for {
			select {
			case <-done:
				return
			case t:= <-ticker.C:
				fmt.Println("Tick at", t)
			}
		}
	}()

	// ticker can be stopped
	// after stopping, can no longer receive value from it
	// unlike timers
	time.Sleep(1600 * time.Millisecond)
	ticker.Stop()
	done <- true
	fmt.Println("Ticker stopped")

}