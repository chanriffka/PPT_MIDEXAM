package main

import (
	"fmt"
)

func array(a1 []string, a2 []string) {
	similiar := true
	for i := 0; i < len(a1); i++ {
		if a1[i] != a2[i] {
			similiar = false
			fmt.Printf("Index ke %d berbeda \n", i)
		}
	}
	if similiar {
		fmt.Printf("both index are same")
	}
}

func main() {
	fmt.Println("Input lenght array")
	var b8 int
	fmt.Scanln(&b8)

	array1 := make([]string, b8)
	fmt.Println("array 1: ")
	for i := 0; i < b8; i++ {
		fmt.Scan(&array1[i])
	}
	array2 := make([]string, b8)
	fmt.Println("array 2: ")
	for i := 0; i < b8; i++ {
		fmt.Scan(&array2[i])
	}
	array(array1, array2)
}
