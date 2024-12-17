package main

import (
	"fmt"
	"os"
)

func main() {
	// read the file at the given path
	txt, err := os.ReadFile("/Users/mike/Documents/apps/go-easy-book/sample.txt")
	// check for errors
	check(err)
	// print the text. string() is needed to translate the byte slice char code numbers
	// into their respective alphabet chars
	fmt.Println(string(txt))

	// open the file at the given path
	file, err := os.Open("/Users/mike/Documents/apps/go-easy-book/sample.txt")
	check(err)
	// close the file after other operations complete
	defer file.Close()

	// specify starting position and check for errors
	// Seek(offset, from where / "whence")
	start, err := file.Seek(42, 0)
	check(err)

	// read 15 characters and check for errors
	// make a slice of 15 byte elements
	phraseSlice := make([]byte, 15)
	// Read returns number of bytes read
	bytesRead, err := file.Read(phraseSlice)
	check(err)

	fmt.Printf("\n%v bytes @ %v: ", bytesRead, start)
	fmt.Printf("%v\n", string(phraseSlice[0:bytesRead]))
	// :bytesRead is 15 which would be the end of the slice
	// since there's no number before the colon, it prints from
	// 0 to the bytesRead index
	// since it's the whole slice, bytesRead can be omitted
}

func check(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
