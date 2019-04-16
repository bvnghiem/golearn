package mashal

import (
	"encoding/json"
	"fmt"
	"golearn/nghiem/entities"
)

func Demo1() {
	p := entities.NewProduct1("p01", "Product 1", 35.3, entities.Category1{
		Id:   "1",
		Name: "Category 1",
	}, []entities.Comment1{
		entities.Comment1{Id: 1, Title: "Title 1", Content: "Content 1"},
		entities.Comment1{Id: 2, Title: "Title 2", Content: "Content 2"},
	})
	result, error := json.Marshal(p)

	if error != nil {
		fmt.Println(error)
	} else {
		fmt.Println(string(result))
	}
}

func Demo2() {
	s := `{"id":"p01","name":"Product 1","price":35.3,"category":{"id":"1","name":"Category 1"},"comments":[{"id":1,"title":"Title 1","content":"Content 1"},{"id":2,"title":"Title 2","content":"Content 2"}]}`
	var p entities.Product1
	json.Unmarshal([]byte(s), &p)
	fmt.Println("Product info")
	fmt.Println("Id: ", p.Id)
	fmt.Println("Name: ", p.Name)
	fmt.Println("Price: ", p.Price)
	fmt.Println("Category ID: ", p.Category.Id)
	fmt.Println("Category Name: ", p.Category.Name)
	fmt.Println("Comments:")
	fmt.Println("--------")
	for _, c := range p.Comments {
		fmt.Println("Comment Id: ", c.Id)
		fmt.Println("Comment title: ", c.Title)
		fmt.Println("Comment content: ", c.Content)
	}
}
