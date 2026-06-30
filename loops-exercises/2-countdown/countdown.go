package main

import "fmt"

func main() {

	var number int

	fmt.Print("Type a number: ")
	fmt.Scan(&number)

	for number := number; number > 0; number-- {

		fmt.Println(number)

	}

}
