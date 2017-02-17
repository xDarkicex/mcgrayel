package main

import (
	"errors"
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"
)

var goHTML *template.Template

// Gets arguments from a form value using a slice of strings as the key.
func getArguments(request *http.Request, keys []string) (arguments []string, err error) {
	err = errors.New("")
	for _, e := range keys {
		value := request.FormValue(e)
		if value == "" {
			err = fmt.Errorf("Required parameter: %s\n%s", e, err.Error())
		} else {
			arguments = append(arguments, value)
		}
	}
	if err.Error() == "" {
		err = nil
	}
	return arguments, err
}

type Product struct {
	Name string
}
type Options struct {
	isProduct bool
	View      string
}

func Parse(opts Options) *Options {
	return &Options{
		isProduct: opts.isProduct,
		View:      opts.View,
	}
}

func (*Product) setName(name string) *Product {
	return &Product{
		Name: name,
	}
}

type obj map[string]interface{}

func setObject(name string) *obj {
	switch name {

	case "pooltec":
		return &obj{
			"name": strings.ToTitle(name),
			"tag":  "Pooltec® 3-in-1 Pool Water Treatment",
			"points": []string{
				"Continuous prevention of green, yellow and black algae. Kills most algae in 4 to 24 hours.",
				"Strong clarifiers create ultra-clear water.",
				"Boosts chlorine’s effectiveness up to 6X (600%).",
				"Improves salt cell’s chlorine output and performance.",
				"Eliminates chlorine odors & skin-eye irritation.",
			},
		}
	case "algatec":
		return &obj{
			"name": strings.ToTitle(name),
			"tag":  "Fast Effective Algae Removal",
			"points": []string{
				"Kills and Controls all algae, green, yellow and black.",
				"Superior effectiveness and fast results in just 4 to 24 hours - Most Algae.",
				"Eradicates difficult Black Algae(fungi) growths 7 to 10 day.",
				"Strongest, fastest acting algaecide available.",
				"Metal-free, non-hazardous to handle, will not foam or stain.",
			},
		}
	default:
		return &obj{}
	}
}

func (p *Product) product(res http.ResponseWriter, req *http.Request) {
	raw := req.URL.Path
	t := strings.TrimPrefix(raw, "/Products/")
	name := strings.Split(t, "/")
	product := p.setName(strings.ToLower(name[2]))
	obj := setObject(product.Name)
	fmt.Println(product.Name)

	render(Parse(Options{isProduct: true}), res, *obj)
}

func render(Options *Options, res http.ResponseWriter, obj map[string]interface{}) {
	if Options.isProduct {
		goHTML = template.Must(template.ParseFiles("product.gohtml"))
		err := goHTML.Execute(res, obj)
		if err != nil {
			log.Fatal(err)
		}
	} else {
		goHTML = template.Must(template.ParseFiles(Options.View + ".gohtml"))
		err := goHTML.Execute(res, obj)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func index(res http.ResponseWriter, req *http.Request) {
	obj := map[string]interface{}{}
	render(Parse(Options{View: "index"}), res, obj)
}

func international(res http.ResponseWriter, req *http.Request) {
	obj := map[string]interface{}{}
	render(Parse(Options{View: "international"}), res, obj)

}

func main() {
	port := flag.String("p", "8000", "port to serve on")
	flag.Parse()
	fileServer := http.FileServer(http.Dir("."))
	http.HandleFunc("/", index)
	http.HandleFunc("/international", international)
	controller := Product{}
	http.HandleFunc("/products/pooltec", controller.product)
	http.HandleFunc("/products/algatec", controller.product)

	http.Handle("/static/", http.StripPrefix("/static/", fileServer))

	log.Fatal(http.ListenAndServe(":"+*port, nil))
}
