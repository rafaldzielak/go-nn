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

	// ARRAYS & SLICES:
	// Array has a fixed length
	var ages [3]int = [3]int{1, 2, 3}
	var agesTwo = [3]int{1, 2, 3}
	names := [4]string{"Rafa", "Dyrek", "tt", "jj"}
	names[1] = "Pawa"
	fmt.Println(ages, agesTwo, names, len(names))

	// Slices (use arrays under the hood):
	var scores = []int{1, 2, 3}
	scores[2] = 25
	scores = append(scores, 4)
	fmt.Println(scores, len(scores))
	// slice rankges:
	range1 := names[1:3] //contains 1 and 2
	range2 := names[2:]  //to the end
	range3 := names[:3]  //0,1,2
	fmt.Println(range1, range2, range3)
	range1 = append(range1, "aa")
	fmt.Println(range1)

}
