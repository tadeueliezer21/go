package main

import "fmt"

func somar(n1 int8, n2 int8) int8 {
	return n1 + n2
}

func calculos(n1, n2 int8) (int8, int8) {
	soma := n1 + n2
	subtracao := n1 - n2

	return soma, subtracao
}

func main() {
	soma := somar(2, 4)

	fmt.Println(soma)

	var f = func() {
		fmt.Println("F")
	}

	f()

	resultadoSoma, _ := calculos(4, 4)

	fmt.Println(resultadoSoma)
}
