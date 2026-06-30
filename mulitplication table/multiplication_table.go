package main

import (
	"fmt"
)

func main() {

	var number int
	var result int

	fmt.Print("Write a number: ")
	fmt.Scan(&number)

	for x := 0; x < 9; x = x + 1 {

		result = number * x
		fmt.Printf("%v x %v = %v \n", number, x, result)

	}

}
