package data

import (
	"fmt"
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

//FromJSON convert the data
func(p *Product) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(p)
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

// AddProduct is adding product to the database
func AddProduct(p *Product) {
	p.ID = GetNextID()
	productList = append(productList, p)
}

// UpdateProduct is updating the list of products
func UpdateProduct(id int, p *Product) error {
	_, pos, err := findProduct(id)
	if err != nil {
		return err
	}

	p.ID = id
	productList[pos] = p

	return nil
}

// ErrProductNotFound is custom error for the productList
var ErrProductNotFound = fmt.Errorf("Product not found")

func findProduct(id int) (*Product, int, error) {
	for i, p := range productList {
		if p.ID == id {
			return p, i, nil
		}
	}

	return nil, -1, ErrProductNotFound
}

//GetNextID serching for the next product ID
func GetNextID() int {
	lp := productList[len(productList) - 1]
	return lp.ID + 1
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



