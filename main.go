package main

import (
	"flag"
	"fmt"
	"net/http"
	"time"

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
	server := &http.Server{
		Addr:           ":" + *port,
		Handler:        Server.Routes,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	err := server.ListenAndServe()
	if err != nil {
		fmt.Println(err)
	}
}
