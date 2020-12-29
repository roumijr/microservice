package handlers

import (
	"net/http"
	"log"


	"working/api/data"
)

// Products struct that listening to new data 
type Products struct {
	l *log.Logger
}

// NewProducts is creating a new product
func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

func(p *Products) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		p.getProducts(rw, r)
		return
	}

	// Catch all
	rw.WriteHeader(http.StatusMethodNotAllowed)
}

func(p *Products) getProducts(rw http.ResponseWriter, r *http.Request) {
	lp := data.GetProducts()

	err := lp.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to marshal data/json", http.StatusInternalServerError)
	}
}