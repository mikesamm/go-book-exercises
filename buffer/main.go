package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// prompt the user for text
	fmt.Print("\nEnter text: ")

	// create a Scanner instance to store user input in a memory buffer
	scanner := bufio.NewScanner(os.Stdin)
	// read the user input
	scanner.Scan()

	// output an error if one occurred
	if scanner.Err() != nil {
		fmt.Println(scanner.Err())
	} else { // output stored data
		fmt.Println(scanner.Text())
	}

	// prompt the user for more text
	fmt.Print("\nEnter text: ")
	// create a new Scanner instance to store another user input
	scanner2 := bufio.NewScanner(os.Stdin)
	// read the new user input
	scanner2.Scan()

	// store individual items of the user input in individual elements
	// of a slice
	words := strings.Fields(scanner2.Text())

	// output an error if one occurred
	if scanner2.Err() != nil {
		fmt.Println(scanner2.Err())
	} else {
		for i, word := range words {
			fmt.Printf("%v: %v\n", i, word)
		}
	}

}
