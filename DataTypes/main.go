package main

import (
	//"errors"
	"fmt"
)

func main() {
	var number int16 = 100
	fmt.Println(number)

	var number2 uint64 = 100
	fmt.Println(number2)

	var number3 rune = 9874 // Rune is a alias for int32
	fmt.Println(number3)

	var number4 byte = 123
	fmt.Println(number4)

	char := 'a'
	fmt.Println(char)

	var boolean bool = true
	var boolean2 bool = false
	fmt.Println(boolean, boolean2)

	var erro error
	fmt.Println(erro)

	// var erro2 error = errors.New("Error from GoLang")
	// fmt.Println(erro2)
}
