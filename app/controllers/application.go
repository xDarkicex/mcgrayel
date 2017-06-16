package controllers

import "github.com/xDarkicex/mcgrayel/helpers"

//Application this is here to pass a type to router
type Application helpers.Controller

func (this Application) Index(a helpers.RouterArgs) {
	data := map[string]interface{}{}
	helpers.Render(a, "application/index", data)
}

func (this Application) International(a helpers.RouterArgs) {
	data := map[string]interface{}{}
	helpers.Render(a, "application/international", data)

}

func (this Application) About(a helpers.RouterArgs) {
	//Changing obj to data will update all programming and delete this comment in future
	// The reasoning is data is a better disc.
	data := map[string]interface{}{}
	helpers.Render(a, "application/about", data)
}

func (this Application) Order(a helpers.RouterArgs) {
	data := map[string]interface{}{}
	helpers.Render(a, "application/order", data)
}

type Locate helpers.Controller

func (this Locate) Locate(a helpers.RouterArgs) {
	helpers.Render(a, "storeLocator/locator", nil)
}
