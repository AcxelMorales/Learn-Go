package main

import "fmt"

func main() {
	dividir(3, 1)
}

func dividir(dividendo, divisor int) {
	defer fmt.Println("Siempre se ejecuta")

	if divisor == 0 {
		panic("No se puede dividir entre cero")
	}

	fmt.Println(dividendo / divisor)
}
