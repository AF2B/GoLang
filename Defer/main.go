package main

import "fmt"

func main() {
	fmt.Println(median(7.0, 7.5, 4.5))
}

func median(numbers ...float32) float32 {
	defer fmt.Println("Finished calculating median.")
	var totalOfNumbers float32 = 0.0
	var result float32 = 0.0

	for _, number := range numbers {
		totalOfNumbers += number
	}
	result = totalOfNumbers / float32(len(numbers))

	return result
}
