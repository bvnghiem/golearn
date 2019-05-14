package auth

import (
	"fmt"
	"net/http"
)

func BasicAuth(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
		username, password, ok := request.BasicAuth()
		if ok && username == "abc" && password == "123" {
			fmt.Println("Login success!")
			handler.ServeHTTP(response, request)
		} else {
			response.WriteHeader(http.StatusUnauthorized)
			response.Write([]byte("Unauthorized!"))
		}
	})
}
