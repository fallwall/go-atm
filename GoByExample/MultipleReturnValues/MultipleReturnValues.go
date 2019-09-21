package main

import "fmt"

// (int, int) means this function return 2 ints
func vals() (int, int) {
	return 3, 7
}

func main() {
	// multiple assignment
	a, b := vals()
	fmt.Println(a)
	fmt.Println(b)

	// blank identifier _
	_, c  := vals()
	fmt.Println(c)
}