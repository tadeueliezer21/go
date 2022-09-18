package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
	idade     uint8
	altura    uint8
}

type estudante struct {
	pessoa

	curso     string
	faculdade string
}

func main() {

	var aluno estudante

	aluno.nome = "Eliezer"
	aluno.sobrenome = "Tadeu"
	aluno.idade = 22
	aluno.faculdade = "FATEC"
	aluno.curso = "ADS"
	aluno.altura = 182

	fmt.Println(aluno)
}
