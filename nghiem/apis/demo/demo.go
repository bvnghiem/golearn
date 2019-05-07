package demo

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func Demo1(response http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(response, "Hello world!")
}

func Demo2(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "text/html")
	fmt.Fprintf(response, "<b><i><u>Hi!</u></i></b>")
}

func Demo3(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	fullName := vars["fullName"]

	fmt.Fprintf(response, "Hello "+fullName)
}
