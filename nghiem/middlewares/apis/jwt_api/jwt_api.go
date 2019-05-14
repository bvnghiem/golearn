package jwt_api

import (
	"encoding/json"
	"golearn/nghiem/middlewares/config"
	"golearn/nghiem/middlewares/entities"
	"golearn/nghiem/middlewares/models"
	"golearn/nghiem/middlewares/security"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func GenerateToken(response http.ResponseWriter, request *http.Request) {
	var account entities.Account
	err := json.NewDecoder(request.Body).Decode(&account)
	if err != nil {
		responseWithError(response, http.StatusBadRequest, err.Error())
	} else {

		db, err1 := config.GetConnection()

		if err1 != nil {
			responseWithError(response, http.StatusBadGateway, err1.Error())
		} else {
			accModel := models.AccountModel{
				Db:         db,
				Collection: "account",
			}
			isValid := accModel.Check(account.Username, account.Password)
			if isValid {
				token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
					"username": account.Username,
					"password": account.Password,
					"exp":      time.Now().Add(time.Hour * 72).Unix(),
				})
				tokenString, err2 := token.SignedString([]byte(security.PrivateKey))
				if err2 != nil {
					responseWithError(response, http.StatusBadRequest, err2.Error())
				} else {
					responseWithJSON(response, http.StatusOK, entities.JwtToken{
						Token: tokenString,
					})
				}
			} else {
				responseWithError(response, http.StatusBadRequest, "Invalid account")
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
