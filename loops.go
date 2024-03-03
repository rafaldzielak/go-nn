package main

import "fmt"

func loops() {
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
}
