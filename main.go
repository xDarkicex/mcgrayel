package main

import (
	"flag"
	"log"
	"net/http"
	"github.com/xDarkicex/mcgrayel/datastore"
	"github.com/xDarkicex/mcgrayel/server"

	mgo "gopkg.in/mgo.v2"
)

var (
	session *mgo.Session
)

func init() {
	// Dial mongo Datastore
	err := datastore.Dial()
	if err != nil {
		log.Fatal(err)
	}
	session = datastore.GetSession()
}

// var goHTML *template.Template

// type Product struct {
// 	Name   string
// 	Method string
// }
// type Options struct {
// 	isProduct bool
// 	View      string
// }

// func Parse(opts Options) *Options {
// 	return &Options{
// 		isProduct: opts.isProduct,
// 		View:      opts.View,
// 	}
// }

// func (*Product) setMethod(method string) *Product {
// 	return &Product{
// 		Method: method,
// 	}
// }

// func (*Product) setName(name string) *Product {
// 	return &Product{
// 		Name: name,
// 	}
// }
var Server *server.Server
func init(){
	Server = server.Initialize()
}

func main() {
	port := flag.String("p", "8000", "port to serve on")
	flag.Parse()

	log.Fatal(http.ListenAndServe(":"+*port, Server.Routes))
}
