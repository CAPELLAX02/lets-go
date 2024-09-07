package main

import "fmt"

func main() {

	// `defer` is a keyword used to delay the execution of a function until the surrounding function finishes. 

	// defer logics work like 'stacks'. (LIFO - Last In First Out)

	defer fmt.Println("hello") // printed out 2nd
	fmt.Println("world!") // printed out 1st

	fmt.Println()

	defer fmt.Println("1st defer") // printed out 3rd
	defer fmt.Println("2nd defer") // printed out 2nd
	defer fmt.Println("3rd defer") // printed out 1st

	fmt.Println()

	defer fmt.Println("before the if-statement") // this line printed out and the "after the if-statement never printed out"
	if true {
		return
	}
	defer fmt.Println("after the if-statement")

}