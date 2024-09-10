package main

import (
	"fmt"
	"mymodule/helper"
	"mymodule/helper/rest"
)

func main() {
	fmt.Println("Main runs")

	helper.Helper1()

	rest.Rest1()
}