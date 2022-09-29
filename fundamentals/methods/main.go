package main

import "fmt"

type usuario struct {
	nome  string
	idade uint8
}

func (u usuario) salvar() {
	fmt.Printf("SALVAR %s \n", u.nome)
}

func (u usuario) get_idade() {
	fmt.Println(u.idade)
}

func main() {

	usuario1 := usuario{"Eliezer", 20}
	usuario2 := usuario{"José", 52}

	usuario1.salvar()
	usuario2.salvar()

	usuario1.get_idade()
	usuario2.get_idade()

}
