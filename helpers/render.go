package helpers

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
)

func Render(a RouterArgs, view string, object map[string]interface{}) {
	device := a.Request.UserAgent()
	expression := regexp.MustCompile("(Mobi(le|/xyz)|Tablet)")
	if !expression.MatchString(device) {
		a.Response.Header().Set("Connection", "keep-alive")
	}
	a.Response.Header().Set("Vary", "Accept-Encoding")
	a.Response.Header().Set("Cache-Control", "private, max-age=7776000")
	a.Response.Header().Set("Transfer-Encoding", "gzip, chunked")
	// var goHTML *template.Template
	templateFilePath := filepath.Join("./app/views/", view+".gohtml")
	fmt.Println(templateFilePath)
	info, err := os.Stat(templateFilePath)
	if err != nil {
		if os.IsNotExist(err) {
			http.NotFound(a.Response, a.Request)
			return
		}
	}
	if info.IsDir() {
		http.NotFound(a.Response, a.Request)
		return
	}
	goHTML := template.Must(template.ParseFiles("./app/views/"+view+".gohtml", "./app/views/layout/includes.gohtml"))
	err = goHTML.Execute(a.Response, object)
	if err != nil {
		log.Println(err.Error())
		http.Error(a.Response, http.StatusText(500), 500)
	}

}
