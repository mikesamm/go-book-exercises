package main

import (
	"flag"
	"fmt"
)

func main() {
	// define three command-line flags
	// the flag function arguments are: <flag name>, <default value>, <help message>
	// the functions return a memory address for an appropriate variable data type
	txt := flag.String("txt", "Hello", "A string")
	num := flag.Int("num", 23, "A number")
	status := flag.Bool("status", true, "A boolean")

	// process flags and display their values *SEE BELOW*
	flag.Parse()
	// dereference the value at the address of `txt`, `num`, and `status`
	fmt.Println("\nText:", *txt)
	fmt.Println("Number:", *num, "Status:", *status)
}

/*
when running this program, you can change the values of the flags with the following syntax:
...
> ./flags

Text: Hello
Number: 23 Status: true
> ./flags -txt="buon giorno" -num=42 -status=false

Text: buon giorno
Number: 42 Status: false
...

for single word string, quotes can be omitted
*/
