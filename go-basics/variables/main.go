package main

import (
	"fmt"
	"reflect"
)

func main() {

	var s1 string = "string1"
	fmt.Println(s1) // string1
	fmt.Println(reflect.TypeOf(s1)) // string

	fmt.Println()

	var int1 int = 10
	fmt.Println(int1) // 10
	fmt.Println(reflect.TypeOf(int1)) // int

	fmt.Println()
	
	var f1 float32 = 3.2
	fmt.Println(f1) // 3.2
	fmt.Println(reflect.TypeOf((f1))) // float32

	fmt.Println()

	var isSomething bool = false
	fmt.Println(isSomething) // false
	fmt.Println(reflect.TypeOf((isSomething))) // bool

	fmt.Println()

	x := 24
	fmt.Println(x) // 24
	fmt.Println(reflect.TypeOf(x)) // int

	name := "John"
	fmt.Printf("My name is %s\n", name) // My name is John!

}