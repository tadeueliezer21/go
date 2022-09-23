package main

import (
	"fmt"
	"reflect"
)

func main() {

	var array [5]int

	array[0] = 1

	array2 := [5]string{"a", "b", "c", "d", "e"}

	fmt.Println(array2)

	array3 := [...]int{1, 2, 3}

	fmt.Println(array3)

	slice1 := []int{1, 2, 3, 4}

	fmt.Println(slice1)

	fmt.Println(reflect.TypeOf(array3))
	fmt.Println(reflect.TypeOf(slice1))

	slice1 = append(slice1, 1)
	slice1 = append(slice1, 2)
	slice1 = append(slice1, 3)

	fmt.Println(slice1)

	slice2 := array2[0:2]

	fmt.Println(slice2)

	array2[1] = "mudou"

	fmt.Println(array2)
	fmt.Println(slice2)

}
