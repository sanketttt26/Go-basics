package main

import "fmt"

func main() {
	fmt.Println("Welcome from main")
	greeter()
	greeterTwo()
	discountAmt := calculateDiscount(12000, 10)
	fmt.Println("You have saved $", discountAmt)
}

func greeter() {
	fmt.Println("Welcome from greeter 1")
}

func greeterTwo() {
	fmt.Println("Welcome from greeter 2")
}

func calculateDiscount(price float64, discountRate float64) float64 {
	discount := price * (discountRate / 100)
	finalPrice := price - discount
	return finalPrice
}
