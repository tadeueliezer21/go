package main

import (
	"command/app"
	"fmt"
	"log"
	"os"
)

func main() {

	fmt.Println("Main")
	application := app.Gerar()

	err := application.Run(os.Args)

	if err != nil {
		log.Fatal(err)
	}
}
