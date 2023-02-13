package routes

import (
	"database/sql"
	"encoding/json"
	"fmt"
	model "go_microservice/models"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

type Context struct {
	db *sql.DB
}

// Initialize the database connections
func InitContext() Context {
	log.Println("Initiating the database connection")
	dsn := fmt.Sprintf("go_root:%s@tcp(db4free.net:3306)/go_microservice?charset=utf8&parseTime=True", "Password!123")
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalln("Unable to connect to the database. Closing service...")
	}

	err = db.Ping()
	if err != nil {
		log.Fatalln("Unable to connect to the database. Closing service...")
	}
	return Context{db}
}

// Update the Order with id.
// Show error if order with given ID Does not exists
func (ctx Context) UpdateOrder(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	op, _ := io.ReadAll(r.Body)
	var order model.OrderQuery
	query := "UPDATE tbl_orders SET "
	query_params := []string{}
	err := json.Unmarshal(op, &order)
	if err != nil {
		log.Println("Unable to Unmarshal")
		HandleError(w, r, err, "Unable to Unmarshal", http.StatusBadRequest)
		return
	}
	if order.Status != nil {
		query_params = append(query_params, fmt.Sprintf("status = '%s'", *order.Status))
	}
	if order.Currency != nil {
		query_params = append(query_params, fmt.Sprintf("currency = '%s'", *order.Currency))
	}

	query += fmt.Sprintf("%s WHERE id = %d", strings.Join(query_params, ", "), id)

	log.Println(query)
	log.Printf("Updating the order with id : %d\n", id)

	_, err = ctx.db.Exec(query)

	if err != nil {
		log.Println("Unable to update : ", err)
		HandleError(w, r, err, "Unable To Update..", http.StatusBadRequest)
		return
	}

	WriteResponse(w, r, []byte("Updated Successfully"), http.StatusCreated)
}

/*
	Search the order with available parameters and sorting options.

Filtering options --> status, currency
Data ordering and limiting options --> orderAsc, orderDesc, limit
*/
func (ctx Context) SearchOrder(w http.ResponseWriter, r *http.Request) {

	request_body, _ := io.ReadAll(r.Body)
	var search_option model.SearchOptions
	var search_value model.SearchValues
	var OrderList []model.Order
	search_value_list := []string{}
	json.Unmarshal(request_body, &search_option)
	json.Unmarshal(request_body, &search_value)

	// Creating SQL Query from give parameters
	search_query := "SELECT * FROM tbl_orders"

	if search_value.Status != nil || search_value.Currency != nil {
		search_query += " WHERE "
	}

	if search_value.Status != nil {
		search_value_list = append(search_value_list, fmt.Sprintf("status = '%s'", *search_value.Status))
	}

	if search_value.Currency != nil {
		search_value_list = append(search_value_list, fmt.Sprintf("currency = '%s'", *search_value.Currency))
	}

	search_query += strings.Join(search_value_list, " AND ")

	if search_option.OrderASC != nil {
		search_query += fmt.Sprintf(" ORDER BY %s ASC", *search_option.OrderASC)
	}

	if search_option.OrderDESC != nil {
		search_query += fmt.Sprintf(" ORDER BY %s DESC", *search_option.OrderDESC)
	}

	if search_option.Limit != nil {
		search_query += fmt.Sprintf(" LIMIT %d", *search_option.Limit)
	}
	log.Println(search_query)

	res, err := ctx.db.Query(search_query)
	if err != nil {
		HandleError(w, r, err, "Error Querying data", http.StatusBadRequest)
	}

	for res.Next() {
		var op model.Order
		var temp []byte
		err := res.Scan(&op.Id, &op.Status, &temp, &op.Total, &op.Currency)
		if err != nil {
			HandleError(w, r, err, "No Data Found", http.StatusNotFound)
		}

		err = json.Unmarshal(temp, &op.Items)
		if err != nil {
			HandleError(w, r, err, "Unable to unmarshal data", http.StatusBadRequest)
		}

		OrderList = append(OrderList, op)
	}

	if len(OrderList) == 0 {
		WriteResponse(w, r, []byte("No Data Available"), http.StatusNotFound)
		return
	}
	string_op, _ := json.Marshal(OrderList)
	WriteResponse(w, r, string_op, http.StatusOK)
}

