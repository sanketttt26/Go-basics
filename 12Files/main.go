package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("Welcome to learn about files in go")
	content := "This text will appear in the file we created just now"
	file, err := os.Create("./myGoFile.txt")

	checkNilError(err)

	length, err := io.WriteString(file, content)
	checkNilError(err)

	fmt.Println("The length of the file is", length)
	readFile("./myGoFile.txt")
	defer file.Close()
}

func readFile(filePath string) {
	data, err := os.ReadFile(filePath)

	checkNilError(err)

	fmt.Println("The data inside the file is \n", string(data))
}

func checkNilError(err error) {
	if err != nil {
		panic(err)
	}
}
