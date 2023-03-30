package main

import "fmt"

// function that takes an integer argument by value
func passByValue(num int) {
	num = num + 10
	fmt.Println("Inside passByValue function, num = ", num)
}

// function that takes a pointer to an integer argument by reference
func passByReference(numPtr *int) {
	*numPtr = *numPtr + 10
	fmt.Println("Inside passByReference function, *numPtr = ", *numPtr)
}

func main() {
	// declaring and initializing an integer variable
	num := 5

	// calling passByValue function with num argument
	passByValue(num)
	fmt.Println("After calling passByValue function, num = ", num)

	// calling passByReference function with address of num variable as argument
	passByReference(&num)
	fmt.Println("After calling passByReference function, num = ", num)
}
