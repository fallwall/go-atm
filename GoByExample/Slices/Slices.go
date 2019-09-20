package main

import "fmt"

func main() {

	// make a slice of strings of length 3 (initially zero-valued).
	s := make([]string, 3)
	fmt.Println("emp: ", s)

	// set and get like arrays
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set: ", s)
	fmt.Println("get: ", s[2])
	fmt.Println("len: ", len(s))

	// built-in append
	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("apd: ", s)

	// to copy
	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("cpy: ", c)
	
	// to slice
	// slice[low:high]. 
	// this gets s[2], s[3], and s[4].
	l := s[2:5]
	fmt.Println("sl1: ", l)

	// this to slice up to s[5] (exclude [5])
	l = s[:5]
	fmt.Println("sl2: ", l)

	// this to slife up from s[2] (include [2])
	l = s[2:]
	fmt.Println("sl3: ", l)

	// declare and initial one liner
	t := []string{"g", "h", "i"}
	fmt.Println("dcl: ", t)

	// Slices can be composed into multi-dimensional data structures. 
	// The length of the inner slices can vary, unlike with multi-dimensional arrays.
	twoD := make([][]int, 3)
    for i := 0; i < 3; i++ {
        innerLen := i + 1
        twoD[i] = make([]int, innerLen)
        for j := 0; j < innerLen; j++ {
            twoD[i][j] = i + j
        }
    }
    fmt.Println("2d: ", twoD)
}
