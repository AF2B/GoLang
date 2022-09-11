package main

import (
	"fmt"
	"reflect"
)

func main() {
	example := map[string]string{
		"a": "apple",
		"b": "banana",
		"c": "cherry",
	}

	fmt.Println(example)

	example2 := map[string]map[string]string{
		"name": {
			"first": "André",
			"last":  "Borba",
		},
		"carrer": {
			"title":   "Software Engineer",
			"company": "ISITics",
		},
	}

	fmt.Println(example2)
	fmt.Println(reflect.TypeOf(example2))
}
