package main

import "fmt"

func main() {
	var x int = 10
	fmt.Println(x)
	var y *int = &x
	fmt.Println(y)
	var z int = *y
	fmt.Println(z)
	fmt.Println("")

	copy := invertSignalNot(x)
	fmt.Println(copy)
	fmt.Println(x)

	fmt.Println("")

	invertSignal(&x)
	fmt.Println(x)
}

func invertSignal(number *int) {
	*number = *number * -1
}

func invertSignalNot(number int) int {
	number = number * -1
	return number
}
