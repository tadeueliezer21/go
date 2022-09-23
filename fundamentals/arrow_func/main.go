package main

import "fmt"

func main() {

	retorno := func(name string) string {
		return fmt.Sprintf("ALOU => %s", name)
	}("ELIEZER")

	funcao := func() {
		fmt.Println("JOSE MARIA")
	}

	funcao()

	fmt.Println(retorno)
}
