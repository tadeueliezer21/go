package main

import "fmt"

func calcMatematico(n1, n2 int) (soma int, sub int) {

	soma = n1 + n2
	sub = n1 - n2

	return

}

func main() {

	fmt.Println(calcMatematico(2, 5))

}
