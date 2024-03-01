package main

import (
	"fmt"
	"math"
	"sort"
	"strings"
)

func sayGreeting(n string) {
	fmt.Printf("Good morning %v\n", n)
}

func sayBye(n string) {
	fmt.Printf("Bye %v\n", n)
}

func cycleNames(n []string, f func(string)) {
	for _, value := range n {
		f(value)
	}
}

func circleArea(r float64) float64 {
	return math.Pi * r * r
}

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

	// STANDARD LIBRARY:
	// strings:
	greeting := "Greetings from the mountain"
	fmt.Println(strings.Contains(greeting, "Greetings"), strings.ReplaceAll(greeting, "Greetings", "Hi"))
	fmt.Println("Original value:", greeting, strings.ToUpper(greeting))
	fmt.Println(strings.Index(greeting, "ee"), strings.Split(greeting, " "))
	// sort:
	nums := []int{45, 20, 35, 30, 75, 60, 50, 25}
	sort.Ints(nums) //alters the slice
	fmt.Println((nums))
	index := sort.SearchInts(nums, 999)
	fmt.Println(index) //if it does not find it - nums.length+1
	strs := []string{"rafa", "dyrek", "mario", "luigi", "peach"}
	sort.Strings(strs)
	fmt.Println(strs)
	fmt.Println(sort.SearchStrings(strs, "dyrek"))

	// LOOPS:
	x := 0
	for x < 5 {
		// fmt.Println("Value of x:", x)
		x++
	}
	for i := 0; i < 5; i++ {
		// fmt.Println("Value of i:", i)
	}
	for i := 0; i < len(strs); i++ {
		// fmt.Println(strs[i])
	}
	// for index, value := range strs {
	// 	fmt.Printf("the value at index %v is %v\n", index, value)
	// }
	for _, value := range strs {
		fmt.Printf("the value is %v\n", value)
		value = "new string" // does NOT update the value
	}
	fmt.Println(strs)

	// BOOLEANS & CONDITIONALS
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

	sayGreeting("Rafa")
	sayBye("Rafa")

	cycleNames(strs, sayGreeting)
	cycleNames(strs, sayBye)

	a1 := circleArea(4)
	a2 := circleArea(5)
	fmt.Printf("circle 1 is %0.3f and circle 2 is %0.3f\n", a1, a2)

	initialOne, initialTwo := getInitials("Rafa Dyrek")
	fmt.Println(initialOne, initialTwo)

}
