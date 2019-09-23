// a goroutine is a lightweight thread of execution

package main

import "fmt"

func f(from string) {
	for i := 0; i <3; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {
	// usual way to call a function
	// running it synchronously
	f("direct")

	// to invoke this function in a goroutine
	// executing concurrently with above
	go f("goroutine") 

	// go on anonymous function ok
	go func(msg string) {
		fmt.Println(msg)
	}("going")

	// Scanln require we press a key before exit
	fmt.Scanln()
	fmt.Println("this is the end")
}