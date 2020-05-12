package models

import (
	"golang-restful-api-angular/entities"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type ProductModels struct {
	Db *mgo.Database
}

func (productModel ProductModels) FindAll() ([]entities.Product, error) {
	var products []entities.Product

	err := productModel.Db.C("product").Find(bson.M{}).All(&products)
	if err != nil {
		return nil, err
	} else {
		return products, nil
	}
}

func (productModel ProductModels) FindByID(id string) (entities.Product, error) {
	var product entities.Product
	err := productModel.Db.C("product").FindId(bson.ObjectIdHex(id)).One(&product)
	return product, err
}

func (productModel ProductModels) Create(product *entities.Product) error {
	return productModel.Db.C("product").Insert(&product)
}

func (productModel ProductModels) Delete(id string) error {
	return productModel.Db.C("product").RemoveId(bson.ObjectIdHex(id))
}

func (productModel ProductModels) Update(product *entities.Product) error {
	return productModel.Db.C("product").Update(product.Id,product)
}
