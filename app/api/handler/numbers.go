package handler

import (
	"encoding/json"
	"github.com/mstovicek/showcase-go-api-number-aggregator/app/loader"
	"net/http"
)

type numbersJSON struct {
	Numbers []int `json:"numbers"`
}

type numbers struct {
	loader loader.Loader
}

func NewNumbers(loader loader.Loader) http.Handler {
	return &numbers{
		loader: loader,
	}
}

func (handler *numbers) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()

	var urls []string

	for key, url := range params {
		if key == "u" {
			urls = append(urls, url...)
			break
		}
	}

	numberCollection := handler.loader.Load(urls)
	dataJSON := numbersJSON{
		Numbers: numberCollection.GetUniqueSortedNumbers(),
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(dataJSON)
}
