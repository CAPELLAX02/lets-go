package main

import "fmt"

func main() {
	
	// Empty array definition with 3 element capacity
	var names [3] string

	fmt.Println(names) // [ ]

	names[0] = "John"
	names[1] = "Jane"
	names[2] = "David"

	fmt.Println(names) // [John Jane David]

	//--------------------------------------------

	// Array definition with values in it
	var fruits = [3] string { "Banana", "Apple", "Cherry" }

	fmt.Println(fruits) // [Banana Apple Cherry]

	fmt.Println(fruits[1]) // Apple

	fruits[0] = "Watermelon"

	fmt.Println(fruits) // [Watermelon Apple Cherry]
	
	fmt.Println(fruits[0:2]) // [Watermelon Apple]
	fmt.Println(fruits[1:3]) // [Apple Cherry]

	// fmt.Println(fruits[3]) // invalid argument: index 3 out of bounds 

	//--------------------------------------------

	// Slice definition with values in it
	var cities = [] string {"Ankara", "Istanbul", "Izmir"}

	fmt.Println(cities) // [Ankara Istanbul Izmir]

	cities = append(cities, "Balikesir")

	fmt.Println(cities) // [Ankara Istanbul Izmir Balikesir]


}