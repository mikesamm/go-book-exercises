package main

import (
	"fmt"
	"os"
)

func main() {
	// create a temp file, check for errors, confirm it exists
	tempFile, err := os.CreateTemp("", "Data-*")
	check(err)
	fmt.Printf("\nCreated File: %v\n", tempFile.Name())

	// write to the temporary file
	bytesWritten, err := tempFile.WriteString("Hey, this is a string in a temp file.\n")
	check(err)

	// read the temporary file
	txt, err := os.ReadFile(tempFile.Name())
	check(err)
	fmt.Printf("\nTemp file contains:\n%v bytes -> %v", bytesWritten, string(txt))

	// close the file after processes complete
	tempFile.Close()
	// remove the temporary file
	os.Remove(tempFile.Name())

	_, err = os.Stat(tempFile.Name())
	check(err)
}

func check(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
