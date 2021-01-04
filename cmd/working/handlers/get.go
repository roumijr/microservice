package handlers



import (
	"net/http"
	
	
	"working/api/data"
)

// swagger: route GET /products products ListProducts 
// Returns a list of products
// responses:
// 200: productsResponse


// GetProducts returns the data from the datastore
func(p *Products) GetProducts(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle GET Products")

	// fetch the data from the datastore
	lp := data.GetProducts()

	// serialize the list to JSON
	err := lp.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to marhsal JSON", http.StatusInternalServerError)
		return
	}
}