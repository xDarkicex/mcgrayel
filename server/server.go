package server

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/xDarkicex/mcgrayel/app/controllers"
	"github.com/xDarkicex/mcgrayel/helpers"
)

type Server struct {
	Routes *httprouter.Router
}

//Initialize will handle new httprouter servemux
func Initialize() *Server {
	return &Server{Routes: GetRoutes()}
}

func route(fn helpers.RoutesHandler) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		a := helpers.RouterArgs{Request: r, Response: w, Params: ps}
		fn(a)
	}
}

// GetRoutes returns pointer to Router
func GetRoutes() *httprouter.Router {
	router := httprouter.New()
	application := controllers.Application{}
	router.GET("/", route(application.Index))
	router.GET("/international", route(application.International))

	products := controllers.Products{}
	router.GET("/product/:name", route(products.Show))
	router.GET("/products/new", route(products.New))
	router.POST("/products/new", route(products.Create))
	router.GET("/product/:name/edit", route(products.Edit))

	fileServer := http.FileServer(http.Dir("public"))
	router.GET("/static/*filepath", func(res http.ResponseWriter, req *http.Request, pm httprouter.Params) {
		res.Header().Set("Vary", "Accept-Encoding")
		res.Header().Set("Cache-Control", "public, max-age=7776000")
		req.URL.Path = pm.ByName("filepath")
		fileServer.ServeHTTP(res, req)
	})

	return router
}
