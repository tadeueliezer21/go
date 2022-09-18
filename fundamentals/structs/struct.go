package main

import "fmt"

type endereco struct {
	logradouro string
	numero     uint8
}

type usuario struct {
	nome     string
	idade    int8
	endereco endereco
}

func main() {

	fmt.Println("Arquivo structs")

	var u usuario

	u.nome = "Eliezer"
	u.idade = 22

	fmt.Println(u)

	en2 := endereco{"Rua", 20}
	u2 := usuario{"Eliezer", 10, en2}

	fmt.Println(u2)

	u3 := usuario{idade: 14}

	fmt.Println(u3)
}
