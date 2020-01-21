package main

import "fmt"

func main() {
	// Slices
	// 1.- Es un apuntador a un array
	// 2.- Tamaño (no es fijo)
	// 3.- Capacidad

	// Functions
	// make(tipo del dato, tamaño inicial, opcional -> capacidad inicial)

	// var numbers []int
	numbers := make([]int, 0)

	numbers = append(numbers, 3)
	fmt.Printf("Its size is %d with a capacity of %d\n", len(numbers), cap(numbers))
	numbers = append(numbers, 1)
	fmt.Printf("Its size is %d with a capacity of %d\n", len(numbers), cap(numbers))
	numbers = append(numbers, 6)
	fmt.Printf("Its size is %d with a capacity of %d\n", len(numbers), cap(numbers))
	numbers = append(numbers, 9)
	fmt.Printf("Its size is %d with a capacity of %d\n", len(numbers), cap(numbers))
	numbers = append(numbers, 4)
	fmt.Printf("Its size is %d with a capacity of %d\n", len(numbers), cap(numbers))

	fmt.Println(numbers)
}
