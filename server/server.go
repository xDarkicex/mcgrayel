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

//NewServer will handle new httprouter servemux
func NewServer() *Server {
	return &Server{Routes: GetRoutes()}
}

// route transforms function into Handle
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
	router.HEAD("/", route(application.Index))
	router.GET("/about-us", route(application.About))
	router.GET("/order", route(application.Order))
	router.GET("/contact-us", route(application.Contact))
	router.POST("/contact-us", route(application.Contact))
	router.GET("/sds", route(application.SDS))
	router.GET("/video-tour", route(application.VideoTour))

	intenernational := controllers.International{}
	router.GET("/international", route(intenernational.Index))
	router.GET("/international/greek", route(intenernational.Greek))
	router.GET("/international/russia", route(intenernational.Russia))

	staff := controllers.Staff{}
	router.GET("/staff", route(staff.Index))
	router.GET("/api/staff", route(staff.IndexData))

	products := controllers.Products{}
	router.GET("/product/:name", route(products.Show))
	router.GET("/products/new", route(products.New))
	router.POST("/products/new", route(products.Create))
	router.GET("/product/:name/edit", route(products.Edit))

	location := controllers.Locate{}
	router.GET("/locator", route(location.Locate))

	fileServer := http.FileServer(http.Dir("public"))
	router.GET("/static/*filepath", func(res http.ResponseWriter, req *http.Request, pm httprouter.Params) {
		res.Header().Set("Vary", "Accept-Encoding")
		res.Header().Set("Cache-Control", "public, max-age=7776000")
		req.URL.Path = pm.ByName("filepath")
		fileServer.ServeHTTP(res, req)
	})

	return router
}
