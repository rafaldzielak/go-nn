package main

import "fmt"

func types() {
	// functions that are exported are with Capital Letter
	fmt.Println("Hi")

	// strings, double ""
	var nameOne string = "mario"
	var nameTwo = "luigi"
	var nameThree string

	nameOne = "beach"
	nameThree = "peach"

	nameFour := "Rafa"
	fmt.Println(nameOne, nameTwo, nameThree, nameFour)

	// ints:
	var ageOne int = 20
	var ageTwo = 30
	ageThree := 40
	fmt.Println(ageOne, ageTwo, ageThree)

	// bits & memory
	var numOne int8 = 127
	var numTwo int8 = -128
	var numThree uint8 = 255
	fmt.Println(numOne, numTwo, numThree)

	var scoreOne float32 = 123847128.98
	var scoreTwo float64 = 2523749238749283479283472398749238749238749283749.98
	scoreThree := 1.5 //float64
	fmt.Println(scoreOne, scoreTwo, scoreThree)

	// Print:
	fmt.Print("Hi, ")
	fmt.Print("Rafa \n")
	fmt.Print("new line \n")
	fmt.Println("cya")

	// Formatted strings
	fmt.Printf("my age is %v and my name is %v\n", ageOne, nameFour)
	fmt.Printf("my age is %q and my name is %q\n", ageOne, nameFour) // with quotes (for strings)
	fmt.Printf("age is of type %T\n", ageOne)                        // type
	fmt.Printf("you scored %f points\n", scoreThree)
	fmt.Printf("you scored %0.2f points\n", scoreThree)

	// Sprintf (save formatted strings)
	str := fmt.Sprintf("my age is %v and my name is %v\n", ageOne, nameFour)
	fmt.Println("Saved string: ", str)
}
