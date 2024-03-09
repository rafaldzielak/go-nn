package main

import "fmt"

type bill struct {
	name  string
	items map[string]float64
	tip   float64
}

func newBill(name string) bill {
	b := bill{
		name:  name,
		items: map[string]float64{},
		tip:   0,
	}
	return b
}

// Receiver function - receives a COPY of the bill
func (b bill) format() string {
	fs := "Bill breakdown:\n"
	var total float64 = 0
	for k, v := range b.items {
		fs += fmt.Sprintf("%-25v ...$%v \n", k+":", v) //first one takes 25 characters
		total += v
	}

	fs += fmt.Sprintf("%-25v ...$%0.2f", "total:", total)
	return fs
}

func structs() {
	fmt.Println("STRUCTS:")
	myBill := newBill("rafa")
	myBill.items["pie"] = 5.99
	myBill.items["cake"] = 3.99
	myBill.tip = 2.5
	fmt.Println(myBill)
	fmt.Println(myBill.format())
}
