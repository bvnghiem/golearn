package log1_middleware

import (
	"fmt"
	"net/http"
)

func Log1Middleware(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
		fmt.Println("Log 1 handler function is handling url:", request.URL)
		handler.ServeHTTP(response, request)
	})
}
