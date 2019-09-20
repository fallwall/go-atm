package main

import (
	"fmt"
	"math"
)

// const declares a constant value
// a constant statment can appear anywhere a var can
const s string = "constant"

func main() {
	fmt.Println(s)

	const n = 500000000

	// constant expressions perform arithmetic w/ arbitrary precison
	const d = 3e20 / n
	fmt.Println(d)

	// a nmeric constant has no type until it's given one
	// such as by an explicit conversion

	fmt.Println(int64(d))

	// a num can be given a type by using it in a context that requires one
	// such as a variable assignment or function call
	fmt.Println(math.Sin(n))
}
