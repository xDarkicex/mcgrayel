package controllers

import "github.com/xDarkicex/mcgrayel/helpers"

//International this is here to pass a type to router
type International helpers.Controller

func (i International) Index(a helpers.RouterArgs) {
	data := map[string]interface{}{"name": "International"}
	helpers.Render(a, "international/international", data)
}

func (i International) Greek(a helpers.RouterArgs) {
	data := map[string]interface{}{"name": "Ελληνικά"}
	helpers.Render(a, "international/greek", data)
}

func (i International) Russia(a helpers.RouterArgs) {
	data := map[string]interface{}{"name": "русский"}
	helpers.Render(a, "international/russia", data)
}
