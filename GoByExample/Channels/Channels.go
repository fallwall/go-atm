package main

import "fmt"

func main() {
	// new channel: make(chan val-type)
	messages := make(chan string)

	// send value into chanel
	// channel <-
	go func() {
		messages <- "ping"} ()
	

	// receive value from channel
	// <-channel
	msg := <-messages

	fmt.Println(msg)
}