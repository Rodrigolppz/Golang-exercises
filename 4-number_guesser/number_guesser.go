package main

import "fmt"

func main() {

	number := 9
	var guess int

	fmt.Print("Guess an integer number: ")
	fmt.Scan(&guess)

	if guess == number {

		fmt.Printf("Good job, the secret number is: %v \n", number)

	} else {
		fmt.Printf("You're wrong!!!, the secret number is not %v \n", guess)
	}

}
