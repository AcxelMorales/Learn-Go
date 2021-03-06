package main

import "fmt"

func main() {
	f()
}

func f() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("%T\n", r)
			fmt.Println("Recuperado en f", r)
		}
	}()

	fmt.Println("Llamando a g")
	g(4)
	fmt.Println("Retorna en g")
}

func g(i int) {
	if i > 3 {
		panic("No puede ser mayor a 3")
	}

	defer fmt.Println("Defer en g")
	fmt.Println("Imprimiendo en g", i)
}
