package main

import "fmt"

func main() {

	var age = 22

	if age >= 18 {
		fmt.Println("You can vote!")
	} else {
		fmt.Println("You cannot vote.")
	}
	// You can vote!

	a := 10
	b := 20
	c := 30

	if a >= b && a >= c {
		fmt.Println("a is the greatest number here!")
	} else if b >= a && b >= c {
		fmt.Println("b is the greatest number here!")
	} else {
		fmt.Println("c is the greatest number here!")
	}
	// c is the greatest number here!

	// NOTE: There is no `while` loop in Go language.

}