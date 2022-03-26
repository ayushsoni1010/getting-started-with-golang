package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome again to our Pizza app")
	fmt.Println("Please rate our Pizza between 1 and 5")

	reader := bufio.NewReader(os.Stdin)

	input, _ := reader.ReadString('\n')

	fmt.Println("Thanks for rating: ", input)

	// type conversion of string to number by using strconv package
	// numRating = input + 1       ----- Will give error becoz input is type of string and adding number to it

	// numRating, err := strconv.ParseFloat(input, 64)    --- Gives Error becoz in input it contains \n so could able to add 1 to it in numRating
	numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Added 1 to your rating: ", numRating+1)
	}
}
