package main

import (
	"fmt"
	"golearn/nghiem/middlewares/apis/account_api"
	"golearn/nghiem/middlewares/apis/demo_api"
	"golearn/nghiem/middlewares/apis/jwt_api"
	"golearn/nghiem/middlewares/middlewares/auth"
	"golearn/nghiem/middlewares/middlewares/jwt_middleware"
	"golearn/nghiem/middlewares/middlewares/log1_middleware"
	"golearn/nghiem/middlewares/middlewares/log2_middleware"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/api/demo/demo", demo_api.Demo)
	router.Handle("/api/demo/demo1",
		log1_middleware.Log1Middleware(http.HandlerFunc(demo_api.Demo1))).Methods("GET")

	router.Handle("/api/demo/demo2",
		log2_middleware.Log2Middleware(http.HandlerFunc(demo_api.Demo2))).Methods("GET")

	router.Handle("/api/demo/demo3",
		log1_middleware.Log1Middleware(log2_middleware.Log2Middleware(
			http.HandlerFunc(demo_api.Demo3)))).Methods("GET")

	router.Handle("/api/demo/demo4",
		auth.BasicAuth(http.HandlerFunc(demo_api.Demo4))).Methods("GET")

	router.HandleFunc("/api/account/create", account_api.Create).Methods("POST")

	router.HandleFunc("/api/jwt/genkey", jwt_api.GenerateToken).Methods("POST")

	router.Handle("/api/demo/demo5",
		jwt_middleware.CheckToken(http.HandlerFunc(demo_api.Demo5))).Methods("POST")
	err := http.ListenAndServe(":5000", router)
	if err != nil {
		fmt.Println(err)
	}
}
