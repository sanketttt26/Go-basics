package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome, let's study time package")

	presentTime := time.Now()

	fmt.Println(presentTime.Format("2006-01-02 15:04:05 Monday"))

	createdDate := time.Date(2021, time.November, 22, 22, 22, 0, 0, time.UTC)
	fmt.Println(createdDate)

}
