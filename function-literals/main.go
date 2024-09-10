package main

import (
	"fmt"
	"reflect"
)

func main() {

	func() {
		fmt.Println("f1")
	} () // f1


	func(x int, y int) {
		fmt.Println(x + y)
	} (3, 5) // 8

	
	add := func(x int, y int) int {
		return x + y
	}

	fmt.Println(reflect.TypeOf(add)) // func(int, int)

	substract := func(x int, y int) int {
		return x - y
	}

	a, b := calculator(8, 5, add, substract)

	fmt.Println(a, b) // 13  3


}

func calculator(x int, y int, f1 func(int, int) int, f2 func(int, int) int) (int, int) {
	return f1(x, y), f2(x, y)
}
