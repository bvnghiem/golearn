package main

import (
	"fmt"
	"golearn/nghiem/apis/demo"
	"golearn/nghiem/apis/mobileapi"
	"golearn/nghiem/apis/productapi"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	//GET
	router.HandleFunc("/api/demo/demo1", demo.Demo1).Methods("GET")
	router.HandleFunc("/api/demo/demo2", demo.Demo2).Methods("GET")
	router.HandleFunc("/api/demo/demo3/{fullName}", demo.Demo3).Methods("GET")
	router.HandleFunc("/api/product/find", productapi.Find).Methods("GET")
	router.HandleFunc("/api/product/findall", productapi.FindAll).Methods("GET")

	//POST
	router.HandleFunc("/api/product/create", productapi.Create).Methods("POST")

	//POST
	router.HandleFunc("/api/product/update", productapi.Update).Methods("PUT")

	//POST
	router.HandleFunc("/api/product/delete/{id}", productapi.Delete).Methods("DELETE")

	//With MongoDB
	router.HandleFunc("/api/mobile/findall", mobileapi.FindAll).Methods("GET")
	router.HandleFunc("/api/mobile/find/{name}", mobileapi.Find).Methods("GET")

	err := http.ListenAndServe(":5000", router)

	if err != nil {
		fmt.Println(err)
	}
}
