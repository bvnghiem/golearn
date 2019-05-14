package jwt_middleware

import (
	"encoding/json"
	"golearn/nghiem/middlewares/security"
	"net/http"

	"github.com/dgrijalva/jwt-go"
)

func CheckToken(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
		tokenStr := request.Header.Get("Token")
		if tokenStr == "" {
			responseWithError(response, http.StatusUnauthorized, "NO token")
		} else {
			result, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
				return []byte(security.PrivateKey), nil
			})
			if err != nil {
				responseWithError(response, http.StatusInternalServerError, err.Error())
			} else {
				if result.Valid {
					handler.ServeHTTP(response, request)
				} else {
					responseWithError(response, http.StatusBadRequest, "Invalid token")
				}
			}
		}
	})
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
