package main

import "fmt"

func main() {
	// var names [3]string // pude tener tamaño fijo o no
	// names[0] = "Acxel"
	// names[1] = "José"
	// names[2] = "Maria"

	names := [3]string{"Acxel", "José", "Maria"}
	size := len(names)

	fmt.Println(names)
	fmt.Println("Size:", size)
	fmt.Println(names[0:2])
	fmt.Println(names[1:])
	fmt.Println(names[:2])
	fmt.Println(names[0 : len(names)-1])
}
