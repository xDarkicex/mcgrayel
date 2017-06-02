package controllers

import "github.com/xDarkicex/mcgrayel/helpers"

//Application this is here to pass a type to router
type Application helpers.Controller

func (this Application) Index(a helpers.RouterArgs) {
	obj := map[string]interface{}{}
	helpers.Render(a, "application/index", obj)
}

func (this Application) International(a helpers.RouterArgs) {
	obj := map[string]interface{}{}
	helpers.Render(a, "application/international", obj)

}

type Locate helpers.Controller

type Product struct {
	Name string
}

func (this Locate) Locate(a helpers.RouterArgs) {
	helpers.Render(a, "storeLocator/locator", nil)
}
