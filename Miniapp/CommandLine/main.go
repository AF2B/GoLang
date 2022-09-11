package main

import (
	"command-line/app"
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println("Looking for some public IP's")

	apping := app.Generate()
	erro := apping.Run(os.Args)
	if erro != nil {
		log.Fatal(erro)
	}
}
