package main

import (
	"fmt"
	"math"
	"strings"
)

// Gophercises
// go run main.go greetings.go conditionals.go arraysSlices.go loops.go standardLibrary.go types.go maps.go passByValue.go pointers.go

func cycleNames(n []string, f func(string)) {
	for _, value := range n {
		f(value)
	}
}

func circleArea(r float64) float64 {
	return math.Pi * r * r
}

var score = 99.5

var strs = []string{"rafa", "dyrek", "mario", "luigi", "peach"}

func getInitials(n string) (string, string) {
	s := strings.ToUpper(n)
	names := strings.Split(s, " ")
	var initials []string
	for _, value := range names {
		initials = append(initials, value[:1])
	}
	// fmt.Println(initials, len(initials))
	if (len(initials)) > 1 {
		return initials[0], initials[1]
	}
	return initials[0], "_"
}

// Does not work outside functions:
// someName := "hello"

// "main" is important, go will look for main at the beginning, only one main function in allowed in the app
func main() {

	types()
	arraySlices()
	standardLibrary()
	loops()
	conditionals()

	// PACKAGE SCOPE:
	sayGreeting("Rafa")
	sayBye("Rafa")

	cycleNames(strs, sayGreeting)
	cycleNames(strs, sayBye)

	a1 := circleArea(4)
	a2 := circleArea(5)
	fmt.Printf("circle 1 is %0.3f and circle 2 is %0.3f\n", a1, a2)

	initialOne, initialTwo := getInitials("Rafa Dyrek")
	fmt.Println(initialOne, initialTwo)

	for _, v := range points {
		fmt.Println(v)
	}

	showScore()

	maps()
	passByValue()
	pointers()

	structs()

}
