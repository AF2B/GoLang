package main

import "fmt"

func main() {
	slice := make([]float32, 3, 5) // The first param is the type, second is the positions, third is the capacity
	slice[0] = 1.0
	slice[1] = 2.0
	slice[2] = 3.0
	slice = append(slice, 4.0)

	fmt.Println(len(slice))
	fmt.Println(cap(slice))
	fmt.Println(slice)
}

/*
The make method allocates an array with 5 elements and returns a slice of type float32 with 3 elements.
*/
