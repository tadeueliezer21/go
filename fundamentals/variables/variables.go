package main

import "fmt"

func main() {

	var variable1 string = "Variavel 1"
	var variable2 string = "Variavel 2"

	fmt.Println(variable1)
	fmt.Println(variable2)

	var (
		title1 string = "hello"
		title2 string = "ola"
	)

	fmt.Println(title1)
	fmt.Println(title2)

	const constant string = "ola"

	fmt.Println(constant)
}
