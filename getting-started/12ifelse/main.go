package main

import "fmt"

func main() {
	fmt.Println("if-else in Golang")
	loginCount := 10

	var result string

	if loginCount < 10 {
		result = "Regular User"
	} else if loginCount > 10 {
		result = "Watch out"
	} else {
		result = "Exactly 10 login count"
	}
	fmt.Println(result)

	if 9 % 2 == 0 {
		fmt.Println("Number is even")
	} else {
		fmt.Println("Number is odd")
	}

	// declaring it by another way when handling with web request, assigning num variable instantly and checking num
	if num :=3; num < 10{
		fmt.Println("Number is less than 10")
	} else{
		fmt.Println("Number is not less than 10")
	}
}
