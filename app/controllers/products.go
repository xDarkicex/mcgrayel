package controllers

import (
	"github.com/xDarkicex/mcgrayel/app/models"
	"github.com/xDarkicex/mcgrayel/helpers"
)

//Products this is here to pass a type to router
type Products helpers.Controller

func (products Products) Show(a helpers.RouterArgs) {

	helpers.Render(a, "products/product", map[string]interface{}{})
}

func (products Products) New(a helpers.RouterArgs) {

	helpers.Render(a, "products/new", map[string]interface{}{})
}

func (products Products) Edit(a helpers.RouterArgs) {

	helpers.Render(a, "products/edit", map[string]interface{}{})
}

func (products Products) Create(a helpers.RouterArgs) {
	// var err error
	// File processing ...
	// file, _, _ := a.Request.FormFile("file")
	// fileBytes, _ := ioutil.ReadAll(file)

	// points := strings.Split(a.Request.FormValue("points"), ".")
	// sellingPoints := strings.Split(a.Request.FormValue("sellingpoints"), ".")
	// features := strings.Split(a.Request.FormValue("features"), ".")
	// benefits := strings.Split(a.Request.FormValue("benefits"), ".")

	// product := setProduct()
	// if models.ProductCreate(product, fileBytes) != nil {
	// 	panic(err)
	// }
}
func setProduct(name string, tag string, points []string, summary string, sellingpoints []string, features []string, benefits []string, howitworks string, notice string) *models.Product {
	return &models.Product{
		Name:          name,
		Tag:           tag,
		Points:        points,
		Summary:       summary,
		SellingPoints: sellingpoints,
		Features:      features,
		Benefits:      benefits,
		HowItWorks:    howitworks,
		Notice:        notice,
	}
}
