package hello

import "fmt"

var Visible bool
var noVisible bool = false

/*
	Mayusculas para exportar
*/
func SayHello(name string) {
	fmt.Println("Hello", name)
}
