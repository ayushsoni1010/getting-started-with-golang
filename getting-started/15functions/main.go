package main

import "fmt"

func main() {
	fmt.Println("Welcome to functions in golang")
	greeter()
	greeterTwo()

	result := adder(3, 5)
	fmt.Println("Result is: ", result)

	proResult, myMessage := proAdder(2, 5, 7, 8, 9, 10)
	fmt.Println("Pro result is: ", proResult)
	fmt.Println("Pro message is: ", myMessage)
}

func proAdder(values ...int) (int, string) {
	total := 0

	for _, val := range values {
		total += val
	}
	return total, "Hi Pro result function"
}

func adder(valOne int, valtwo int) int {
	return valOne + valtwo
}

func greeterTwo() {
	fmt.Println("Another method")
}

func greeter() {
	fmt.Println("Namaste from golang")
}
