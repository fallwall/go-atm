// timers represent a single event in the future
// time.NewTimer(2 * time.Second)
// time.NewTimer(2 * time.Second).Stop()
package main

import (
	"fmt"
	"time"
)

func main() {

	timer1 := time.NewTimer(2 * time.Second)
	
	// below blocks on the timer's channel C until
	// it sends a value indicating that the timer expired
	<-timer1.C
	fmt.Println("Timer 1 expired")

	// timer instead of sleep because cancellable before expiration
	timer2 := time.NewTimer(time.Second)
	go func() {
		<-timer2.C
		fmt.Println("Timer 2 expired")
	}()
	stop2 := timer2.Stop()
	if stop2 {
		fmt.Println("Timer 2 stopped")
	}
}