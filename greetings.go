package main

import "fmt"

// functions and variables are automatically accessible in any other file in package main
// to run both programs: go run main.go greetings.go

var points = []int{20, 30, 50, 100, 200, 500, 900, 1000}

func sayGreeting(n string) {
	fmt.Printf("Good morning %v\n", n)
}

func sayBye(n string) {
	fmt.Printf("Bye %v\n", n)
}

func showScore() {
	fmt.Println("Score:", score)
}
