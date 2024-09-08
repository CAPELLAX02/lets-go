package main

import "fmt"

func main() {
	a, b := 12, 13

	add(a, b) // 12 + 13 = 25

	fmt.Println(multiply(a, b)) // 156

	fmt.Println(get_squares(a, b)) // 144 169

	var numbers = [] int { 10, 20, 30, 40, 50 }
	fmt.Println(get_sum(numbers)) // 150

	fmt.Println(get_sum_premium(3, 4, 5)) // 12
	fmt.Println(get_sum_premium(1, 2, 3, 4, 5)) // 15
}

func add(x int, y int) {
	fmt.Printf("%d + %d = %d", x, y, x + y)
	fmt.Println()
}

func multiply(x int, y int) int {
	return x * y
}

func get_squares(x int, y int) (int, int) {
	return x * x, y * y
}

func get_sum(arr [] int) int {
	sum := 0
	for _, value := range arr {
		sum += value
	}
	return sum
}

func get_sum_premium(numbers ...int) int {
	sum := 0
	for _, value := range numbers {
		sum += value
	}
	return sum
}