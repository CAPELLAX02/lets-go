package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {

	var x int
	var y float32
	var z bool
	var ptr *int

	fmt.Println(x) // 0
	fmt.Println(y) // 0
	fmt.Println(z) // false
	fmt.Println(ptr) // <nil>

	//********************************************

	fileData, err := os.ReadFile("sample.txt")

	if (err != nil) {
		fmt.Println("error occured:", err)
	} else {
		fmt.Println(fileData)
	}

	//********************************************

	result, err  := divide(10, 0)

	if err != nil {
		fmt.Println("Division error:", err)
	} else {
		fmt.Println("Division result:", result)
	}

	//********************************************

	productService := ProductService{}

	productErr := productService.Add(Product{id: 1, name: "washing machine", price: 2000})

	if productErr != nil {
		fmt.Println(productErr)
	}

}

func divide(x int, y int) (int, error) {
	if y == 0 {
		return 0, errors.New("denominator cannot be zero.")
	}
	return x / y, nil
}

type Product struct {
	id int
	name string
	price int
}

type ProductService struct {

}

func(productService *ProductService) Add(product Product) error {
	if len(product.name) == 0 {
		return errors.New("Product name cannot be empty.")
	}
	if product.price < 0 {
		return errors.New("Product price must be nonnegative.")
	} 
	fmt.Println("Product service successfull.")
	return nil
}