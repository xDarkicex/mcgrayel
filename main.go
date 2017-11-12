package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/xDarkicex/mcgrayel/server"
)

//Server pointer
var Server *server.Server

func init() {
	Server = server.NewServer()
}

func main() {
	port := flag.String("p", "8000", "port to serve on")
	flag.Parse()
	log.Fatal(http.ListenAndServe(":"+*port, Server.Routes))
}
