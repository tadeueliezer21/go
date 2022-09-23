package main

import "fmt"

func main() {

	slice := make([]string, 1)

	fmt.Println(slice)

	slice = append(slice, "a")

	fmt.Println(slice)
	fmt.Println(cap(slice))
	fmt.Println(len(slice))
}
