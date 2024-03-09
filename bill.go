package main

import "fmt"

type bill struct {
	name  string
	items map[string]float64
	tip   float64
}

// make a new bill

func newBill(name string) bill {
	b := bill{
		name:  name,
		items: map[string]float64{},
		tip:   0,
	}
	return b
}

func structs() {
	b := newBill("rafa")
	b.items["pie"] = 5.99
	b.items["cake"] = 3.99
	b.tip = 2.5
	fmt.Println(b)
}
