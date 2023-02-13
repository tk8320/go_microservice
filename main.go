package main

import (
	"fmt"
	"log"
	"net/http"

	route "go_microservice/routes"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

type App struct {
	ctx route.Context
}

func main() {
	app := App{route.InitContext()}
	// Defining mux rounter to handle the rounting
	log.Println("Starting the server...")
	router := app.CreateHandler()

	// Handling the rounter
	http.Handle("/", router)

	app.SetUpDb()

	// Address binding. In this case every IP of my machine
	// Second parameter is handler
	log.Fatal(http.ListenAndServe(":8080", nil))

}

func (app App) CreateHandler() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/order/{id:[0-9]+}", app.ctx.ViewOrderByID).Methods("GET")
	router.HandleFunc("/order/create", app.ctx.CreateOrder).Methods("POST")
	router.HandleFunc("/order", app.ctx.ViewOrder).Methods("GET")
	router.HandleFunc("/order/{id:[0-9]+}", app.ctx.DeleteOrder).Methods("DELETE")
	router.HandleFunc("/order/{id:[0-9]+}", app.ctx.UpdateOrder).Methods("PUT")
	router.HandleFunc("/order/search", app.ctx.SearchOrder).Methods("POST")
	router.HandleFunc("/hello", Hello)

	// Assigning the path to handler
	// Using Middleware for logging
	router.Use(route.LoggingMiddleware)
	// Initializing the context handlers

	return router
}

func Hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello!")
}

func (app App) SetUpDb() {
	app.ctx.Db.Exec("TRUNCATE tbl_orders")
	app.ctx.Db.Exec(`INSERT INTO tbl_orders(status, items, total, currency) VALUES ('ORDER_PLACED','[{"id":1,"Description":"Shampoo","price":150,"quantity":1},{"id":5,"Description":"Hair Oil","price":10,"quantity":2}]',170.000000,'INR')`)
	app.ctx.Db.Exec(`INSERT INTO tbl_orders(status, items, total, currency) VALUES ('ORDER_UPDATED','[{"id":2,"Description":"Comb","price":20,"quantity":1},{"id":5,"Description":"Hair Oil","price":10,"quantity":2}]',40.000000,'INR')`)
	app.ctx.Db.Exec(`INSERT INTO tbl_orders(status, items, total, currency) VALUES ('ORDER_PLACED','[{"id":13,"Description":"Butter","price":10,"quantity":5}]',50.000000,'USD')`)
	app.ctx.Db.Exec(`INSERT INTO tbl_orders(status, items, total, currency) VALUES ('ORDER_CREATED','[{"id":12,"Description":"Fruit Slushie","price":29.99,"quantity":1}]',29.990000,'USD')`)
}
