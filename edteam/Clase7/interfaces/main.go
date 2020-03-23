package main

import "./animales"

func main() {
	var p animales.Perro
	p.Nombre = "Kalvin"
	adoptarMascota(p)

	var g animales.Gato
	g.Nombre = "Julien"
	adoptarMascota(g)
}

// func adoptarPerro(p animales.Perro) {
// 	p.Comunicarse()
// }

// func adoptarGato(g animales.Gato) {
// 	g.Comunicarse()
// }

func adoptarMascota(m animales.Mascota) {
	m.Comunicarse()
}
