package main

import (
	"encoding/json"
	"net/http"
)

type handler struct {
	Numbers []int `json:"numbers"`
}

type wrongHandler struct {
	Numbers []int `json:"wrong"`
}

func newOdd() http.Handler {
	return &handler{
		Numbers: []int{1, 3, 5, 7, 9, 11, 13, 15, 17, 19},
	}
}

func newEven() http.Handler {
	return &handler{
		Numbers: []int{2, 4, 6, 8, 10, 12, 14, 16, 18, 20},
	}
}

func newTen() http.Handler {
	return &handler{
		Numbers: []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1},
	}
}

func newFibo() http.Handler {
	return &handler{
		Numbers: []int{13, 8, 5, 3, 2, 1},
	}
}

func newPrimes() http.Handler {
	return &handler{
		Numbers: []int{19, 17, 13, 11, 7, 5, 3, 2},
	}
}

func newWrong() http.Handler {
	return &wrongHandler{
		Numbers: []int{1, 2, 3, 4, 5},
	}
}

func (handler *handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(handler)
}

func (handler *wrongHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(handler)
}
