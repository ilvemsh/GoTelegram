package main

import (
	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/quotes", CreateQuote).Methods("POST")
	router.HandleFunc("/quotes", GetQuotes).Methods("GET")
	router.HandleFunc("/quotes/random", GetRandQuote).Methods("GET")
	router.HandleFunc("/quotes/{id}", DeleteQuote).Methods("DELETE")

	return router
}
