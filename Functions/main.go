package main

import "fmt"

func main() {
	result := sum(5, 6)
	fmt.Println(result)

	var f = func() {
		fmt.Println("f Function")
	}
	f()

	var print = func(text string) string {
		return fmt.Sprintf("%s", text)
	}

	text := print("Testing my variable")
	fmt.Println(text)

	result_less, result_division := less_division(5, 6)
	fmt.Println(result_less, result_division)

	just_first, _ := less_division(5, 6)
	fmt.Println(just_first)

	fmt.Println(multiply(10, 10))

	power(1, 2, 3)

	fmt.Println(sumAllNumbers(2, 2, 2, 2, 2))

	func(str string) {
		fmt.Println(str)
	}("Testing my anonymous function")

	fmt.Println(sumAllNumbers(3, 2))
}

func less_division(a, b int8) (int8, int8) {
	var less int8 = a - b
	var division int8 = a / b
	return less, division
}

func multiply(number1, number2 int8) (multiply, power int8) {
	multiply = number1 * number2
	power = number1 ^ number2
	return
}

func sum(number int8, number2 int8) int8 {
	return number + number2
}

func power(numbers ...int) {
	fmt.Println(numbers)
}

func sumAllNumbers(numbers ...int) int {
	var total int = 0

	for _, number := range numbers {
		total += number
	}

	if total != 0 {
		func(number int) int {
			return number
		}(total)
	} // Just to test the anonymous function, that is not necessary!!!
	return total
}
