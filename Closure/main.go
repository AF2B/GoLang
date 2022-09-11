package main

import "fmt"

func main() {
	printerMain := closure()
	printerMain()
}

func closure() func() {
	text := "Testing closure function"

	printer := func() {
		fmt.Println(text)
	}

	return printer
}
