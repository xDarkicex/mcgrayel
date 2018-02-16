package controllers

import (
	"encoding/json"
	"fmt"

	"github.com/xDarkicex/mcgrayel/helpers"
)

//Staff type controller
type Staff helpers.Controller

type employee struct {
	Name        string
	Title       string
	Email       string
	Phone       string
	Experience  []string
	Description string
	Location
}

type Location struct {
	City    string
	State   string
	Country string
}

//Index list all Staff
func (s Staff) Index(a helpers.RouterArgs) {
	a.Response.Header().Set("content-type", "text/html")
	helpers.Render(a, "staff/index", nil)
}

func (s Staff) IndexData(a helpers.RouterArgs) {
	data := getStaff()
	a.Response.Header().Set("Content-Type", "application/json")
	a.Response.Write(data)
}

func getStaff() []byte {
	data := []employee{
		{Name: "Marvin Rezac", Title: "C.E.O.", Phone: "(559) 246-3719", Email: "mrezac@easycarewater.com", Location: Location{City: "Fresno", State: "CA", Country: "United States"}},
		{Name: "Evangelina Serrano", Title: "President", Phone: "(559) 246-4480", Email: "eserrano@easycarewater.com", Location: Location{City: "Fresno", State: "CA", Country: "United States"}},
		{Name: "Monica Morgan", Title: "Accountant", Phone: "(559) 260-9692", Email: "accounting@easycarewater.com", Location: Location{City: "Fowler", State: "CA", Country: "United States"}},
		{Name: "Gentry Rolofson", Title: "Technomancer", Phone: "(559) 676-0527", Email: "grolofson@bitdev.io", Location: Location{City: "Fresno", State: "CA", Country: "United States"}},
		{Name: "Tiffany Rolofson", Title: "Senior sales Represenitive", Phone: "(559) 694-1267", Email: "trolofson@easycarewater.com", Location: Location{City: "Fresno", State: "CA", Country: "United States"}},
		{Name: "Jacqueline Hubbard", Title: "Sales Represenative", Phone: "(559) 321-3857", Email: "jkitchens@easycarewater.com", Location: Location{City: "Fresno", State: "CA", Country: "United States"}},
		{Name: "Rosemarie Arenas", Title: "Senior Sales Executive", Phone: "(559) 974-8252", Email: "rarenas@easycarewater.com"},
		{Name: "Scott Nichols", Title: "Sales Represenative", Phone: "(469) 401-2130", Email: "snichols@easycarewater.com"},
		{Name: "Bryan Wiegand", Title: "Sales Executive", Phone: "(210) 857-5729", Email: "bweigand@easycarewater.com"},
		{Name: "Todd Wilson", Title: "Senior Sales Executive", Phone: "(909) 973-3824", Email: "twilson@easycarewater.com"},
		{Name: "Matt Wyant", Title: "Product Specialist", Phone: "(623) 738-5061", Email: "mwyant@easycarewater.com"},
		{Name: "Will Bond", Title: "Product Specialist", Phone: "(619) 208-4148", Email: "wbond@easycarewater.com"},
		{Name: "Vicki Brown", Title: "Product Specialist", Phone: "(619) 208-4148", Email: "vbrown@easycarewater.com"},
		{Name: "Bill Hills", Title: "Sales Represenative", Phone: "(609) 970-0871", Email: "agpoolman@yahoo.com"},
		{Name: "Kyle Morgan", Title: "Product Specialist", Phone: "(480) 442-4646", Email: "kmorgan@easycarewater.com"},
		{Name: "Carlos Mejia", Title: "Product Specialist", Phone: "(786) 369-9903", Email: "cmejia@easycarewater.com"},
		{Name: "Jose Valdovinos", Title: "Factory Sales Represenative", Phone: "(818) 581-0991", Email: "jvaldovinos@easycarewater.com"},
		{Name: "Jospeh Mendez", Title: "Plant Manager", Phone: "", Email: "jmendez@easycarewater.com"},
		{Name: "Pio Trinidad", Title: "Production Employee", Phone: "", Email: ""},
		{Name: "Michael Alvarez", Title: "Shipping Manager", Phone: "", Email: ""},
		{Name: "Jimmy Blajos", Title: "Plant Maintenance Manager", Phone: "", Email: ""},
		{Name: "Debbie Saldivar", Title: "Administrative Assistant", Phone: "(559) 270-7698", Email: "dsaldivar@easycarewater.com"},
		{Name: "Jennifer Weaver", Title: "Order Processor", Phone: "(805) 901-3953", Email: "orderprocessing@easycarewater.com"},
	}
	return convertStruct2Byte(data)
}
func convertStruct2Byte(data []employee) []byte {
	d, err := json.Marshal(data)
	if err != nil {
		fmt.Println(err)
	}
	return d
}
