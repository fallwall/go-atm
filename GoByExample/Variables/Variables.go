package main

import  "fmt"

func main() {

	// var decalires 1 or more variables
	var a = "initial"
	fmt.Println(a)

	var b, c int = 1, 2
	fmt.Println(b, c)

	// Go will infer the type of initialzied vars
	var d = true
	fmt.Println(d)

	// Variables declared without a corresponding initialization are zero-valued
	// For example, the zero value for an int is 0.
	var e int
	fmt.Println(e)

	// The := syntax is shorthand for declaring and initializing a variable, 
	// e.g. for var f string = "apple" in this case.
	f := "apple"
	fmt.Println(f)

}