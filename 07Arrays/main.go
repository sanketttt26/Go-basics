package main

import (
	"fmt"
)

func main() {
	fmt.Println("Arrays in golang ")

	var fruitList [4]string
	fruitList[0] = "Apple"
	fruitList[1] = "Banana"
	fruitList[2] = "Peach"

	fmt.Println("Values in my fruitlist are", fruitList)
	fmt.Println("Values in my fruitlist are", len(fruitList))

	var veggiesList = []string{"Potato", "GreenPeas", "Beans"}

	fmt.Println("Veggies list containes ", veggiesList)

	var itemList [3]string = [3]string{"Bucket", "Towel", "Fan"}
	fmt.Println(itemList)

}
