package main

import (
	"bytes"
	route "go_microservice/routes"
	"net/http"
	"net/http/httptest"
	"testing"
)

var A = App{route.InitContext()}
var Router = A.CreateHandler()

func TestConnection(t *testing.T) {
	req, err := http.NewRequest("GET", "/hello", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := Router
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected := `Hello!`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestCreateOrder(t *testing.T) {
	data := []byte(`{
				"id": "5",
				"status": "TEST_ORDER",
				"items": [
					{
						"id": 13,
						"description": "Fruit Salad",
						"price": 29.99,
						"quantity": 1
					}],
				"total": 29.99,
				"currencyUnit": "USD"
				}`)
	req, err := http.NewRequest("POST", "/order/create", bytes.NewBuffer(data))
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := Router
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := `Data Created Successfully`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestGetAllOrder(t *testing.T) {
	req, err := http.NewRequest("GET", "/order", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := Router
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

func TestGetOrderByID(t *testing.T) {

	req, err := http.NewRequest("GET", "/order/2", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := Router
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	expected := `{"id":"2","status":"ORDER_UPDATED","items":[{"id":2,"Description":"Comb","price":20,"quantity":1},{"id":5,"Description":"Hair Oil","price":10,"quantity":2}],"total":40,"currencyUnit":"INR"}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestUpdateOrder(t *testing.T) {
	body := []byte(`{"status" : "ORDER_UPDATED"}`)
	req, err := http.NewRequest("PUT", "/order/3", bytes.NewBuffer(body))
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := Router
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusCreated)
	}

	expected := "Updated Successfully"
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestDeleteOrder(t *testing.T) {

	req, err := http.NewRequest("DELETE", "/order/4", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := Router
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	expected := `Deleted Successfully`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestSearchOrder(t *testing.T) {
	data := []byte(`{"status":"TEST_ORDER"}`)
	req, err := http.NewRequest("POST", "/order/search", bytes.NewBuffer(data))
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := Router
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}