// Delete order with give ID. Return 404 Not found if order dosen't exist
func (ctx Context) DeleteOrder(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	query := fmt.Sprintf("DELETE from tbl_orders WHERE id = %d", id)
	res, err := ctx.db.Exec(query)
	count, _ := res.RowsAffected()
	if count == 0 {
		log.Println("Data Not Availabe")
		WriteResponse(w, r, []byte("Data Not Availabe"), http.StatusNotFound)
		return
	}
	if err != nil {
		log.Println("No Record found")
		HandleError(w, r, err, "No Record found", http.StatusNotFound)
		return
	}
	log.Println("Deleted Successfully")
	WriteResponse(w, r, []byte("Deleted Successfully"), http.StatusOK)
}

// Create Order. Returns 201 Created after successful creation
func (ctx Context) CreateOrder(w http.ResponseWriter, r *http.Request) {
	log.Println("Creating Orders")
	request_body, _ := io.ReadAll(r.Body)
	var order_data model.Order
	err := json.Unmarshal(request_body, &order_data)
	if err != nil {
		log.Println("Json Unmarshalling failed :  ", err)
		WriteResponse(w, r, []byte("Json Marshalling failed"), http.StatusBadRequest)
		return
	} else {
		item_string, _ := json.Marshal(order_data.Items)
		query := fmt.Sprintf("INSERT INTO tbl_orders(status, items, total, currency) VALUES ('%s','%s',%f,'%s')", order_data.Status, item_string, order_data.Total, order_data.Currency)
		log.Println(query)
		_, err = ctx.db.Query(query)
		if err != nil {
			log.Println(err)
			WriteResponse(w, r, []byte("Insert failed"), http.StatusBadRequest)
			return
		}
	}
	WriteResponse(w, r, []byte("Data Created Successfully"), http.StatusCreated)
}

// Shows all Orders
func (ctx Context) ViewOrder(w http.ResponseWriter, r *http.Request) {

	var OrderList []model.Order
	query := "SELECT * FROM tbl_orders"
	res, err := ctx.db.Query(query)
	if err != nil {
		HandleError(w, r, err, "Error Querying data", http.StatusBadRequest)
	}

	for res.Next() {
		var op model.Order
		var temp []byte
		err := res.Scan(&op.Id, &op.Status, &temp, &op.Total, &op.Currency)
		if err != nil {
			HandleError(w, r, err, "No Data Found", http.StatusNotFound)
		}

		err = json.Unmarshal(temp, &op.Items)
		if err != nil {
			HandleError(w, r, err, "Unable to unmarshal data", http.StatusBadRequest)
		}

		OrderList = append(OrderList, op)
	}
	if len(OrderList) == 0 {
		WriteResponse(w, r, []byte("No Data Available"), http.StatusNotFound)
		return
	}
	string_op, _ := json.Marshal(OrderList)
	WriteResponse(w, r, string_op, http.StatusOK)
}

// Show order by id. Return 404 Not found if order is not available
func (ctx Context) ViewOrderByID(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	query := fmt.Sprintf("Select * from tbl_orders where id = %d", id)
	res := ctx.db.QueryRow(query)
	var op model.Order
	var temp []byte
	err := res.Scan(&op.Id, &op.Status, &temp, &op.Total, &op.Currency)
	if err != nil {
		HandleError(w, r, err, "No Data Found", http.StatusNotFound)
		return
	}

	err = json.Unmarshal(temp, &op.Items)
	if err != nil {
		HandleError(w, r, err, "Unable to unmarshal data", http.StatusBadRequest)
		return
	}

	byte_op, _ := json.Marshal(op)
	WriteResponse(w, r, byte_op, http.StatusOK)

}

// Middleware for logging the requests
func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Do stuff here
		log.Printf("[%s] %s", r.Method, r.RequestURI)
		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(w, r)
	})
}

// Adding custom headers to the request
func WriteResponse(w http.ResponseWriter, r *http.Request, data []byte, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(data)
}

// Handle Errors and send to the response writer
func HandleError(w http.ResponseWriter, r *http.Request, err error, message string, status int) {
	log.Println("Error : ", err.Error())
	WriteResponse(w, r, []byte(message), status)

}
