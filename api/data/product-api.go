package data

import (
	"encoding/json"
	"io"
	"time"
)
 

// Product defines a structure for an api product
type Product struct {
	ID			int		`json:"id"`
	Name		string	`json:"name"`
	Description string	`json:"description"`
	Price 		float32	`json:"price"`
	SKU			string	`json:"sku"`
	CreatedOn	string	`json:"-"`
	UpdatedOn	string	`json:"-"`
	DeletedOn	string	`json:"-"`
}

// Products is a receiver
type Products []*Product 

// ToJSON method converts data to json format
func(p *Products) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}

// GetProducts method returns all data from productList
func GetProducts() Products {
	return productList
}

var productList = []*Product {
	&Product{
	ID:			 1,			
	Name:		 "latte",
	Description: "Milk coffe",
	Price: 		 2.45,
	SKU:		 "abcd1",	
	CreatedOn:	 time.Now().UTC().String(),
	UpdatedOn:	 time.Now().UTC().String(),
	DeletedOn:	 time.Now().UTC().String(),

	},
	&Product{
	ID:			 2,
	Name:		 "Espresso",
	Description: "Strong coffe without milk",
	Price: 		 1.99,
	SKU:		 "abcd2",
	CreatedOn:	 time.Now().UTC().String(),
	UpdatedOn:	 time.Now().UTC().String(),
	DeletedOn:	 time.Now().UTC().String(),

	},
} 



