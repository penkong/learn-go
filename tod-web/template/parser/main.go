package main

import (
	"os"
	"log"
	"text/template"
)

func main()  {
	tpl, err := template.ParseFiles("tp.gohtml")
	if err != nil {
		log.Fatalln(err)
	}

	err = tpl.Execute(os.Stdout, "sdfsdf")
	if err != nil {
		log.Fatalln(err)
	}
}

