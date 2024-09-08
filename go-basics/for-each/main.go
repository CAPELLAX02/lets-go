package main

import "fmt"

func main() {

	var numbers = [] int { 1, 2, 3, 4 }

	for i := 0; i < len(numbers); i++ {
		fmt.Print(numbers[i], " ")
	}
	// 1 2 3 4

	fmt.Println()

	for index, value := range numbers {
		fmt.Println("value of the index",index,"is",value)
	}
	// value of the index 0 is 1
	// value of the index 1 is 2
	// value of the index 2 is 3
	// value of the index 3 is 4
	
	fmt.Println()

	var language = "Golang"

	for _, char := range language {
		fmt.Println(char)
	}
	// 71
	// 111
	// 108
	// 97
	// 110
	// 103

	fmt.Println()

	for _, char := range language {
		fmt.Printf("%c", char)
	}
	// Golang

	fmt.Println()

	var names = map [string] int {
		"John": 10,
		"Jane": 20,
		"Smith": 30,
	}

	for key, value := range names {
		fmt.Printf("key: %s, value: %d\n", key, value)
	}
	// key: Jane, value: 20
	// key: Smith, value: 30
	// key: John, value: 10

	fmt.Println()

}