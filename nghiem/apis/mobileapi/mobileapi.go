package mobileapi

import (
	"encoding/json"
	"fmt"
	"golearn/nghiem/apis/config"
	"golearn/nghiem/apis/model"
	"net/http"

	"github.com/gorilla/mux"
)

func responseWithJSON(response http.ResponseWriter, statusCode int, data interface{}) {
	result, _ := json.Marshal(data)
	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(statusCode)
	response.Write(result)
}

func responseWithError(response http.ResponseWriter, statusCode int, msg string) {
	responseWithJSON(response, statusCode, map[string]string{"error": msg})
}

func FindAll(response http.ResponseWriter, request *http.Request) {
	fmt.Println("Get all mobile")
	db, err := config.GetConnection()
	if err != nil {
		responseWithError(response, http.StatusBadGateway, err.Error())
	}

	mobileModel := model.MobileModel{
		Db:         db,
		Collection: "mobile",
	}
	mobiles, err2 := mobileModel.FindAll()

	if err2 != nil {
		responseWithError(response, http.StatusBadRequest, err2.Error())
	}
	db.Session.Close()
	responseWithJSON(response, http.StatusOK, mobiles)
	fmt.Println("Done get all mobile")
}

func Find(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	name := vars["name"]
	fmt.Println("Find mobile by name")
	db, err := config.GetConnection()
	if err != nil {
		responseWithError(response, http.StatusBadGateway, err.Error())
	}

	mobileModel := model.MobileModel{
		Db:         db,
		Collection: "mobile",
	}
	mobiles, err2 := mobileModel.Find(name)

	if err2 != nil {
		responseWithError(response, http.StatusBadRequest, err2.Error())
	}
	db.Session.Close()
	responseWithJSON(response, http.StatusOK, mobiles)
	fmt.Println("Done find  mobile by name")
}
