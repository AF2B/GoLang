package main

import (
	"fmt"
	"reflect"
)

func main() {
	var arr [3]int
	arr[0] = 1
	arr[1] = 2
	arr[2] = 3

	arr2 := [6]string{"a", "b", "c", "d", "e", "f"}
	fmt.Println(arr2)

	slice := []int{5, 8}
	fmt.Println(reflect.TypeOf(slice))
	fmt.Println(reflect.TypeOf(arr))

	slice = append(slice, 6)
	fmt.Println(slice)

	take := arr2[0:1]
	fmt.Println(take)
}
