package main

import "fmt"

type user struct {
	name    string
	age     uint8
	address address
}

type address struct {
	city  string
	state string
}

func main() {
	var user1 user
	user1.name = "John"
	user1.age = 30
	fmt.Println(user1)

	var user2 = user{name: "André", age: 24}
	fmt.Println(user2)

	user_address := address{city: "Recife", state: "PE"}

	user3 := user{"Fonsêca", 24, user_address}
	fmt.Println(user3)

	user4 := user{name: "Borba"}
	fmt.Println(user4)

	user4.age = 24
	fmt.Println(user4)
}
