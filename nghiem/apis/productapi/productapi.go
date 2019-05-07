package productapi

import (
	"encoding/json"
	"fmt"
	"golearn/nghiem/apis/entities"
	"net/http"

	"github.com/gorilla/mux"
)

func responseWithJSON(response http.ResponseWriter, statusCode int, data interface{}) {
	result, _ := json.Marshal(data)
	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(statusCode)
	response.Write(result)
}

func Find(response http.ResponseWriter, request *http.Request) {
	product := entities.Product{
		Id:       "p01",
		Name:     "Product 01",
		Price:    345.23,
		Quantity: 24,
	}
	responseWithJSON(response, http.StatusOK, product)
}

func FindAll(response http.ResponseWriter, request *http.Request) {
	products := []entities.Product{
		entities.Product{
			Id:       "p01",
			Name:     "Product 01",
			Price:    345.23,
			Quantity: 24,
		}, entities.Product{
			Id:       "p02",
			Name:     "Product 02",
			Price:    12.3,
			Quantity: 45,
		},
		entities.Product{
			Id:       "p03",
			Name:     "Product 03",
			Price:    675,
			Quantity: 65,
		},
	}
	responseWithJSON(response, http.StatusOK, products)
}

func Create(response http.ResponseWriter, request *http.Request) {
	var product entities.Product
	json.NewDecoder(request.Body).Decode(&product)

	fmt.Printf("Create new product: \n id: %s\n name: %s\n price: %.3f\n quantity: %d\n",
		product.Id, product.Name, product.Price, product.Quantity)
}

func Update(response http.ResponseWriter, request *http.Request) {
	var product entities.Product
	json.NewDecoder(request.Body).Decode(&product)

	fmt.Printf("Update product: \n id: %s\n name: %s\n price: %.3f\n quantity: %d\n",
		product.Id, product.Name, product.Price, product.Quantity)
}

func Delete(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	id := vars["id"]

	fmt.Printf("Delete product: \n id: %s\n", id)
}
