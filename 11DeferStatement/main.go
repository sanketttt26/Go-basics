package main

import "fmt"

func main() {
	fmt.Println("Hello")
	defer fmt.Println("Kya bhai")
	defer fmt.Println("World")
	myDefer()
}

func myDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Print(i)
	}
}
