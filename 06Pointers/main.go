package main

import "fmt"

func main() {
	fmt.Println("Learning Pointers")
	var a int = 10
	var p = &a // assigning the address of 'a' to the pointer 'p'

	fmt.Println("Value of a:", a)
	fmt.Println("Address of a:", &a)
	fmt.Println("Value of p (address of a):", p)
	fmt.Println("Value at address stored in p:", *p)

	*p = 20 // changing the value at the address stored in 'p'
	fmt.Println("New value of a:", a)

}
