package main

import "fmt"

var status bool
var option int

func main() {
	status = false

	if status = true; status {
		fmt.Println(true)
	} else {
		fmt.Println(false)
	}

	switch option = 9; option {
	case 1:
		fmt.Println(1)
	case 2:
		fmt.Println(2)
	case 3:
		fmt.Println(3)
	case 4:
		fmt.Println(4)
	case 5:
		fmt.Println(5)
	default:
		fmt.Println(nil)
	}
}
