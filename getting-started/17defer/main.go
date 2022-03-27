package main

import "fmt"

func main() {
	defer fmt.Println("Hey I am executing")
	defer fmt.Println("One")
	defer fmt.Println("Two")
	fmt.Println("Welcome to Defer in golang")
	myDefer()
}

func myDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}
