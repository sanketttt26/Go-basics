package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Learn Slices")
	var fruitList = []string{"Apple", "Peach", "Tomato"}
	fruitList = append(fruitList, "Mango", "Banana")
	fmt.Println(fruitList)

	fmt.Println(fruitList[:2])

	highScores := make([]int, 4)
	highScores[0] = 3232
	highScores[1] = 33234
	highScores[2] = 786
	highScores[3] = 987654

	fmt.Println(highScores)
	highScores = append(highScores, 111, 444, 12312, 342)
	fmt.Println(highScores)

	sort.Ints(highScores)
	fmt.Println("Sorted order", highScores)

	var courses = []string{"react", "javascript", "c#", ".net", "python"}
	var index int = 2

	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println("Available courses", courses)
}
