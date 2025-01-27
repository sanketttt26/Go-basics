package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Write something valuable..."
	fmt.Println(welcome)
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Add your ratings")
	input, _ := reader.ReadString('\n')
	fmt.Println("Your ratings are ", input)

}
