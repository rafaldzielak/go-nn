package main

import (
	"fmt"
	"sort"
	"strings"
)

func standardLibrary() {
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
	sort.Strings(strs)
	fmt.Println(strs)
	fmt.Println(sort.SearchStrings(strs, "dyrek"))
}
