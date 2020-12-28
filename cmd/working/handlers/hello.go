package handlers

import (
	"io/ioutil"
	"log"
	"net/http"
)

// Hello struct listening to response from NewHello func 
type Hello struct {
	l *log.Logger
}

// NewHello receive data to Hello struct 
func NewHello(l *log.Logger) *Hello {
	return &Hello{l}
}




func(h *Hello) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	h.l.Println("Handle Hello Request")


	d, err := ioutil.ReadAll(r.Body)
	if err != nil {
		h.l.Println("Error reading body", err)

		http.Error(rw, "Unable to read request body", http.StatusBadRequest)
		return
	}

		log.Printf("data %s", d)
}