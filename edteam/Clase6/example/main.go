package main

import "fmt"

func main() {
	hello("Acxel", 22)
}

func hello(name string, yearsOld uint8) {
	fmt.Printf("Hola %s tienes %d años\n", name, yearsOld)

	if yearsOld >= 30 {
		return
	}

	fmt.Println("Eres joven")
}
