package main

import "fmt"

func main() {

	numero := 10

	if numero > 15 {
		fmt.Println("MAIOR")
	} else {
		fmt.Println("MENOR")
	}

	//Variavel dentro do escopo
	if outroNumero := numero; outroNumero > 5 {
		fmt.Println("MAIOR")
	} else {
		fmt.Println("MENOR")
	}

}
