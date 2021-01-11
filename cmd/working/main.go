package main

import (
	"os/signal"
	"context"
	"time"
	"log"
	"net/http"
	"os"


	"github.com/gorilla/mux"


	"working/cmd/working/handlers"
)

// BindAddress represents config for server connection 
var BindAddress = env.String("BIND_ADDRESS", false, ":9090", "Bind address for the server")

func main() {

	env.Parse()


	l := log.New(os.Stdout, "product-api", log.LstdFlags)

	// create the handlers
	ph := handlers.NewProducts(l)

	// create a new serve mux and register the handlers
	sm := mux.NewRouter()

	getRouter := sm.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/", ph.GetProducts)

	putRouter := sm.Methods(http.MethodPut).Subrouter()
	putRouter.HandleFunc("/(id:[0-9]+)", ph.UpdateProducts)
	putRouter.Use(ph.MiddlewareValidateProduct)

	postRouter := sm.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("/", ph.AddProduct)
	postRouter.Use(ph.MiddlewareValidateProduct)


	// create a new server
	s := &http.Server{
		Addr: *BindAddress, 			// config 
		Handler: sm,				// set the default handler
		ErrorLog: l,				// set the logger for the server
		WriteTimeout: 10 * time.Second,		// max time to write response to the client
		ReadTimeout: 5 * time.Second,		// max time to read request from the client
		IdleTimeout: 120 * time.Second,		// max time for connection using TCP Keep-Alive
	}

	// start the server
	go func() {
		l.Printf("Starting server on port 9090")

		err := s.ListenAndServe()
		if err != nil {
			l.Printf("Error starting server: %s/n", err)
			os.Exit(1)
		}
	}()

	// trap sigterm or interrupt and gracefully shutdown the server
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, os.Kill)

	sig := <-c
	l.Println("Receive termenate , graceful shutdown", sig)

	tc, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(tc)
} 
