package main

import "fmt"

func closure() func() {

	texto := "dentro da funcao closure"

	funcao := func() {
		fmt.Println(texto)
	}

	return funcao

}

func main() {

	texto := "dentro da func main"
	fmt.Println(texto)

	funcNova := closure()

	funcNova()

}
