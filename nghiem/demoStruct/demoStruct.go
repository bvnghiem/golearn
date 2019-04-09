package demoStruct

import (
	"fmt"
	"nghiem/entities"
	"sort"
	"strings"
)

func Demo1() {
	var p entities.Product
	p.Id = "p1"
	p.Name = "Name 1"
	p.Price = 10.2
	p.Quantity = 3
	p.Status = true

	fmt.Println("Product info")
	fmt.Println("id: ", p.Id)
	fmt.Println("name: ", p.Name)
	fmt.Println("price: ", p.Price)
	fmt.Println("quantity: ", p.Quantity)
	fmt.Println("status: ", p.Status)
	fmt.Println("total: ", int(p.Price)*p.Quantity)
	display(p)
}

func Demo2() {
	p := entities.Product{
		Id:       "p2",
		Name:     "Name 2",
		Price:    5.6,
		Quantity: 4,
		Status:   false,
	}

	fmt.Println("Product info")
	fmt.Println("id: ", p.Id)
	fmt.Println("name: ", p.Name)
	fmt.Println("price: ", p.Price)
	fmt.Println("quantity: ", p.Quantity)
	fmt.Println("status: ", p.Status)
	fmt.Println("total: ", int(p.Price)*p.Quantity)

	display(p)
}

func display(p entities.Product) {
	fmt.Println("id: ", p.Id,
		"name:", p.Name, "price:", p.Price,
		"quantity:", p.Quantity, "status:", p.Status,
		"total:", p.Price*float64(p.Quantity),
		"category name:", p.Category.Name,
		"comments:", p.Comments[0].Content)
}

func Find() entities.Product {

	p := entities.Product{
		Id:       "p2",
		Name:     "Name 2",
		Price:    5.6,
		Quantity: 4,
		Status:   false,
	}
	return p
}

func changeValue1(p entities.Product) {
	p.Price = 6
}

func changeValue2(p *entities.Product) {
	p.Price = 6
	(*p).Quantity = 8
}

func Demo3() {
	p := Find()
	display(p)
	changeValue1(p)
	display(p)
	changeValue2(&p)
	display(p)
}

func Demo4() {
	products := []entities.Product{}
	products = append(products, entities.Product{
		Id:       "p1",
		Name:     "Name 1",
		Price:    5.6,
		Quantity: 4,
		Status:   false,
	})

	products = append(products, entities.Product{
		Id:       "p2",
		Name:     "Name 2",
		Price:    2.6,
		Quantity: 5,
		Status:   false,
	})
	products = append(products, entities.Product{
		Id:       "p3",
		Name:     "Name 3",
		Price:    6.1,
		Quantity: 6,
		Status:   true,
	})

	fmt.Println("Product list")
	for i := 0; i < len(products); i++ {
		display(products[i])
	}

	fmt.Println("=================")
	fmt.Println("Product list")
	for _, p := range products {
		display(p)
	}
	displayProducts(products)
}

func displayProducts(products []entities.Product) {
	fmt.Println("=================")
	for _, p := range products {
		display(p)
	}
}

func tongTien(products []entities.Product) float64 {
	tongTien := 0.0
	for _, p := range products {
		tongTien += p.Price * float64(p.Quantity)
	}
	return tongTien
}

func giaMaxMin(products []entities.Product, minPrice, maxPrice float64) (max, min entities.Product) {
	for i := 0; i < len(products); i++ {
		if maxPrice < products[i].Price {
			maxPrice = products[i].Price
			max = products[i]
		}
		if minPrice > products[i].Price {
			minPrice = products[i].Price
			min = products[i]
		}
	}
	return max, min
}

func timTheoTen(products []entities.Product, key string) (results []entities.Product) {
	for _, p := range products {
		if strings.Contains(strings.ToLower(p.Name), strings.ToLower(key)) {
			results = append(results, p)
		}
	}
	return results
}

func sortProducts(products []entities.Product) {
	sort.Slice(products, func(i, j int) bool {
		return products[i].Price > products[j].Price
	})
}

func countMinMax(products []entities.Product, minPrice, maxPrice float64) int {
	result := 0
	for i := 0; i < len(products); i++ {
		if maxPrice >= products[i].Price && minPrice <= products[i].Price {
			result += products[i].Quantity
		}
	}
	return result
}

func Demo5() {

	products := []entities.Product{}
	products = append(products, entities.Product{
		Id:       "p1",
		Name:     "Name 1",
		Price:    5.6,
		Quantity: 4,
		Status:   false,
	})

	products = append(products, entities.Product{
		Id:       "p2",
		Name:     "Name 2",
		Price:    2.6,
		Quantity: 5,
		Status:   false,
	})
	products = append(products, entities.Product{
		Id:       "p3",
		Name:     "Name 3",
		Price:    6.1,
		Quantity: 6,
		Status:   true,
	})
	products = append(products, entities.Product{
		Id:       "p4",
		Name:     "Name 4",
		Price:    15,
		Quantity: 8,
		Status:   true,
	})

	displayProducts(products)

	fmt.Println("Subtotal: ", tongTien(products))
	pMax, pMin := giaMaxMin(products, 3, 8)
	fmt.Println("---------------------")
	fmt.Println("Max price product:")
	display(pMax)
	fmt.Println("Min price product:")
	display(pMin)
	fmt.Println("---------------------")
	fmt.Println("Product has 2 in name:")
	displayProducts(timTheoTen(products, "2"))
	fmt.Println("---------------------")
	fmt.Println("Sorted")
	sortProducts(products)
	displayProducts(products)
	fmt.Println("---------------------")
	fmt.Println("Number of product has 3 < price < 20:", countMinMax(products, 3, 20))
}

func Demo6() {
	product := entities.Product{
		Id:       "p4",
		Name:     "Name 4",
		Price:    15,
		Quantity: 8,
		Status:   true,
		Category: entities.Category{
			Id:   "c1",
			Name: "Category 1",
		},
		Comments: []entities.Comment{
			entities.Comment{
				Id:      1,
				Title:   "Title 1",
				Content: "Content 1",
			},
		},
	}
	display(product)
}
