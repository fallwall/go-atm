package main

import "fmt"

func main() {
	// make a channel of strings buffering up to 2 values
	messages := make(chan string, 2)

	// because this channel is buffered, 
	// we can send these values into the channel wo a corresponding concurrent receive
	messages <- "buffered"
	messages <- "channel"

	fmt.Println(<-messages)
	fmt.Println(<-messages)
}