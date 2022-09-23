package main

import (
	"fmt"
	"time"
)

func main() {

	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	nomes := []string{"JOAO", "JOSE", "MARIA"}

	for _, nome := range nomes {
		fmt.Println(nome)
	}

	for indice, letra := range "PALAVRA" {
		fmt.Println(indice, string(letra))
	}

	usuario := map[string]string{
		"nome":      "Leo",
		"sobrenome": "José",
	}

	for chave, valor := range usuario {
		fmt.Println(chave, valor)
	}

	for {
		fmt.Println("Executando")
		time.Sleep(time.Second)
	}
}
