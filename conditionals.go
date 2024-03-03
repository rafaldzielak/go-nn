package main

import "fmt"

func conditionals() {
	age := 45
	if age < 30 {
		fmt.Println("Age is less than 30")
	} else if age < 40 {
		fmt.Println("Age is less than 40")
	} else {
		fmt.Println("Age is more or equal to 40")
	}

	for index, value := range strs {
		if index == 1 {
			fmt.Println("continuing at pos", index)
			continue
		}
		if index > 2 {
			fmt.Println("Breaking at pos", index)
			break
		}
		fmt.Printf("the value at pos %v is %v\n", index, value)
	}
}
