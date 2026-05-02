package main

import "fmt"

// printInit is a regular function that can be called from main.
func printInit() {
	fmt.Println("This is the init function. It runs before the main function.")
}

func main() {
	printInit()
}
