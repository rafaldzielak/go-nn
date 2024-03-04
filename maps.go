package main

import "fmt"

// Maps are collections of key-value pairs
// All of the keys must have the same type and all of the values must have the same type
func maps() {
	menu := map[string]float64{
		"soup":           4.99,
		"pie":            7.99,
		"salad":          6.99,
		"toffee pudding": 3.55,
	}

	fmt.Println(menu)
	fmt.Println(menu["pie"])
	for key, value := range menu {
		fmt.Println(key, "-", value)
	}

	// ints as key type
	phonebook := map[int]string{
		123: "Rafa",
		456: "Dyrek",
	}
	fmt.Println(phonebook)
	fmt.Println(phonebook[123])

	phonebook[123] = "Rafael"
	fmt.Println(phonebook[123])
}
