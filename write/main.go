package main

import (
	"fmt"
	"os"
)

func main() {
	// write an entire text file, overwrite if file already there
	txt := []byte("Here's a multiple line text file\nso that I can practice writing\nfiles with the GO book exercises.\nAnother line for kicks.")
	err := os.WriteFile("/Users/mike/Documents/apps/go-easy-book/sampleWrite.txt", txt, 0666)
	check(err)

	// open the same file for an append operation
	file, err := os.OpenFile("/Users/mike/Documents/apps/go-easy-book/sampleWrite.txt", 0666, os.FileMode(os.O_APPEND))
	check(err)
	// close file at end of function after other operations finish
	defer file.Close()

	// append text to the file content and check for errors
	addText := []byte("\nby Extreme Lord Mastermind Dave")
	// get total number of bytes to know where to start write
	totalBytes, err := os.ReadFile("/Users/mike/Documents/apps/go-easy-book/sampleWrite.txt")
	check(err)
	lenTotalBytes := len(totalBytes)
	bytesWritten, err := file.WriteAt(addText, int64(lenTotalBytes))
	check(err)

	// display number of chars written and what has been written
	fmt.Printf("%v chars appended: %v", bytesWritten, string(addText[:bytesWritten]))
}

func check(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
