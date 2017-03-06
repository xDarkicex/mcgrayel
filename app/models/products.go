package models

import (
	"time"

	"github.com/xDarkicex/mcgrayel/datastore"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

//Product models
//This is our current model of a product.
//In the future I may add cost and doage along with links.
type Product struct {
	ID            bson.ObjectId `bson:"_id,omitempty"`
	Image         string        `bson:"image"`
	PDFs          string        `bson:"pdfs"`
	Name          string        `bson:"name"`
	Tag           string        `bson:"tag"`
	Points        []string      `bson:"points"`
	Summary       string        `bson:"summary"`
	HowItWorks    string        `bson:"HowItWorks"`
	Features      []string      `bson:"Features"`
	SellingPoints []string      `bson:"sellingPoints"`
	Benefits      []string      `bson:"benefits"`
	Notice        string        `bson:"notice"`
	Time          time.Time     `bson:"time"`
}

//ProductCreate Creates new product in database
//returns error if data can not be entered
func ProductCreate(product *Product, file []byte) error {
	var err error
	session := datastore.GetSession()
	defer session.Close()
	collection := session.DB("Mcgrayel").C("Products")
	index := mgo.Index{
		Key:        []string{"name"},
		Unique:     true,
		DropDups:   true,
		Background: true,
		Sparse:     true,
	}
	if collection.EnsureIndex(index) != nil {
		return err
	}
	gridFS := session.DB("Mcgrayel").GridFS("fs")
	gridFile, err := gridFS.Create("")
	if err != nil {
		return err
	}
	defer gridFile.Close()
	_, err = gridFile.Write(file)
	if err != nil {
		return err
	}
	err = collection.Insert(&Product{
		Name:          product.Name,
		Image:         gridFile.Id().(bson.ObjectId).Hex(),
		Tag:           product.Tag,
		Points:        product.Points,
		Summary:       product.Summary,
		SellingPoints: product.SellingPoints,
		Features:      product.Features,
		Benefits:      product.Benefits,
		HowItWorks:    product.HowItWorks,
		Notice:        product.Notice,
		Time:          time.Now(),
	})
	if err != nil {
		return err
	}
	return err
}

func FindProductByIndex() {

}
