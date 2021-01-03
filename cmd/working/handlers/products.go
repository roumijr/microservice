package handlers

import (
	"strconv"
	"net/http"
	"log"
	"context"


	"github.com/gorilla/mux"


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



func(p *Products) getProducts(rw http.ResponseWriter, r *http.Request) {	
	p.l.Println("Handle GET Products")

	//fetch the products from the database
	lp := data.GetProducts()

	//serialize the list to JSON
	err := lp.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to marshal data/JSON", http.StatusInternalServerError)
	}
}

func(p *Products) addProduct(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle POST Product")

	prod := r.Context().Value(KeyProduct{}).(data.Product)
	data.AddProduct(&prod)
}

func(p *Products) UpdateProducts(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		 http.Error(rw, "Unable to convert ID", http.StatusBadRequest)
		 return
	}
	p.l.Println("Handle PUT Product", id)
	prod := r.Context().Value(KeyProduct{}).(data.Product)
	
	

	err = data.UpdateProduct(id, &prod)
	if err == data.ErrProductNotFound {
		http.Error(rw, "Product not found", http.StatusNotFound)
		return
	}

	if err != nil {
		http.Error(rw, "Product not found", http.StatusInternalServerError)
		return
	}

}

type KeyProduct struct{}

func(p *Products) MiddlewareValidateProduct(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		prod := &data.Product{}
		
		err := prod.FromJSON(r.Body)
		if err != nil {
		http.Error(rw, "Unable to unmarshal JSON", http.StatusBadRequest)
		return
		}

		ctx := context.WithValue(r.Context(), KeyProduct{}, prod)
		req := r.WithContext(ctx)

		next.ServeHTTP(rw, req)
	})
}