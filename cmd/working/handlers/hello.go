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
	h.l.Println("hello world")
	d, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(rw, "Oops", http.StatusBadRequest)
			
		return
	}

		log.Printf("data %s", d)
}