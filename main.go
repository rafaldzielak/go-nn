package main

import "fmt"

// Does not work outside functions:
// someName := "hello"

// "main" is important, go will look for main at the beginning, only one main function in allowed in the app
func main() {
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

}
