package main

import "fmt"

func main() {
	a := 6

	fmt.Println("Antes de duplicar, a = ", a)
	// fmt.Println(&a)
	duplicar(&a)
	fmt.Println("Despues de duplicar, a = ", a)
}

func duplicar(x *int) {
	*x = *x * 2
	// fmt.Println("Dentro de la función x = ", x)
	fmt.Println("Dentro de la función x = ", *x)
}
