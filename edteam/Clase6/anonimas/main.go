package main

import "fmt"

func main() {
	func() {
		fmt.Println("Hola !!")
	}()

	anonima := func() {
		fmt.Println("Desde variable")
	}

	anonima()

	sec := secuencia()
	for i := 1; i <= 3; i++ {
		fmt.Println(sec())
	}
}

func secuencia() func() int32 {
	var x int32

	return func() int32 {
		x++
		return x
	}
}
