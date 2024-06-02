package main

import "fmt"

type ItemSale struct {
	name    string
	cost    float32
	taxRate float32
}

func main() {
	itemsSold := []ItemSale{
		{name: "Cake", cost: 0.99, taxRate: 7.5},
		{name: "Milk", cost: 2.75, taxRate: 1.5},
		{name: "Butter", cost: 0.87, taxRate: 2},
	}
	salesTaxTotal := float32(0)

	for _, value := range itemsSold {
		salesTaxTotal += (value.cost * value.taxRate * 0.01)
	}
	fmt.Printf("Total sales tax %#v", salesTaxTotal)
}
