package main

import (
	"os"
	"text/template"
)

type Todo struct {
	Name string
	Description string
}

func main(){

	td:= Todo {"Test templates", "Let's test a template to see the magic."}
	t,err := template.New("todos").Parse("You have a task named \"{{ .Name}}\" with description: \"{{ .Description}}\"")
	 // error checking
    if err != nil {
        panic(err)
    }

    err = t.Execute(os.Stdout, td)
    if err != nil {
        panic(err)
    }

    // it is possible to reuse the same template, without
    // needing to create or parse it again by providing the struct you want
    // to use to the "Execute" function again
    rzfNew := Todo{"Make breakfast", "Grind and brew coffee, and find a croissant"}

    err = t.Execute(os.Stdout, rzfNew) // reuse our same template
    // error checking
    if err != nil {
        panic(err)
    }
}
