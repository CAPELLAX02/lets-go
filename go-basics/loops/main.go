package main

import "fmt"

func main() {
	
	index := 1

	for index <= 10 {
		fmt.Print(index, " ")
		index++
	}
	// 1 2 3 4 5 6 7 8 9 10 

	fmt.Println()

	for i:=1; i <= 10; i++ {
		fmt.Print(i, " ")
	}
	// 1 2 3 4 5 6 7 8 9 10 

	fmt.Println()

	var names = [3] string {"Go", "Programming", "Language"}

	x := 0
	for x < len(names) {
		fmt.Print(names[x], " ")
		x++
	}
	// Go Programming Language 

	fmt.Println()

	for i := 0; i < len(names); i++ {
		fmt.Print(names[i], " ")
	}
	// Go Programming Language 
}