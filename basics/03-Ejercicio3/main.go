package main

import (
	"bufio"
	"fmt"
	"os"
)

var number1 int
var number2 int
var result int
var desc string

func main() {
	fmt.Print("Type number1: ")
	fmt.Scanf("%d", &number1)

	fmt.Print("Type number2: ")
	fmt.Scanf("%d", &number2)

	fmt.Print("Description: ")
	scanner := bufio.NewScanner(os.Stdin)

	if scanner.Scan() {
		desc = scanner.Text()
	}

	result = number1 + number2
	fmt.Println(desc, result)
}
