package main

import "fmt"

func main() {
	// for loop - single condition
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	// for loop - ini/cond/after
	for j := 7; j <= 9; j ++ {
		fmt.Println(j)
	}

	// for wo condition: continut until nreak or return
	for {
		fmt.Println("loooop")
		break
	}

	// you can also continue to the enxt iteration
	for n := 0; n <=5; n++ {
		if n%2 ==0 {
			continue
		}
		fmt.Println(n)
	}
	
}