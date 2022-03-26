package main

import "fmt"

func main() {
	fmt.Println("Welcome to pointers in golang")

	// var ptr *int
	// fmt.Println("Value of pointer is: ", ptr)
	// default value of pointer is nil

	myNumber := 23

	var myPtr *int = &myNumber

	fmt.Println("Value of actual pointer is: ", myPtr)
	fmt.Println("Value of actual pointer is: ", *myPtr)
	fmt.Println("Value of actual pointer is: ", &myNumber)
	fmt.Println("Value of actual pointer is: ", *&myNumber)

	*myPtr = *myPtr + 2
	fmt.Println("New value is: ", myNumber)
}
