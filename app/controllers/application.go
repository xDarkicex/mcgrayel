package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"

	"github.com/xDarkicex/mcgrayel/helpers"
)

//Application this is here to pass a type to router
type Application helpers.Controller

func (this Application) Index(a helpers.RouterArgs) {
	if a.Request.Method == "HEAD" {
		a.Response.WriteHeader(200)
		return
	}
	helpers.Render(a, "application/index", map[string]interface{}{})
}

func (this Application) VideoTour(a helpers.RouterArgs) {
	if a.Request.Method == "HEAD" {
		a.Response.WriteHeader(200)
		return
	}
	helpers.Render(a, "application/videotour", map[string]interface{}{})
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

//Load loads config file
func Load() {

}

func ReadJSON() *helpers.Secret {
	data := helpers.Secret{}
	secret, err := ioutil.ReadFile("./secret.json")
	if err != nil {
		fmt.Println(err, secret)
	}
	// fmt.Println(config)
	err = json.Unmarshal(secret, &data)
	if err != nil {
		fmt.Println(err)
	}
	return &data
}

//Note to self add a router for emails
// example would be code like this
// whois := a.Request.FormValue("whois")
// if whois == "orderprocessing" {
// email jen
//}
// if whois == "accounting" {
// email monica
//}
// a switch would be more efficent code

func (this Application) Contact(a helpers.RouterArgs) {
	if a.Request.Method == "POST" {
		msg := NewMessage(a)
		department, err := strconv.Atoi(msg.Department)
		if err != nil {
			log.Println(err)
		}

		// password := ReadJSON()
		switch d := department; d {
		case 1:
			fmt.Println("Customer Support")
			// subject := "Contact Request from " + msg.Name
			// data := "New Contact request from " + msg.Name + "\n" +
			// 	"Email: " + msg.Email + "\n" +
			// 	"Phone Number: " + msg.Telephone + "\n" +
			// 	"Product: " + msg.Product + "\n" +
			// 	"Contact Message: " + "\n" +
			// 	msg.Body
			// m := email.NewMessage(subject, data)
			// m.From = mail.Address{Name: "Services", Address: "orderprocessing@easycarewater.com"}
			// m.To = []string{"orderprocessing@easycarewater.com"}
			// auth := smtp.PlainAuth("", "orderprocessing@easycarewater.com", password.Password, "smtp.gmail.com")
			// SMTP := "smtp.gmail.com" + ":" + strconv.Itoa(587)
			// if err := email.Send(SMTP, auth, m); err != nil {
			// 	fmt.Println(err)
			// }
		case 2:
			fmt.Println("Order Processing")
			// subject := "Contact Request from " + msg.Name
			// data := "New Contact request from " + msg.Name + "\n" +
			// 	"Email: " + msg.Email + "\n" +
			// 	"Phone Number: " + msg.Telephone + "\n" +
			// 	"Product: " + msg.Product + "\n" +
			// 	"Contact Message: " + "\n" +
			// 	msg.Body
			// m := email.NewMessage(subject, data)
			// m.From = mail.Address{Name: "Services", Address: "orderprocessing@easycarewater.com"}
			// m.To = []string{"orderprocessing@easycarewater.com"}
			// auth := smtp.PlainAuth("", "orderprocessing@easycarewater.com", password.Password, "smtp.gmail.com")
			// SMTP := "smtp.gmail.com" + ":" + strconv.Itoa(587)
			// if err := email.Send(SMTP, auth, m); err != nil {
			// 	fmt.Println(err)
			// }
		case 3:
			fmt.Println("Water treatment")
		case 4:
			fmt.Println("Web Master")
		}
	}
	data := map[string]interface{}{}
	helpers.Render(a, "application/contact", data)
}

func (this Application) SDS(a helpers.RouterArgs) {
	data := map[string]interface{}{}
	helpers.Render(a, "application/sds", data)
}

type Locate helpers.Controller

func (this Locate) Locate(a helpers.RouterArgs) {
	helpers.Render(a, "storeLocator/locator", nil)
}

func loadSDSSheet(name string) *data {
	switch name {
	case "algatec":
		return &data{
			"revision":                "06/26/2017",
			"productCode":             "#10064",
			"productName":             "Algatec® - Green, Yellow and Black Algaecide",
			"recommendedUse":          "Swimming pool water treatment",
			"restrictionsOnUseOrSale": "none",
			"companyName":             "EasyCare Products USA, div. of McGrayel Company Inc.",
			"address":                 "5361 S. Villa Avenue Fresno, CA 93725",
			"emergencyContacts": []string{
				"Debbie Saldivar",
				"Marvin J. Rezac",
			},
			"contactEmails": []string{
				"dsaldivar@easycarewater.com",
				"mrezac@easycarewater.com",
			},
			"hazards": []string{
				"Oral, Category 4 &#10; Aquatic Toxicity (Acute), Category 1",
				"H302- Harful is swallowed. &#10; H400 - Very toxic to aquatic life.",
				"P264 - Wash hands thoroughly after handling. &#10; P270 - Do not drink, eat or smoke when using this product. &#10; P273 - Avoid release to the environment.",
				"P301 + 312 - IF SWALLOWED: Call a POISON CENTER or doctor if you fell unwell. &#10; P330 - Rinse mouth. &#10; P391 - Collect spillage.",
				"P501 - Disose of contents / container to regulatory requirements.",
				"This material is classi ed as hazardous under OSHA regulations.",
				"Prolonged inhalation may be harmful. Inhalation may be harmful.",
				"Prolonged or repeated skin contact may cause irritation.",
				"Contact may cause eye irritation.",
				"Harmful is swallowed. If medical advise is needed, have product container or label at hand.",
			},
			"HazardousRatingSystem": []string{
				"1",
				"0",
				"0",
				"B",
			},
			"composition": []string{
				"31512-74-0",
				"Poly[oxy-1,2-ethanediyl(dimethyliminio)-1,2-ethan ediyl(dimethyliminio)-1,2-ethanediyl dichloride]",
				"30.0 %",
				"TR1650000",
				"7732-18-5",
				"Water",
				"70.0%",
				"ZC0110000",
			},
			"firstAid": []string{
				"IF INHALED: If breathing is di cult, remove to fresh air and keep at rest in a position comfortable for breathing. If breathing is di cult, give oxygen. IF NOT BREATHING, call 911 or and ambulance, then give arti cial respiration.",
				"Wash with soap and water. Take o  contaminated clothing and wash before re-use. If skin irritation or rash occurs, seek medical advice / attention.",
				"Hold eyelids apart and  ush eyes with plent of water. After initial  ushings, remove any contact lenses and continue  ushing for at least 15 minutes. Have eyes examined and tested by medical personnel.",
				"If swallowed, do NOT induce vomiting. Give victim a glass of water or milk. Call a physician or poison control center immediately. Never give anything by mouth to an unconscious person.",
				"Treat symptomatically and supportively.",
			},
			"fire": []string{
				"> 210 F (100.0 C) Method Used: Cleveland Open Cup.",
				"LEL: No Data UEL: No Data",
				"N/A",
				"Use extinguishing agent suitable for type of surrounding fire.",
				"As in any  re, wear a self-contained breathing apparatus in pressure-demand, MSHA / NIOSH (approved or equivalent), and full protective gear. Material will not burn.",
				"No Data available",
			},
			"accidental": []string{
				"Avoid release to the environment. This product is toxic to  sh and aquatic organisms. Do not discharge into e uent containing this product into lakes, streams, ponds or estaries, oceans or other water unless in accordance with the requirements of a National Pollutant Discharge Elimination System (NPDES) permit and the permitting authority has been noti ed in writing prior to discharging. Do no discharge e uent containing this product to sewer systems without previously notifying the local sewage treatment plant authority. For guidance call your State Water Board Authority or Regional O ce of the EPA.",
				"Use proper personal protective equipment as indicated in Section 8. Spills / Leaks: Absorb spill with inert material (e.g. vermiculite, sand or earth), then place in suitable container.",
			},
			"handling": []string{
				"No special handling procedures are required. Do not ingest. Avoid contact with eyes, skin or clothing. Avoid release to the environment. Keep container closed when not in use.",
				"No special storage requirements. Storage temperature - ambient (avoid freezing). Store in accordance with local requlations. Do not reuse container, rinse bottle and o er for recycling.",
			},
			"exposureControls": []string{
				"31512-74-0",
				"Poly[oxy-1,2-ethanediyl(dimethyliminio)-1,2-ethanediyl(dimethyliminio)1, 2-ethanediyl dichloride]",
				"No Data",
				"No Data",
				"No Data",
				"7732-18-5",
				"water",
				"No Data",
				"No Data",
				"No Data",
			},
		}
	}
	return &data{}
}

//NewMessage returns a ContactForm struct
// Name
// Email
// Telephone
// Product
// Body
func NewMessage(a helpers.RouterArgs) *helpers.ContactForm {
	return &helpers.ContactForm{
		Department: a.Request.FormValue("department"),
		Name:       a.Request.FormValue("name"),
		Email:      a.Request.FormValue("email"),
		Telephone:  a.Request.FormValue("telephone"),
		Product:    a.Request.FormValue("product"),
		Body:       a.Request.FormValue("body"),
	}
}
