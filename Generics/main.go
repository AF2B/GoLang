package main

import "fmt"

func generic(interface{}) {
	fmt.Println("generic")
}

func main() {
	generic(1)
	generic("hello")
	generic(true)
}
