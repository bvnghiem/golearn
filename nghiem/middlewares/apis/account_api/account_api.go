package account_api

import (
	"encoding/json"
	"golearn/nghiem/middlewares/config"
	"golearn/nghiem/middlewares/entities"
	"golearn/nghiem/middlewares/models"
	"net/http"

	"gopkg.in/mgo.v2/bson"
)

func Create(response http.ResponseWriter, request *http.Request) {
	var account entities.Account
	err := json.NewDecoder(request.Body).Decode(&account)
	if err != nil {
		responseWithError(response, http.StatusBadRequest, err.Error())
	} else {
		db, err2 := config.GetConnection()

		if err2 != nil {
			responseWithError(response, http.StatusBadGateway, err.Error())
		} else {
			accModel := models.AccountModel{
				Db:         db,
				Collection: "account",
			}
			account.Id = bson.NewObjectId()
			err3 := accModel.Create(&account)
			if err3 != nil {
				responseWithError(response, http.StatusBadGateway, err.Error())
			} else {
				responseWithJSON(response, http.StatusOK, account)
			}
		}
	}
}

func responseWithError(response http.ResponseWriter, statusCode int, msg string) {
	responseWithJSON(response, statusCode, map[string]string{"error": msg})
}
func responseWithJSON(response http.ResponseWriter, statusCode int, data interface{}) {
	result, _ := json.Marshal(data)
	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(statusCode)
	response.Write(result)
}
