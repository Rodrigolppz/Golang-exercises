package main

import "fmt"

func main() {

	var number int
	var sum int

	fmt.Print("Give me a number: ")
	fmt.Scan(&number)

	for index := 0; index <= number; index++ {
		sum = sum + index
		fmt.Println(sum)

	}

}
