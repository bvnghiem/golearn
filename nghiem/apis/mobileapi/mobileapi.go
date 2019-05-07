package mobileapi

import (
	"encoding/json"
	"fmt"
	"golearn/nghiem/apis/config"
	"golearn/nghiem/apis/model"
	"net/http"
)

func responseWithJSON(response http.ResponseWriter, statusCode int, data interface{}) {
	result, _ := json.Marshal(data)
	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(statusCode)
	response.Write(result)
}

func FindAll(response http.ResponseWriter, request *http.Request) {
	fmt.Println("Get all mobile")
	db, err := config.GetConnection()
	if err != nil {
		fmt.Println(err)
		return
	}

	mobileModel := model.MobileModel{
		Db:         db,
		Collection: "mobile",
	}
	mobiles, err2 := mobileModel.FindAll()

	if err2 != nil {
		fmt.Println(err2)
	}
	db.Session.Close()
	responseWithJSON(response, http.StatusOK, mobiles)
	fmt.Println("Done get all mobile")
}
