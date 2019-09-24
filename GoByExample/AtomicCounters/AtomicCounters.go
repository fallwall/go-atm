// sync/stomic package for atomic counters accessed by multiple goroutines
// managing states other than worker pools

package main

import (
	"sync"
	"sync/atomic"
	"fmt"
)

func main() {
	// unassigned interger to rep counter
	var ops int64
	// waitgroup will help us wait for all goroutines to finish
	var wg sync.WaitGroup

	// 50 goroutines, each increment the counter 1000x
	for i := 0; i < 50; i++ {
		wg.Add(1)

		go func() {
			for c := 0; c < 1000; c++ {

				//to atomically increment the counter
				// giving it the memory address of our ops counter with &
				atomic.AddInt64(&ops, 1)
			}
			wg.Done()
		}()
	}

	//wait until all goroutines are done
	wg.Wait()
	fmt.Println("ops:", ops)
}