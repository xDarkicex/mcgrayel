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
	Phone string
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
		&employee{Name: "Marvin Rezac", Title: "C.E.O.", Phone: "(559) 246-3719", Email: "mrezac@easycarewater.com"},
		&employee{Name: "Evangelina Serrano", Title: "President", Phone: "(559) 246-4480", Email: "eserrano@easycarewater.com"},
		&employee{Name: "Monica Morgan", Title: "Accountant", Phone: "(559) 260-9692", Email: "accounting@easycarewater.com"},
		&employee{Name: "Gentry Rolofson", Title: "Director of Information Technology", Phone: "(559) 676-0527", Email: "grolofson@bitdev.io"},
		&employee{Name: "Tiffany Rolofson", Title: "Senior sales Represenitive", Phone: "(559) 694-1267", Email: "trolofson@easycarewater.com"},
		&employee{Name: "Jacqueline Hubbard", Title: "Sales Represenative", Phone: "(559) 321-3857", Email: "jkitchens@easycarewater.com"},
		&employee{Name: "Rosemarie Arenas", Title: "Senior Sales Executive", Phone: "(559) 974-8252", Email: "rarenas@easycarewater.com"},
		&employee{Name: "Scott Nichols", Title: "Sales Represenative", Phone: "(469) 401-2130", Email: "snichols@easycarewater.com"},
		&employee{Name: "Bryan Wiegand", Title: "Sales Executive", Phone: "(210) 857-5729", Email: "bweigand@easycarewater.com"},
		&employee{Name: "Todd Wilson", Title: "Senior Sales Executive", Phone: "(909) 973-3824", Email: "twilson@easycarewater.com"},
		&employee{Name: "Matt Wyant", Title: "Product Specialist", Phone: "(623) 738-5061", Email: "mwyant@easycarewater.com"},
		&employee{Name: "Will Bond", Title: "Product Specialist", Phone: "(619) 208-4148", Email: "wbond@easycarewater.com"},
		&employee{Name: "Vicki Brown", Title: "Product Specialist", Phone: "(619) 208-4148", Email: "vbrown@easycarewater.com"},
		&employee{Name: "Bill Hills", Title: "Sales Represenative", Phone: "(609) 970-0871", Email: "agpoolman@yahoo.com"},
		&employee{Name: "Kyle Morgan", Title: "Product Specialist", Phone: "(480) 442-4646", Email: "kmorgan@easycarewater.com"},
		&employee{Name: "Carlos Mejia", Title: "Product Specialist", Phone: "(786) 369-9903", Email: "cmejia@easycarewater.com"},
		&employee{Name: "Jose Valdovinos", Title: "Factory Sales Represenative", Phone: "(818) 581-0991", Email: "jvaldovinos@easycarewater.com"},
		&employee{Name: "Jospeh Mendez", Title: "Plant Manager", Phone: "", Email: "jmendez@easycarewater.com"},
		&employee{Name: "Pio Trinidad", Title: "Production Employee", Phone: "", Email: ""},
		&employee{Name: "Michael Alvarez", Title: "Shipping Manager", Phone: "", Email: ""},
		&employee{Name: "Jimmy Blajos", Title: "Plant Maintenance Manager", Phone: "", Email: ""},
		&employee{Name: "Debbie Saldivar", Title: "Administrative Assistant", Phone: "(559) 270-7698", Email: "dsaldivar@easycarewater.com"},
		&employee{Name: "Jennifer Weaver", Title: "Order Processor", Phone: "(805) 901-3953", Email: "jweaver@easycarewater.com"},
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
