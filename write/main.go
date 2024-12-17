package main

import (
	"fmt"
	"os"
)

func main() {
	txt := []byte("Here's a multiple line text file\nso that I can practice writing\nfiles with the GO book exercises.\nAnother line for kicks.")
	err := os.WriteFile("/Users/mike/Documents/apps/go-easy-book/sampleWrite.txt", txt, 0644)
	check(err)
}

func check(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
