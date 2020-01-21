package main

import "fmt"

type Persona struct {
	Nombre  string
	Edad    uint8
	IsAlive bool
	Friends []Persona
}

func main() {
	/*
		Aplica la restricción de coma
		puede tener llaves de valor o no (Nombre, Edad, )
	*/
	// var persona1 Persona
	persona1 := Persona{
		"Acxel",
		22,
		true,
		[]Persona{
			Persona{
				"Pedro",
				23,
				true,
				nil,
			},
			Persona{
				"Daniel",
				26,
				true,
				nil,
			},
		},
	}

	// persona1.Nombre = "Acxel"
	// persona1.Edad = 22
	// persona1.IsAlive = true

	fmt.Println(persona1)
}
