package main

import (
	"fmt"
	"modules/assistant"

	"github.com/badoux/checkmail"
)

func main() {
	fmt.Println("New project with GoLang!")
	assistant.Aux_function()

	erro := checkmail.ValidateFormat("dev_andre.filipe@outlook.com")
	fmt.Println(erro)
}

/*
That project was created by running the following command: go mod init modules.
And after that, we can run the following command to build the project: go build.
And he builds your project and places the name you gave it in the new archive file.
That new archive is Binary. And how can i run it? Simple, just run: ./modules.

GoLang, doesn't a language oriented to the object programming.
So, if you need to declare a "public, protected or private", you do this using the first letter of
name of the variable or method.

Functions with the first letter lowercase are private.

Another question is, private methods inside an package are visible only in the package.
If you want to use an private method wich is in another package, you don't need to import it.
just use! But if you want to use an private method outside his package, then you need to import.

To install a module, i use the following command: go get <module_name>.
*/
