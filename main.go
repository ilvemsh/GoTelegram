package main

import (
	"fmt"
	"net/http"
)

var port = "8080"

func main() {
	router := NewRouter()
	err := http.ListenAndServe(":"+port, router)
	if err != nil {
		fmt.Println("Error!!", err)
	} else {
		fmt.Println("starting server at :", port)
	}
}
