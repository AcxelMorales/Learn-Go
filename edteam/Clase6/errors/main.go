package main

import (
	"errors"
	"fmt"
)

func main() {
	// err := errors.New("Mensaje de error")
	// fmt.Println(err)

	r, err := division(100, 0)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println(r)
}

func division(dividendo int, divisor int) (resultado int, err error) {
	if divisor == 0 {
		err = errors.New("No se puede dividir entre 0")
	} else {
		resultado = dividendo / divisor
	}

	return
}
