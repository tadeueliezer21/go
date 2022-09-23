package main

import "fmt"

func main() {

	usuario := map[string]string{

		"nome":      "Pedro",
		"sobrenome": "Silva",
	}

	fmt.Println(usuario)

	usuario2 := map[string]map[string]string{

		"nome": {
			"primeiro": "Joao",
			"ultimo":   "Pedro",
		},
	}

	fmt.Println(usuario2)

	delete(usuario, "nome")
	usuario["signo"] = "virgem"

	fmt.Println(usuario)
}
