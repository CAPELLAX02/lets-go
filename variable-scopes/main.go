package main

import "fmt"

// global scope
var x = 30

func some_function() {
	// functional scope
	x := 20
	fmt.Println(x)
}

func main() {

	var condition = true

	if condition {
		// block scope
		var x = 10 
		fmt.Println(x)
	}

	// fmt.Println(x) // Error: undefined: x

}