package main

import "fmt"

func arraySlices() {
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
