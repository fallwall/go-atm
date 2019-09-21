package main

import "fmt"

func main() {

	// to create an empty map
	// make(map[key-type]val-type)
	m := make(map[string]int)

	// set syntax: name[key] = val syntax
	m["k1"] = 7
	m["k2"] = 13

	fmt.Println("map:", m)

	// get value with name[key]
	v1 := m["k1"]
	fmt.Println("v1:", v1)

	// len -> key/value pair number
	fmt.Println("len:", len(m))

	// delete: to delete k2 from m
	delete(m, "k2")
	fmt.Println("map:", m)

	// ? check if a key is present in the map
	_, prs := m["k2"]
  fmt.Println("prs:", prs)

	// to initialize and declar in one line
	n := map[string]string{"roe": "cat", "simba": "dog"}
	fmt.Println("map:", n)

}
