package main

import "fmt"

func main() {
	messages := make(chan string)
	signals := make(chan bool)

	// non-blocking receiver
	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	default:
		fmt.Println("no message received")
	}

	// non-blocking sender
	msg := "hi"
	select {
	case messages <- msg:
		fmt.Println("sent message", msg)
	default:
		fmt.Println("mo message sent")
	}

	// multi-way non-blocking select
	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	case sig := <-signals:
		fmt.Println("received signals", sig)
	default:
		fmt.Println("nothing to see here")
	}
}