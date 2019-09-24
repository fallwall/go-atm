package main

import "fmt"

func main() {
	jobs := make(chan int, 5)
	done := make(chan bool)

	go func() {
		for {
			// 2-value form of receive
			// more = false if jobs closed
			// & all values in the channel received
			// notify on done when we've worked all our jobs
			j, more := <-jobs
			if more {
				fmt.Println("received job", j)
			} else {
				fmt.Println("received all jobs")
				done <- true
				return
			}
		}
	}()

	// send 3 jobs to the worker over the jobs channel, then close it:
	for j := 1; j <=3; j++ {
		jobs <- j
		fmt.Println("sent job", j)
	}
	close(jobs)
	fmt.Println("sent all jobs")

	// await the worker , synch approach
	<-done
}