package main

import "fmt"

type user struct {
	name string
	age  int
	isLoggedIn bool
	greet func()
}

func main() {
	user1 := user{
		name: "Emon",
		age: 21,
		isLoggedIn: true,
		greet: func() {
			fmt.Printf("Hello, %s! You are %d years old.\n", user1.name, user1.age)
		},
	} 

	user2 := user{
		name: "Sza",
		age: 2,
		isLoggedIn: false,
		greet: func() {
			fmt.Printf("Hello, %s! You are %d years old.\n", user2.name, user2.age)
		},
	}		

	user1.greet()
	user2.greet()
}
