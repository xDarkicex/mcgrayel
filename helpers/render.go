package helpers

import (
	"html/template"
	"log"
)

func Render(a RouterArgs, view string, object map[string]interface{}) {
	var goHTML *template.Template
	goHTML = template.Must(template.ParseFiles("./app/views/"+view+".gohtml", "./app/views/layout/includes.gohtml"))
	err := goHTML.Execute(a.Response, object)
	if err != nil {
		log.Fatal(err)
	}
}
