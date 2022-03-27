package main

import "fmt"

func main() {
	fmt.Println("Welcome to Methods in golang")

	ayush := User{"Ayush", "ayushsoni1010.work@gmail.com", true, 21}
	fmt.Println(ayush)
	fmt.Printf("Ayush details are: %+v\n", ayush)
	fmt.Printf("Name is %v and email is %v\n\n", ayush.Name, ayush.Email)
	ayush.GetStatus()
	ayush.NewEmail()
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus() {
	fmt.Println("Is user active: ", u.Status)
}

func (u User) NewEmail() {
	u.Email = "test@go.dev"
	fmt.Println("Email is this user is: ", u.Email)
}
