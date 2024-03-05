package main

import "fmt"

func updateNamePointer(x *string) {
	*x = "dyrek"
}

func pointers() {
	fmt.Println("------POINTERS-------")
	name := "rafa"
	namePointer := &name
	fmt.Println("Memory address of name is", &name, "or", namePointer)
	fmt.Println("Value of name is", name, "or", *namePointer) // *pointer gets the value of the variable pointer points to

	updateNamePointer(namePointer)
	fmt.Println("Updated value of name is", name, "or", *namePointer)
}
