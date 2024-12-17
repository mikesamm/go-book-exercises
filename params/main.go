package main

import (
	"fmt"
	"os"
)

func main() {
	// display default location of program along with other params
	for i, param := range os.Args {
		fmt.Printf("Argument %v: %v\n", i, param)
	}

	// display a specific argument from Args slice
	fmt.Printf("\nFirst argument: %v", os.Args[1])
	fmt.Printf("\nLast argument: %v", os.Args[len(os.Args)-1])
}
