package models

import (
	"golearn/nghiem/mongodb/entities"

	"gopkg.in/mgo.v2/bson"

	"gopkg.in/mgo.v2"
)

type ProductModel struct {
	Db         *mgo.Database
	Collection string
}

func (p ProductModel) FindAll() (products []entities.Product, err error) {
	err = p.Db.C(p.Collection).Find(bson.M{}).All(&products)
	return
}

func (p ProductModel) Search1(min float64) (products []entities.Product, err error) {
	err = p.Db.C(p.Collection).Find(bson.M{
		"price": bson.M{"$gte": min},
	}).All(&products)
	return
}

func (p ProductModel) Search2(min, max float64) (products []entities.Product, err error) {
	err = p.Db.C(p.Collection).Find(bson.M{
		"$and": []bson.M{
			bson.M{"price": bson.M{"$gte": min}},
			bson.M{"price": bson.M{"$lte": max}},
		},
	}).All(&products)
	return
}

func (p ProductModel) Like1(key string) (products []entities.Product, err error) {
	err = p.Db.C(p.Collection).Find(bson.M{
		"name": bson.RegEx{Pattern: "^" + key, Options: "i"}, // like %key non-case sensitvie
	}).All(&products)
	return
}

func (p ProductModel) Like2(key string) (products []entities.Product, err error) {
	err = p.Db.C(p.Collection).Find(bson.M{
		"name": bson.RegEx{Pattern: key + "$", Options: "i"}, // like key% non-case sensitvie
	}).All(&products)
	return
}

func (p ProductModel) Like3(key string) (products []entities.Product, err error) {
	err = p.Db.C(p.Collection).Find(bson.M{
		"name": bson.RegEx{Pattern: key, Options: "i"}, // like %key% non-case sensitvie
	}).All(&products)
	return
}

func (p ProductModel) SortAsc(key string) (products []entities.Product, err error) {
	err = p.Db.C(p.Collection).Find(bson.M{}).Sort("price").All(&products)
	return
}

func (p ProductModel) SortDesc(key string) (products []entities.Product, err error) {
	err = p.Db.C(p.Collection).Find(bson.M{}).Sort("-price").All(&products)
	return
}

func (p ProductModel) Paging1(index, size int) (products []entities.Product, err error) {
	err = p.Db.C(p.Collection).Find(bson.M{}).Sort("-price").Limit(size).Skip(index).All(&products)
	return
}

func (p ProductModel) FindById(id string) (product entities.Product, err error) {
	err = p.Db.C(p.Collection).FindId(bson.ObjectIdHex(id)).One(&product)
	return
}

func (p ProductModel) Create(product *entities.Product) error {
	err := p.Db.C(p.Collection).Insert(&product)
	return err
}

func (p ProductModel) Update(product *entities.Product, id string) error {
	err := p.Db.C(p.Collection).Update(product.Id, &product)
	return err
}

func (p ProductModel) Delete(id string) error {
	err := p.Db.C(p.Collection).RemoveId(bson.ObjectIdHex(id))
	return err
}
