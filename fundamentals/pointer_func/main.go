package main

import "fmt"

func invertSinal(numero int) int {

	return numero * -1

}

func invertSinalPointer(numero *int) {
	*numero = *numero - 1
}

func main() {

	numero := 20

	// numeroInvertido := invertSinal(numero)

	fmt.Println(numero)
	fmt.Println(&numero)
	invertSinalPointer(&numero)
	fmt.Println(numero)

}
