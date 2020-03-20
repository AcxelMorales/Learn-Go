package main

import (
	"fmt"

	"./hello"
)

func main() {
	hello.SayHello("Acxel")

	hello.Visible = true
	fmt.Println(hello.Visible)
}
