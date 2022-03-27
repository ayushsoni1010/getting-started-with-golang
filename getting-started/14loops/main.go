package main

import "fmt"

func main() {
	fmt.Println("Welcome to loops in golang")
	days := []string{"Sunday", "Tuesday", "Wednesday", "Friday", "Saturday "}
	fmt.Println(days)

	// for day := 0; day < len(days); day++{
	// 	fmt.Println(days[day])
	// }

	// another way of printing the slice (array)
	// for i := range days {
	// 	fmt.Println(days[i])
	// }

	// For each kind of loop
	for index, day := range days {
		fmt.Printf("Index is %v and value is %v\n", index, day)
	}

	rougueValue := 1

	for rougueValue < 10 {

		if rougueValue == 8 {
			goto portfolio
		}

		if rougueValue == 5 {
			rougueValue++
			continue
		}
		// if rougueValue == 5 {
		// 	break
		// }
		fmt.Println("Value is : ", rougueValue)
		rougueValue++
	}

portfolio:
	fmt.Println("Jumping to ayushsoni1010.vercel.app")
}
