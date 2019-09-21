// Variadic functions can be called with any number of trailing args.package VariadicFunctions
// i.e. fmt.Println

package main

import "fmt"

// take an arbitrary nums of ints as args
func sum(nums ...int) {
	fmt.Print(nums, " ")
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}

func main() {
	sum(1, 2)
	sum(1, 2, 3)

	// already have multiple args in a slice
	nums := []int{1,2,3,4}
	sum(nums...)
}