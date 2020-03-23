package main

import (
	"fmt"
)

type Persona struct {
	nombre string
	edad   int8
}

func main() {
	var persona Persona
	persona.saludar("Acxel")
	persona.setEdad(22)
	persona.setNombre("Acxel")
	fmt.Println(persona)
}

func (p Persona) saludar(nombre string) {
	fmt.Println("Hola", nombre)
}

func (p *Persona) setEdad(edad int8) {
	if edad >= 0 {
		p.edad = edad
	} else {
		fmt.Println("Valor negativo")
	}
}

func (p *Persona) setNombre(nombre string) {
	if nombre == "" {
		fmt.Println("Sin nombre")
	}

	p.nombre = nombre
}
