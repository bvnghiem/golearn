package entities

type Mobile struct {
	Id       string  `bson:"_id" json:"id"`
	Name     string  `bson:"name" json:"name"`
	Price    float64 `bson:"price" json:"price"`
	Quantity int     `bson:"quantity" json:"quantity"`
	Status   bool    `bson:"status" json:"status"`
}
