package main

import (
	"fmt"
	"golearn/nghiem/mongodb/config"
	"golearn/nghiem/mongodb/entities"
	"golearn/nghiem/mongodb/models"

	"gopkg.in/mgo.v2/bson"
)

func Demo1() {
	db, err := config.GetConnection()
	if err != nil {
		fmt.Println(err)
		return
	}

	productModel := models.ProductModel{
		Db:         db,
		Collection: "product",
	}

	fmt.Println(db)
	fmt.Println("Find all: ")

	fmt.Println(*productModel.Db, productModel.Collection)

	prods, err2 := productModel.FindAll()

	if err2 != nil {
		fmt.Println(err2)
	}

	for _, prod := range prods {
		fmt.Println(prod.ToString())
	}

	fmt.Println("Search price >= 3")
	prods, err2 = productModel.Search1(3)
	if err2 != nil {
		fmt.Println(err2)
	}

	for _, prod := range prods {
		fmt.Println(prod.ToString())
	}

	product, err := productModel.FindById("")

	product = entities.Product{
		Id:       bson.NewObjectId(),
		Name:     "laptop 1",
		Price:    423.2,
		Quantity: 1,
		Status:   true,
	}

	productModel.Create(&product)
}

func main() {
	Demo1()
}
