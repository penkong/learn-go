package main

import (
	"os"
	"log"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tp.gohtml"))
}

func main()  {
	err := tpl.ExecuteTemplate(os.Stdout, "tp.gohtml", ` release focus `)
	if err != nil {
		log.Fatalln(err)
	}
}

