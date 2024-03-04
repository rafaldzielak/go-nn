package main

func updateName(x string) {
	x = "dyrek"
}

func updateMenu(y map[string]float64) {
	y["pie"] = 999.99
}

func passByValue() {
	// strings, ints, bools, floats, arrays, structs are passed by value -> non-pointer values
	name := "rafa"
	updateName(name)
	println(name) // rafa

	// slices, maps and functions are passed by reference -> pointer wrapper values
	updateMenu(menu)
	println(menu["pie"]) // 999.99
}
