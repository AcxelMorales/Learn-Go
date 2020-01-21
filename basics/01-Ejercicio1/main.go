package main

import "fmt"

// Creación normal de variables
//var num int
//var texto string
//var status bool

func main() {
	lastName, status, yearsOld, height := "Morales", false, 22, 1.69
	name := "Acxel"

	fmt.Println("Hello world " + name)
	fmt.Println(lastName)
	fmt.Println(status)
	fmt.Println(yearsOld)
	fmt.Println(height)
}
