package main

import "fmt"

func soma(mult string, numeros ...int) int {
	total := 0

	fmt.Println(mult)

	for _, numero := range numeros {
		total += numero
	}

	return total
}

func main() {

	total := soma("SOMANDO", 2, 4, 6, 7, 8)

	fmt.Println(total)

}
