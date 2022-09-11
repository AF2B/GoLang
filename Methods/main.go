package main

import "fmt"

type user struct {
	name string
	age  string
}

func (u user) greetings() {
	fmt.Printf("Hello, my name is %s and i have %s years old", u.name, u.age)
} // That method just working on the user struct, but we can do it on any struct that we want.

func main() {
	var u user = user{name: "André Borba", age: "24"}
	u.greetings()
}
