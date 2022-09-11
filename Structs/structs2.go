package main

import "fmt"

type user struct {
	name   string
	age    uint8
	height float32
}

type student struct {
	user
	school string
	state  string
}

func main() {
	student1 := student{user: user{name: "Borba", age: 24}, school: "MIT", state: "MA"}
	student1.height = 1.69
	fmt.Println(student1)

	var student2 student = student{
		user:   user{name: "Fonsêca", age: 24, height: 1.69},
		school: "MIT",
		state:  "MA",
	}

	fmt.Println(student2)
}
