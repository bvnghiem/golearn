package log2_middleware

import (
	"fmt"
	"net/http"
	"time"
)

func Log2Middleware(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
		today := time.Now()
		fmt.Println("Date: " + today.Format("02/01/2006 03:04:05"))
		handler.ServeHTTP(response, request)
	})
}
