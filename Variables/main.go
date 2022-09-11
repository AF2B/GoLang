package main

import "fmt"

func main() {
	var variable string = ""
	variable_with_no_types := ""
	fmt.Println(variable, variable_with_no_types)

	var (
		variable2 string = "..."
		variable3 string = "..."
	)
	fmt.Println(variable2, variable3)

	var variable4, variable5 int = 0, 0
	fmt.Println(variable4, variable5)

	const PI float32 = 3.14
	fmt.Println(PI)

	var (
		number1 int = 1
		number2 int = 2
	)

	number1, number2 = number2, number1
	fmt.Println(number1, number2)
}
