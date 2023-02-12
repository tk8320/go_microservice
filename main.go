package main

import (
	"log"
	"net/http"

	route "go_microservice/routes"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func main() {

	// Defining mux rounter to handle the rounting
	r := mux.NewRouter()
	log.Println("Starting the server...")

	// Initializing the context handlers
	ctx := route.InitContext()

	// Handling the rountes
	r.HandleFunc("/order/{id:[0-9]+}", ctx.ViewOrderByID).Methods("GET")
	r.HandleFunc("/order/create", ctx.CreateOrder).Methods("POST")
	r.HandleFunc("/order", ctx.ViewOrder).Methods("GET")
	r.HandleFunc("/order/{id:[0-9]+}", ctx.DeleteOrder).Methods("DELETE")
	r.HandleFunc("/order/{id:[0-9]+}", ctx.UpdateOrder).Methods("PUT")
	r.HandleFunc("/order/search", ctx.SearchOrder).Methods("POST")

	// Assigning the path to handler
	// Using Middleware for logging
	r.Use(route.LoggingMiddleware)
	http.Handle("/", r)

	// Address binding. In this case every IP of my machine
	// Second parameter is handler
	log.Fatal(http.ListenAndServe(":8080", nil))
}
