package demo_api

import (
	"fmt"
	"net/http"
)

func Demo(response http.ResponseWriter, request *http.Request) {
	fmt.Fprint(response, "Hello!")
}

func Demo1(response http.ResponseWriter, request *http.Request) {
	fmt.Fprint(response, "Hello demo 1!")
}

func Demo2(response http.ResponseWriter, request *http.Request) {
	fmt.Fprint(response, "Hello demo 2!")
}

func Demo3(response http.ResponseWriter, request *http.Request) {
	fmt.Fprint(response, "Hello demo 3!")
}

func Demo4(response http.ResponseWriter, request *http.Request) {
	fmt.Fprint(response, "Hello demo 4!")
}

func Demo5(response http.ResponseWriter, request *http.Request) {
	fmt.Fprint(response, "Hello demo 5!")
}
