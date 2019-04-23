package entities

import (
	"fmt"

	"gopkg.in/mgo.v2/bson"
)

type Product struct {
	Id       bson.ObjectId `bson:"_id"`
	Name     string        `bson:"name"`
	Price    float64       `bson:"price"`
	Quantity int           `bson:"quantity"`
	Status   bool          `bson:"status"`
}

func (p Product) ToString() string {
	return fmt.Sprintf("id : %s, name: %s, price: %.1f, quantity: %d, status: %s",
		p.Id.Hex(), p.Name, p.Price, p.Quantity, p.Status)
}
