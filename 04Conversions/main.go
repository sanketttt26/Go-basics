package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome user !!!")
	fmt.Println("Enter your rating for the app")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for rating us with ", input)

	numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)
	if err != nil {
		fmt.Println(err)
	} else {
		numRating = numRating + 1
		fmt.Println(numRating)
	}

	s := strconv.FormatFloat(3.1423, 'f', 2, 64)
	fmt.Println(s)

	s1 := strconv.Itoa(42)
	fmt.Println(s1)

}
