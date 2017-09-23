package controllers

import (
	"encoding/json"
	"fmt"

	"github.com/xDarkicex/mcgrayel/helpers"
)

//Staff type controller
type Staff helpers.Controller

type employee struct {
	Name  string
	Title string
	Email string
}

//Index list all Staff
func (s Staff) Index(a helpers.RouterArgs) {
	// data := convertStruct2Byte(getStaff())
	a.Response.Header().Set("content-type", "text/html")
	helpers.Render(a, "staff/index", nil)
}

func (s Staff) IndexData(a helpers.RouterArgs) {
	data := convertStruct2Byte(getStaff())
	a.Response.Header().Set("Content-Type", "application/json")
	a.Response.Write(data)
}

func getStaff() interface{} {
	data := []interface{}{
		&employee{Name: "Marvin Rezac", Title: "C.E.O.", Email: "mrezac@easycarewater.com"},
		&employee{Name: "Evangelina Serrano", Title: "President", Email: "eserrano@easycarewater.com"},
		&employee{Name: "Monica Morgan", Title: "Accountant", Email: "accounting@easycarewater.com"},
		&employee{Name: "Gentry Rolofson", Title: "Director of Information Technology", Email: "grolofson@bitdev.io"},
		&employee{Name: "Tiffany Rolofson", Title: "Senior sales Represenitive", Email: "trolofson@easycarewater.com"},
		&employee{Name: "Jacqueline Hubbard", Title: "Sales Represenative", Email: "jkitchens@easycarewater.com"},
		&employee{Name: "Rosemarie Arenas", Title: "Senior Sales Executive", Email: "rarenas@easycarewater.com"},
		&employee{Name: "Scott Nichols", Title: "Sales Represenative", Email: "snichols@easycarewater.com"},
		&employee{Name: "Bryan Wiegand", Title: "Sales Executive", Email: "bweigand@easycarewater.com"},
		&employee{Name: "Todd Wilson", Title: "Senior Sales Executive", Email: "twilson@easycarewater.com"},
		&employee{Name: "Matt Wyant", Title: "Product Specialist", Email: "mwyant@easycarewater.com"},
		&employee{Name: "Will Bond", Title: "Product Specialist", Email: "wbond@easycarewater.com"},
		&employee{Name: "Vicki Brown", Title: "Product Specialist", Email: "vbrown@easycarewater.com"},
		&employee{Name: "Bill Hills", Title: "Sales Represenative", Email: "agpoolman@yahoo.com"},
		&employee{Name: "Kyle Morgan", Title: "Product Specialist", Email: "kmorgan@easycarewater.com"},
		&employee{Name: "Carlos Mejia", Title: "Product Specialist", Email: "cmejia@easycarewater.com"},
		&employee{Name: "Jose Valdovinos", Title: "Factory Sales Represenative", Email: "jvaldovinos@easycarewater.com"},
		&employee{Name: "Jospeh Mendez", Title: "Plant Manager", Email: "jmendez@easycarewater.com"},
		&employee{Name: "Pio Trinidad", Title: "Production Employee", Email: ""},
		&employee{Name: "Michael Alvarez", Title: "Shipping Manager", Email: ""},
		&employee{Name: "Jimmy Blajos", Title: "Plant Maintenance Manager", Email: ""},
		&employee{Name: "Debbie Saldivar", Title: "Administrative Assistant", Email: "dsaldivar@easycarewater.com"},
		&employee{Name: "Jennifer Weaver", Title: "Order Processor", Email: "jweaver@easycarewater.com"},
	}
	return data
}
func convertStruct2Byte(data interface{}) []byte {
	d, err := json.Marshal(data)
	if err != nil {
		fmt.Println(err)
	}
	return d
}
