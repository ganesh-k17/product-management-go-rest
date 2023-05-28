package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

var app App

func TestMain(m *testing.M) {
	err := app.Initialise(DBUser, DBPassword, "test")
	if err != nil {
		log.Fatal("Error occured while initialising the database.")
	}
	createTable()
	m.Run()
}

func createTable() {
	createTableQuery := ` CREATE TABLE IF NOT EXISTS products (
		id int NOT NULL AUTO_INCREMENT,
		name varchar(255) NOT NULL,
		quantity int,
		price float(10,7),
		PRIMARY KEY(id)
	);`

	_, err := app.DB.Exec(createTableQuery)
	if err != nil {
		log.Fatal(err)
	}
}

func clearTable() {
	app.DB.Exec("DELETE from products")
	app.DB.Exec("ALTER table products AUTO_INCREMENT=1")
}

func addProduct(name string, quantity int, price float64) {
	query := fmt.Sprintf("insert into products(name, quantity, price) values('%v', %v, %v)", name, quantity, price)
	app.DB.Exec(query)
}

func checkStatusCode(t *testing.T, expectedStatusCode int, actualStatusCode int) {
	if expectedStatusCode != actualStatusCode {
		t.Errorf("Expected status: %v, Received: %v", expectedStatusCode, actualStatusCode)
	}
}

func sendRequest(request *http.Request) *httptest.ResponseRecorder {
	recorder := httptest.NewRecorder()
	app.Router.ServeHTTP(recorder, request)
	return recorder
}

func TestGetProducts(t *testing.T) {
	log.Printf("Test Getting products")
	clearTable()
	addProduct("keyboard", 100, 500)
	request, _ := http.NewRequest("GET", "/product/1", nil)
	response := sendRequest(request)
	checkStatusCode(t, http.StatusOK, response.Code)
}

func TestCreateProduct(t *testing.T) {
	log.Printf("Test Posting product")
	clearTable()
	var product = []byte(`{"name":"chair","quantity":1,"price":100}`)
	req, _ := http.NewRequest("POST", "/product", bytes.NewBuffer(product))
	req.Header.Set("Contenty_Type", "application/json")

	response := sendRequest(req)
	checkStatusCode(t, http.StatusOK, response.Code)

	var m map[string]interface{}
	json.Unmarshal(response.Body.Bytes(), &m)

	if m["name"] != "chair" {
		t.Errorf("Expected name: %v, Got %v", "chair", m["name"])
	}
	if m["quantity"] != 1.0 {
		t.Errorf("Expected Quantity: %v, Got %v", 1.0, m["quantity"])
	}
}

func TestDeleteProduct(t *testing.T) {
	log.Printf("Test Deleting product")
	clearTable()
	addProduct("connector", 2, 10)

	req, _ := http.NewRequest("DELETE", "/product/1", nil)
	response := sendRequest(req)
	checkStatusCode(t, http.StatusOK, response.Code)

	req, _ = http.NewRequest("GET", "/product/1", nil)
	response = sendRequest(req)
	checkStatusCode(t, http.StatusNotFound, response.Code)
}

func TestUpdateProduct(t *testing.T) {
	log.Printf("Test Updating product")
	clearTable()

	// var product = []byte(`{"name":"chair","quantity":1,"price":100}`)
	// req, _ := http.NewRequest("POST", "/product", bytes.NewBuffer(product))
	// req.Header.Set("Contenty_Type", "application/json")
	// sendRequest(req) // Adding Product

	addProduct("chair", 1, 100)

	product := []byte(`{"name":"Connector","quantity":1,"price":100}`)
	req, _ := http.NewRequest("PUT", "/product/1", bytes.NewBuffer(product))
	req.Header.Set("Contenty_Type", "application/json")
	response := sendRequest(req) // Updating the product

	checkStatusCode(t, http.StatusOK, response.Code)

	var m map[string]interface{}
	json.Unmarshal(response.Body.Bytes(), &m)

	oldValue := "chair"

	if m["name"] != "Connector" {
		t.Errorf("Expected name: %v, Got %v", oldValue, m["name"])
	}
}
