package main

import "fmt"

// global variables
// jwtToken := 3000.0         ----- Not allowed
// jwtToken = 3000.0          ----- Allowed
// var jwtToken int = 30000   ----- Allowed

// Constant Variables
const LoginToken string = "gibberish"

// captial L is having significant importance here
// variable name starting with first letter capital then it is said to be a public variable
// incasae if it's small case its private const variable

func main() {
	fmt.Println("Variables")
	var username string = "Ayush"
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T \n", username)

	var isLoggedIn bool = false
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type: %T \n", isLoggedIn)

	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("Variable is of type: %T \n", smallVal)

	// uint8       the set of all unsigned  8-bit integers (0 to 255)
	// uint16      the set of all unsigned 16-bit integers (0 to 65535)
	// uint32      the set of all unsigned 32-bit integers (0 to 4294967295)
	// uint64      the set of all unsigned 64-bit integers (0 to 18446744073709551615)

	// int8        the set of all signed  8-bit integers (-128 to 127)
	// int16       the set of all signed 16-bit integers (-32768 to 32767)
	// int32       the set of all signed 32-bit integers (-2147483648 to 2147483647)
	// int64       the set of all signed 64-bit integers (-9223372036854775808 to 9223372036854775807)

	// float32     the set of all IEEE-754 32-bit floating-point numbers
	// float64     the set of all IEEE-754 64-bit floating-point numbers

	// complex64   the set of all complex numbers with float32 real and imaginary parts
	// complex128  the set of all complex numbers with float64 real and imaginary parts

	// byte        alias for uint8
	// rune        alias for int32

	var smallFloat float64 = 255.237492349234
	fmt.Println(smallFloat)
	fmt.Printf("Variable is of type: %T \n", smallFloat)

	// default values and some aliases

	var anotherVariable int
	fmt.Println(anotherVariable)
	fmt.Printf("Variable is of type: %T \n", anotherVariable)

	var anotherVariable2 string
	fmt.Println(anotherVariable2)
	fmt.Printf("Variable is of type: %T \n", anotherVariable2)

	// implicit way of declaring the variable
	// not declaring type, lexer will do the work for me 😎
	var website = "ayushsoni1010.vercel.app"
	fmt.Println(website)

	// No var style
	numberOfUser := 30000.0
	fmt.Println(numberOfUser)
	fmt.Printf("Variable is of type: %T \n", numberOfUser)

	fmt.Println(LoginToken)
	fmt.Printf("Variable is of type: %T \n", LoginToken)

}
