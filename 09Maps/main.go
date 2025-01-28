package main

import "fmt"

func main() {
	fmt.Println("MAPS IN GOLANG")

	languages := make(map[string]string)
	languages["JS"] = "Javascript"
	languages["RB"] = "Ruby"
	languages["PY"] = "Python"
	fmt.Println("List of all languages", languages)
	delete(languages, "RB")
	fmt.Println("List of all languages", languages)

	for key, value := range languages {
		fmt.Printf("The key is %v and the value is %v \n", key, value)
	}
}
