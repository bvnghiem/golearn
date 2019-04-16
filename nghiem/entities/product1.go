package entities

// mashalling/unmashalling
type Product1 struct {
	Id       string     `json:"id"`
	Name     string     `json:"name"`
	Price    float64    `json:"price"`
	Category Category1  `json:"category"`
	Comments []Comment1 `json:"comments"`
}

func NewProduct1(id, name string, price float64, category Category1, comments []Comment1) Product1 {
	return Product1{id, name, price, category, comments}
}
