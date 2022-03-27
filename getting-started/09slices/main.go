package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to slices in golang")
	var fruitList = []string{"Apple", "Peach", "Tomato"}
	fmt.Printf("Type of fruit list is: %T\n", fruitList)

	fruitList = append(fruitList, "Mango", "Banana")
	fmt.Println(fruitList)

	fruitList = append(fruitList[1:])
	fruitList = append(fruitList[1:3])
	fruitList = append(fruitList[:3])
	fmt.Println(fruitList)

	highScores := make([]int, 4)

	highScores[0] = 234
	highScores[1] = 34
	highScores[2] = 4
	highScores[3] = 22
	// highScores[4] = 777

	highScores = append(highScores, 555, 666, 777)
	fmt.Println(len(highScores))

	fmt.Println(highScores)
	fmt.Println(sort.IntsAreSorted(highScores))

	sort.Ints(highScores)
	fmt.Println(highScores)
	fmt.Println(sort.IntsAreSorted(highScores))

	// how to remove a value from slices based on index

	var courses = []string{"reactjs", "swift", "javascript", "python", "ruby"}
	fmt.Println(courses)
	var index int = 1

	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)
}
