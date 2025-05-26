package main

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var quotes = make(map[int]Quote)
var nextID = 1

type Quote struct {
	ID     int    `json:"id"`
	Author string `json:"author"`
	Quote  string `json:"quote"`
}

func CreateQuote(w http.ResponseWriter, r *http.Request) {
	var q Quote
	err := json.NewDecoder(r.Body).Decode(&q)
	if err != nil || q.Author == "" || q.Quote == "" {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}
	q.ID = nextID
	nextID++
	quotes[q.ID] = q

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(q)
}

func GetQuotes(w http.ResponseWriter, r *http.Request) {
	var result []Quote
	author := r.URL.Query().Get("author")
	for _, q := range quotes {
		if author == "" || author == q.Author {
			result = append(result, q)
		}
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

func GetRandQuote(w http.ResponseWriter, r *http.Request) {
	if len(quotes) == 0 {
		http.Error(w, "No quotes available", http.StatusNotFound)
		return
	}
	keys := make([]int, 0, len(quotes))
	for k := range quotes {
		keys = append(keys, k)
	}
	randIdx := rand.Intn(len(keys))
	q := quotes[keys[randIdx]]
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(q)
}

func DeleteQuote(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.Atoi(idStr)

	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	if _, exists := quotes[id]; !exists {
		http.Error(w, "Quote not found", http.StatusNotFound)
		return
	}
	delete(quotes, id)

	w.WriteHeader(http.StatusNoContent)
}
