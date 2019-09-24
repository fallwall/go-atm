package main

import "fmt"

func main() {

	// iterate over 2 values in queue channel
	queue := make(chan string, 2)
	queue <- "one"
	queue <- "two"
	close(queue)

	for elem := range queue {
		fmt.Println(elem)
	}
}
//  close a non-empty channel but still have the remaining values be received