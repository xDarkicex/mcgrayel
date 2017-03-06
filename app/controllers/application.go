package controllers

import "github.com/xDarkicex/mcgrayel/helpers"

//Application this is here to pass a type to router
type Application helpers.Controller

func (app Application) Index(a helpers.RouterArgs) {
	obj := map[string]interface{}{}
	helpers.Render(a, "application/index", obj)
}

func (app Application) International(a helpers.RouterArgs) {
	obj := map[string]interface{}{}
	helpers.Render(a, "application/international", obj)

}
