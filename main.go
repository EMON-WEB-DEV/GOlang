package main

import "fmt"



func main() {
	a := 10
	b := 20

	a = 50
	p := &a
	sum := a + b

	fmt.Printf("The sum of %d and %d is %d\n", a, b, sum)
	fmt.Printf("The address of a is %p\n", p)
}
