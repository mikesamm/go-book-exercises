package main

import (
	"fmt"
	"math/rand"
)

// this is a cli number guessing game

func main() {
	// generate a random number up to 20
	var num int = rand.Intn(20) + 1
	// initialize guess
	var guess int = 0
	// initialize value for loop execution/iteration
	var flag bool = true

	// prompt the user for a number guess
	fmt.Printf("\nGuess my number, 1-20: ")

	// while the number has not been guessed
	for flag {
		// store error if one occurs during scan
		_, err := fmt.Scan(&guess)

		// if there's an error, print it
		// otherwise, prompt user accordingly
		if err != nil {
			fmt.Println(err)
		} else if guess > num {
			fmt.Print("Too high, try again: ")
		} else if guess < num {
			fmt.Print("Too low, try again: ")
		} else if guess == num {
			fmt.Println("Correct! The number is", num)
			flag = false
		}
	}

}
