package main

import (
	"errors"
	"fmt"
)

func main() {

	// int8, int16, int32, int64
	var numero1 int8 = 127
	fmt.Println(numero1)

	// uint
	var numero2 uint8 = 255
	fmt.Println(numero2)

	// alias
	// int32 = RUNE
	var numero3 rune = 1000000000
	fmt.Println(numero3)

	// byte == 8 bits
	// byte == uint8
	var numero4 byte = 255
	fmt.Println(numero4)

	var numero5 float32 = 4124.32
	fmt.Println(numero5)

	var numero6 float64 = 34.32
	fmt.Println(numero6)

	var numero7 bool = true
	fmt.Println(numero7)

	var erro error = errors.New("Novo erro")
	fmt.Println(erro)
}
