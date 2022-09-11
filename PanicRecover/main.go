package main

import "fmt"

func recoverExecution() {
	if recov := recover(); recov != nil {
		fmt.Println("The program execution is recovered with succefully!")
	}
}

func median(number1, number2 float64) bool {
	defer recoverExecution()
	median := (number1 + number2) / 2

	if median > 6 {
		return true
	} else if median < 6 {
		return false
	}

	panic("The median for this is exactly 6!")
}

func main() {
	fmt.Println(median(6, 6))
	fmt.Println("After execution!")
}
