package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

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
func (b *bill) format() string {
	fs := "Bill breakdown:\n"
	var total float64 = 0
	for k, v := range b.items {
		fs += fmt.Sprintf("%-25v ...$%v \n", k+":", v) //first one takes 25 characters
		total += v
	}
	fs += fmt.Sprintf("%-25v ...$%v\n", "tip:", b.tip)

	fs += fmt.Sprintf("%-25v ...$%0.2f", "total:", total+b.tip)
	return fs
}

func (b *bill) updateTip(tip float64) {
	b.tip = tip //we don't need to dereference here (*b)
}

func (b *bill) addItem(name string, price float64) {
	b.items[name] = price

}

func structs() {
	fmt.Println("STRUCTS:")
	myBill := newBill("rafa")
	myBill.addItem("pie", 5.99)
	myBill.addItem("cake", 3.99)
	myBill.updateTip(2.5)
	fmt.Println(myBill)
	fmt.Println(myBill.format())
}

func promptOptions(b bill) {
	reader := bufio.NewReader(os.Stdin)
	option, _ := getInput("Choose option (a - add item, s - save the bill, t = add tip) ", reader)
	fmt.Println(option)

}

func createBill() bill {
	reader := bufio.NewReader(os.Stdin)
	name, _ := getInput("Create a new bill name: ", reader)

	name = strings.TrimSpace(name)
	b := newBill(name)
	fmt.Println("Created the bill - ", b.name)
	return b
}
