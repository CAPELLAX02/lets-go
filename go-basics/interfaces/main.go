package main

import "fmt"

type Book struct {
	desi int
}

type Electronic struct {
	desi int
}

type IShippible interface {
	shippingCost() int
}

func (book *Book) shippingCost() int {
	return 5 + (book.desi * 2)
}

func (electronic *Electronic) shippingCost() int {
	return 10 + (electronic.desi * 3)
}

func calculateTotalShippingCostOfBooks(books []Book) int {
	total := 0
	for _, book := range books {
		total += book.shippingCost()
	}
	return total
}

func calculateTotalShippingCostOfElectronics(electronics []Electronic) int {
	total := 0
	for _, electronic := range electronics {
		total += electronic.shippingCost()
	}
	return total
}

func calculateTotalShippingForShippibleProducts(products []IShippible) int {
	total := 0
	for _, product := range products {
		total += product.shippingCost()
	}
	return total
}

func main() {
	var book1 = Book{ desi: 10 }
	book2 := Book{ desi: 20 }

	fmt.Println(book1.shippingCost()) // 25
	fmt.Println(book2.shippingCost()) // 45

	//--------------------------------------------------

	books := []Book{ { desi: 10 }, { desi: 20 } }
	fmt.Println(calculateTotalShippingCostOfBooks(books)) // 70

	//--------------------------------------------------

	electronic1 := Electronic{ desi: 20 }
	fmt.Println(electronic1.shippingCost()) // 70
	
	//--------------------------------------------------

	electronics := []Electronic{ { desi: 10 }, { desi: 20 } }
	fmt.Println(calculateTotalShippingCostOfElectronics(electronics)) // 110

	//--------------------------------------------------

	// Using interfaces:

	var product1 IShippible
	product1 = &Book{ desi: 10 }
	fmt.Println(product1.shippingCost()) // 25

	//------------------------------------------------
  var products2 []IShippible = []IShippible{ &Book{desi: 10}, &Book{desi: 20}, &Electronic{desi: 15} }

	fmt.Println(calculateTotalShippingForShippibleProducts(products2)) // 125 (interface works!)

}