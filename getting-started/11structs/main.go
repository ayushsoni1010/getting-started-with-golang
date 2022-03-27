package main

import "fmt"

func main() {
	fmt.Println("Structs in golang")
	// no inheritance in golang; No super or parent or child

	ayush := User{"Ayush", "ayushsoni1010.work@gmail.com", true, 21}

	fmt.Println(ayush)
	fmt.Printf("Ayush details are: %+v\n", ayush)
	fmt.Printf("Name is %v and Email is %v", ayush.Name, ayush.Email)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
