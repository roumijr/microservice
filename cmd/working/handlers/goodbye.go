package handlers

import (
	"log"
	"net/http"
)


// Goodbye implements struct for handling data
type Goodbye struct {
	l *log.Logger
}

// NewGoodbye receive data to struct Goodbye
func NewGoodbye(l *log.Logger) *Goodbye {
	return &Goodbye{l}
}

func(g *Goodbye) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	rw.Write([]byte("Byee"))
}

