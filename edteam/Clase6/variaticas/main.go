package main

import "fmt"

func main() {
	saludarVarios(22, "José", "Acxel", "Manuel", "Jeronimo", "Samuel")
}

func saludarVarios(edad uint8, nombres ...string) {
	fmt.Printf("%T\n", nombres)

	for c, i := range nombres {
		fmt.Println("Hola", i, "- Edad:", edad, "- Index:", c)
	}
}
