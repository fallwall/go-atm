package main

import "fmt"

func main() {

	// use range to sum the numbers in a slice
	// array workss like this too
	// _ to ignore index
	nums := []int{2,3,4}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum)

	// when we need index too
	for i, num := range nums {
		if num == 3 {
			fmt.Println("index:", i)
		}
	}

	// range on map iterates over key/value pairs
	kvs := map[string]string{"a":"apple","b":"banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

	// range on map, can iiterate over just keys
	for k := range kvs {
		fmt.Println("key:", k)
	}

	// range on string iterates over unicode code points
	for i, c := range "stuff" {
		fmt.Println(i,c)
	}
}