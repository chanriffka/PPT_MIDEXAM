package main

import (
	"fmt"
)

func main() {
	e1 := "Hai "
	e2 := "My Name"
	e3 := " Is caca"
	hasil := concatString(e1, e2, e3)
	fmt.Printf(hasil)
}

func concatString(strs ...string) string {
	var hasil string
	for _, str := range strs {
		hasil += str
	}
	return hasil
}
