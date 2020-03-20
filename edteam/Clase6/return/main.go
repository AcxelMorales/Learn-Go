package main

import "fmt"

func main() {
	fmt.Println(suma(6, 6))
	fmt.Println(typeOld(22))

	n := []int8{45, 60, 100, -5, 12, 43, 99}
	min, max := minAndMax(n)

	fmt.Println("Max", max)
	fmt.Println("Min", min)
}

// se puede omitir el tipado en el primer valor
// si son del mismo las demas
func suma(number1, number2 uint16) uint16 {
	result := number1 + number2
	return result
}

func typeOld(yearsOld uint8) string {
	var typee string

	switch {
	case yearsOld < 12:
		typee = "Childhood"
	case yearsOld < 30:
		typee = "Adolescence"
	default:
		typee = "Maturity"
	}

	return typee
}

func minAndMax(numbers []int8) (int8, int8) {
	var min, max int8

	for _, v := range numbers {
		if v > max {
			max = v
		}

		if v < min {
			min = v
		}
	}

	return min, max
}
