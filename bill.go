package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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

	switch option {
	case "a":
		fmt.Println("you chose a")
		name, _ := getInput("Item name: ", reader)
		price, _ := getInput("Item price: ", reader)
		p, err := strconv.ParseFloat(price, 64)
		if err != nil {
			fmt.Println("The price must be a number")
			promptOptions(b)
		}
		b.addItem(name, p)
		fmt.Println("Item added - ", name, price)
		promptOptions(b)

	case "t":
		fmt.Println("you chose t")
		tip, _ := getInput("Tip amount: ", reader)
		t, err := strconv.ParseFloat(tip, 64)
		if err != nil {
			fmt.Println("The tip must be a number")
			promptOptions(b)
		}
		b.updateTip(t)
		fmt.Println("Tip added - ", t)
		promptOptions(b)

	case "s":
		fmt.Println("you chose to save the bill", b.format())
		b.save()
	default:
		fmt.Println("that was not a valid option...")
		promptOptions(b)

	}

}

func createBill() bill {
	reader := bufio.NewReader(os.Stdin)
	name, _ := getInput("Create a new bill name: ", reader)

	name = strings.TrimSpace(name)
	b := newBill(name)
	fmt.Println("Created the bill - ", b.name)
	return b
}

func (b *bill) save() {
	data := []byte(b.format())

	err := os.WriteFile("./bills/"+b.name+".txt", data, 0644)
	if err != nil {
		panic(err)
	}
	fmt.Println("Bill was saved to file - ", b.name)
}
