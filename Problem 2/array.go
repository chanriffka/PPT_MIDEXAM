package main

import (
	"fmt"
)

func array(a1 []string, a2 []string) { // function named array, which takes two string s;ice "a1" and "a2" as parameters
	similiar := true // this variable will be used to determine if the two arrays are same or not
	//this block of code loops through each index "a1 and a2" using for loop, comparing the value at each index
	//if the values are not the same, the "similiar" variable is set to "false" and message is printed to the console indicating
	//that two arrays are differents at that index
	for i := 0; i < len(a1); i++ {
		if a1[i] != a2[i] {
			similiar = false
			fmt.Printf("Index ke %d berbeda \n", i)
		}
	}
	//the similiar variable is still "true", it means that all emenets in both arrays are the same and a message
	//is printed to the console indicatinng that both arrays are the same
	if similiar {
		fmt.Printf("Index ke 1 dan Index ke 2 Sama")
	}
}

// inside the "main" function, a message is printed to the console asking the user to input the leght of arrays
func main() {
	fmt.Println("Input lenght array")
	var b8 int //the variable b8 is then declared and its value is read from the user input using the "fmt.Scanln" function
	fmt.Scanln(&b8)
	//this block of code creates a string slice named array1 with a length of b8. A message is printed to the console asking the user to input values for array1
	array1 := make([]string, b8)
	fmt.Println("array 1: ")
	for i := 0; i < b8; i++ { //a for loop is used to read values from the user input and store them in the array1 slice
		fmt.Scan(&array1[i])
	}
	//this block of code creates another string slice named array2 with a length of b8. a message is printed to the console asking the user to input values for array2
	array2 := make([]string, b8)
	fmt.Println("array 2: ")
	for i := 0; i < b8; i++ { //a "for" loop is used to read values from the user input and store them in the "array2" slice
		fmt.Scan(&array2[i])
	}
	array(array1, array2)
}
