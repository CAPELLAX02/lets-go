package main

import "fmt"

func main() {

	// Define map (key: string, value: int)
	var names map [string] int

	names = make(map[string]int, 0)

	names["Java"] = 1
	names["Go"] = 2
	names["TypeScript"] = 3

	fmt.Println(names) // map[Go:2 Java:1 TypeScript:3]
	fmt.Println(names["Go"]) // 2
	fmt.Println(names["Python"]) // 0 (does not exist in the map)

	fmt.Println()

	//-------------------------------------------

	fruits := map [string] int {
		"Apple": 20,
		"Banana": 30,
		"Peach": 10,
	}

	fmt.Println(fruits) // map[Apple:20 Banana:30 Peach:10]

	delete(fruits, "Banana")

	fmt.Println(fruits) // map[Apple:20 Peach:10]
}