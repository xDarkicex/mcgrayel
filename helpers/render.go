package helpers

import (
	"html/template"
	"log"
)

var goHTML *template.Template

func Render(a RouterArgs, view string, object map[string]interface{}) {
	goHTML = template.Must(template.ParseFiles("./app/views/"+view+".gohtml", "./app/views/layout/includes.gohtml"))
	err := goHTML.Execute(a.Response, object)
	if err != nil {
		log.Fatal(err)
	}
}
