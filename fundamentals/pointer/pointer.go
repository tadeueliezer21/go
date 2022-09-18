package main

import "fmt"

func main() {
	fmt.Println("Ponteiros")

	var variavel1 int = 10
	var variavel2 int = variavel1

	fmt.Println(variavel1, variavel2)

	variavel1++

	fmt.Println(variavel1, variavel2)

	// Ponteiro referência de memoria

	var variavel3 int
	var ponteiro *int

	variavel3 = 3
	ponteiro = &variavel3

	// desreferenciação
	fmt.Println(*ponteiro)

	variavel3 = 100

	fmt.Println(*ponteiro)
}
