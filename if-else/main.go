package main

import "fmt"


func main() {

	age := 21

	if age < 18 {
		fmt.Println("you are a minor")
	} else {
		fmt.Println("you are an adult")
	}
}