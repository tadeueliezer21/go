package main

import "fmt"

func recuperar() {

	fmt.Println("tentando recuperar")

	if r := recover(); r != nil {
		fmt.Println("Funcao recuperada")
	}

}

func alunoEstaAprovado(n1, n2 float32) bool {
	defer recuperar()
	fmt.Println("Verificando aluno aprovado")

	media := (n1 + n2) / 2

	if media > 6 {
		return true
	} else if media < 6 {
		return false
	}

	panic("Media é exatamente 6")

}

func main() {
	alunoEstaAprovado(6, 6)
	fmt.Println("OPA")
}
