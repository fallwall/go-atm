package main

import "fmt"

// functions that takes two ints
// and return their sum as an int
func plus(a int, b int) int {
	return a + b
}

// if all you have multiple consecutive parameters of same type
// omit the type name for the like-typed params up to the final one
func plusPlus(a, b, c int) int {
	return a + b + c
}

func main() {
	res := plus(1, 2)
	fmt.Println(res)

	res = plusPlus(1, 2, 3)
	fmt.Println(res)
}
