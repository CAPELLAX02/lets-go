package main

import "fmt"

func main() {

	x := 24
	p := &x
	
	fmt.Println(x) // 24
	fmt.Println(p) // 0xc0000120e0 (address of x in the memory)

	*p = 32

	fmt.Println(*p) // 32
	fmt.Println(x) // 32
	
	fmt.Println(&p) // 0xc000070040 (address of the pointer of the x in the memory)

	// NOTE: *&x === &*x === x are identical

	//-------------------------------

	var numbers = []int { 12, 41, 7, 23, 15, 19 }

	fmt.Println(numbers) // [12 41 7 23 15 19]

	// Passing in the address of the numbers array because we constructed the bubble sort function using pointers
	bubble_sort(&numbers)

	fmt.Println(numbers) // [7 12 15 19 23 41]

}

// Bubble sort using pointers (actually changing the value references in the memory)
func bubble_sort(arr *[]int) {
	for i := 0; i < len(*arr)-1; i++ {
		for j := 0; j < len(*arr)-i-1; j++ {
			if (*arr)[j] > (*arr)[j+1] {
				(*arr)[j], (*arr)[j+1] = (*arr)[j+1], (*arr)[j]
			}
		}
	}
}