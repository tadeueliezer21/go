package main

import "fmt"

func f1() {
	fmt.Println("Executando a funcao 1")
}

func f2() {
	fmt.Println("Executando a funcao 2")
}

func alunoEstaAprovado(n1, n2 float32) bool {

	fmt.Println("Verificando aluno aprovado")

	media := n1 + n2

	if media >= 6 {
		return true
	}

	return false

}

func main() {
	// defer == adiar
	defer f1()
	f2()

	alunoEstaAprovado(4.45, 6.53)
}
