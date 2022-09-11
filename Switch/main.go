package main

import "fmt"

func main() {
	fmt.Println(dayOfWeek(1))
}

func dayOfWeek(number int) string {
	var dayOfWeek string

	switch {
	case number == 1:
		dayOfWeek = "Domingo"
		fallthrough
	case number == 2:
		dayOfWeek = "Monday"
	case number == 3:
		dayOfWeek = "Tuesday"
	case number == 4:
		dayOfWeek = "Wednesday"
	case number == 5:
		dayOfWeek = "Thursday"
	case number == 6:
		dayOfWeek = "Friday"
	case number == 7:
		dayOfWeek = "Saturday"
	default:
		dayOfWeek = "Invalid number"
	}

	return dayOfWeek
}
