package main

import (
	"fmt"
	"time"
)

func main() {
	var indice int

	for indice < 10 {
		indice++
		fmt.Println(indice)
		time.Sleep(time.Second)
	}

	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for indice, value := range arr {
		value += indice
		fmt.Println(value)
	}
}
