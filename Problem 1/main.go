package main

//import fmt, whic provides functions for formatting and printing text
import (
	"fmt"
)

// this block of code defines 'main' function. insine main function
// we declare variabels 'e1,'e2','e3' which contain the valies
// "Hai", "My Name", and "Is caca" respectively
func main() {
	e1 := "Hai "
	e2 := "My Name"
	e3 := " Is caca"
	hasil := concatString(e1, e2, e3) //we then call the "concatString" function with "e1, e2, e3" as arguments, and store
	//the returned concatened string in variable named "hasil"
	fmt.Printf(hasil)
}

// take variabel number of string arguments ang returns a concatened string
func concatString(strs ...string) string {
	var hasil string //declare a variable named "hasil" and initialize it to an empty string
	//use for loop with the range keyword to iterate over each argument in the "strs" slice
	// for each argument, we concatenate it to the "hasil1" variable using the += operator
	for _, str := range strs {
		hasil += str
	}
	return hasil
}
