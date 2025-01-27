package main

import "fmt"

func main() {
	fmt.Println("Variables")
	var username string = "sankettt"
	fmt.Println(username)
	fmt.Printf("The variable is of type : %T \n", username)

	var username2 string
	fmt.Println(username2)
	fmt.Printf("The variable is of type : %T \n", username2)
}
