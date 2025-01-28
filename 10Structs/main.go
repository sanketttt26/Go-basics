package main

import "fmt"

func main() {
	fmt.Println("Structs in Go")

	user1 := User{"Sanket", 20, true, "sanket123@gmail.com"}

	fmt.Printf("The name of the user is %v , the age is %v , the status is %v the email of the user is %v", user1.Name, user1.Age, user1.IsActive, user1.Email)
}

type User struct {
	Name     string
	Age      int
	IsActive bool
	Email    string
}
