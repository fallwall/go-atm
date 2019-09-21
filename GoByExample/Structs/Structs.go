// Go's structs are typed collections of fields
// useful for grouping data together to form records
package main

import "fmt"

// this person struct type has name and age fields
type person struct {
	name string
	age  int
}

// NewPerson constrcuts a new person struct w/ given name
func NewPerson(name string) *person {
	// You can safely return a pointer to local variable
	// ? as a local variable will survive the scope of the function
	p := person{name: name}
	p.age = 42
	return &p
}

func main() {
	// this creates a new struct
	fmt.Println(person{"Simba", 1})
	// you can name the fields when init
	fmt.Println(person{name: "Roe", age: 2})
	// ommitted filed will be zdro-valued
	fmt.Println(person{name: "Nala"})
	// & pre-fix yields a pointer to the struct
	fmt.Println(&person{name: "Sarah", age: 26})

	// idiomatic to encapsulate new struct creation in constructor functions
	fmt.Println(NewPerson("John"))

	// dot syntax to access struct fields
	s := person{name: "Sabrina Spellman", age: 16}
	fmt.Println(s.name)

	// dots w/ struct pointers
	// ? the pointers are automatically dereferenced
	sp := &s
	fmt.Println(sp.age)

	//structs are mutable
	sp.age = 51
	fmt.Println(sp.age)
}
