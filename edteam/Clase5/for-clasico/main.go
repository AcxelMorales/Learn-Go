package main

import "fmt"

func main() {
	/*
		// ++ --
		for i := 0; i < 5; i++ {
			fmt.Println("Value", i)
		}

		fmt.Println("********************************")

		for j := 4; j >= 0; j-- {
			fmt.Println("Value", j)
		}
	*/

	/*
		// continue, break
		for i := 0; i < 5; i++ {
			if i == 2 {
				fmt.Println("CONTINUE")
				continue
			}

			fmt.Println(i)
		}

		for i := 0; i < 5; i++ {
			if i == 2 {
				fmt.Println("BREAK")
				break
			}

			fmt.Println(i)
		}
	*/

	// var matriz [3][3]uint8
	matriz := [3][3]uint8{}
	value := 0

	for i := 0; i < len(matriz); i++ {
		for j := 0; j < len(matriz); j++ {
			value++
			matriz[i][j] = uint8(value)
		}
	}

	fmt.Println(matriz)

	for row := 0; row < len(matriz); row++ {
		for column := 0; column < len(matriz); column++ {
			fmt.Print(matriz[row][column])

			if matriz[row][column] == 3 || matriz[row][column] == 6 || matriz[row][column] == 9 {
				fmt.Println()
			}
		}
	}

}
