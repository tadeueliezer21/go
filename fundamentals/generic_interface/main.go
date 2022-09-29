package main

import "fmt"

func generica(interf interface{}) {
	fmt.Println(interf)
}

func main() {

	generica(1)
	generica(true)
	generica("abc")

	mapa := map[interface{}]interface{}{
		1:       "abc",
		"float": true,
		true:    'C',
	}

	fmt.Println(mapa)
}
