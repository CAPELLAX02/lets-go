package main

import "fmt"

type Customer struct {
	id int64
	name string
	age int
	address Address
}

type Address struct {
	country string
	city    string
}

func (customer Customer) toString() {
	fmt.Printf("Name: %s, Age: %d\n", customer.name, customer.age)
}

func changeName(customer Customer) {
	customer.name = "Changed Name"
}

func changeNameByPassingByReference(customer *Customer) {
	customer.name = "Changed Name fr this time!"
}

func (customer *Customer) changeNameAsMemberOfStruct(newName string) {
	customer.name = newName
}

func (customer *Customer) changeAgeAsMemberOfStruct(newAge int) {
	customer.age = newAge
}

func main() {
	var customer1 = Customer{ id: 1, name: "Ahmet", age: 22 }

	fmt.Println(customer1) // {1 Ahmet 22 { }}

	fmt.Println(customer1.id) // 1
	fmt.Println(customer1.name) // Ahmet
	fmt.Println(customer1.age) // 22

	customer1.name = "John"

	fmt.Println(customer1) // {1 John 22 { }}

	customer1.address.country = "Turkey" 
	customer1.address.city = "Ankara"

	fmt.Println(customer1) // {1 John 22 {Turkey Ankara}}

	fmt.Println()

	var customer2 = Customer{ id: 2, name: "Alex", age: 25, address: Address{ country: "France", city: "Paris" } }

	fmt.Println(customer2) // {2 Alex 25 {France Paris}}

	fmt.Println()

	customer2.toString() // Name: Alex, Age: 25

	fmt.Println()

	customer2.toString() // Name: Alex, Age: 25
	changeName(customer2)
	customer2.toString() // Name: Alex, Age: 25
	// Name did not change here, because we defined the function by passing value. We should have defined it using passing by reference logic. Let's try to pass by reference below.

	changeNameByPassingByReference(&customer2)
	customer2.toString() // Name: Changed Name fr this time!, Age: 25

	fmt.Println()

	var customer3 = Customer{ id: 3, name: "Smith", age: 24, address: Address{ country: "Germany", city: "Berlin" } }
	customer3.toString()
	// Here, we are changing the name, and this function is defined specifically for the Customer struct, we made it clear when defining the function.
	customer3.changeNameAsMemberOfStruct("Hans") // New name is passed in
	customer3.toString()

	customer3.toString() // Name: Hans, Age: 24
	customer3.changeAgeAsMemberOfStruct(45)
	customer3.toString() // Name: Hans, Age: 45

}