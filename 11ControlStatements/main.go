package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Control statements")

	expectedUsername := "sanket123"
	expectedPassword := "password123"

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the username")
	inputUsername, _ := reader.ReadString('\n')
	fmt.Println("Enter the password")
	inputPassword, _ := reader.ReadString('\n')

	inputUsername = strings.TrimSpace(inputUsername)
	inputPassword = strings.TrimSpace(inputPassword)

	if inputUsername == expectedUsername && inputPassword == expectedPassword {
		fmt.Println("Login successful !!!")
	} else {
		fmt.Println("Invalid username or password. try again")
	}

	for i := 0; i < 10; i++ {
		if i == 5 {
			goto end
		}
		fmt.Println(i)
	}

end: // label
	fmt.Println("Exited the loop using goto ")
}
