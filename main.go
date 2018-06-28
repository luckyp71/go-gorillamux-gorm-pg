package main

import (
	"encoding/json"
	"fmt"
	m "go-gorillamux-gorm-pg/models"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var db *gorm.DB
var e error

func main() {
	db, e = gorm.Open("postgres", "user=postgres password=pratama dbname=postgres sslmode=disable")
	if e != nil {
		fmt.Println(e)
	} else {
		fmt.Println("Connection Established")
	}
	defer db.Close()
	db.SingularTable(true)
	db.AutoMigrate(&m.Customer{}, &m.Contact{})
	db.Model(&m.Contact{}).AddForeignKey("cust_id", "customer(customer_id)", "CASCADE", "CASCADE")
	db.Model(&m.Customer{}).AddIndex("index_customer_id_name", "customer_id", "customer_name")

	router := mux.NewRouter()
	router.HandleFunc("/customers", getCustomers).Methods("GET")
	router.HandleFunc("/customers/{id}", getCustomerById).Methods("GET")
	router.HandleFunc("/customers/{name}/list", getCustomersByName).Methods("GET")
	router.HandleFunc("/customers", insertCustomer).Methods("POST")
	router.HandleFunc("/customers/{id}", updateCustomer).Methods("PUT")
	router.HandleFunc("/customers/{id}", deleteCustomer).Methods("DELETE")
	http.ListenAndServe(":1991", router)
}

// Get customers
func getCustomers(w http.ResponseWriter, r *http.Request) {
	var customers []m.Customer
	if e := db.Preload("Contacts").Find(&customers).Error; e != nil {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Response-Code", "06")
		w.Header().Set("Response-Desc", "Data Not Found")
		w.WriteHeader(404)
		w.Write([]byte(`{"message":"data not found"}`))
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Response-Code", "00")
		w.Header().Set("Response-Desc", "Success")
		json.NewEncoder(w).Encode(customers)
	}
}

// Get customers by name
func getCustomersByName(w http.ResponseWriter, r *http.Request) {
	var customers []m.Customer
	param := mux.Vars(r)
	if e := db.Where("customer_name = ?", param["name"]).Preload("Contacts").Find(&customers).Error; e != nil {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Response-Code", "06")
		w.Header().Set("Response-Desc", "Data Not Found")
		w.WriteHeader(404)
		w.Write([]byte(`{"message":"data not found"}`))
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Response-Code", "00")
		w.Header().Set("Response-Desc", "Success")
		json.NewEncoder(w).Encode(&customers)
	}
}

// Get customer by id
func getCustomerById(w http.ResponseWriter, r *http.Request) {
	var customer m.Customer
	param := mux.Vars(r)
	if e := db.Where("customer_id = ?", param["id"]).Preload("Contacts").First(&customer).Error; e != nil {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Response-Code", "06")
		w.Header().Set("Response-Desc", "Data Not Found")
		w.WriteHeader(404)
		w.Write([]byte(`{"message":"data not found"}`))
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Response-Code", "00")
		w.Header().Set("Response-Desc", "Success")
		json.NewEncoder(w).Encode(&customer)
	}
}

// Insert cusotmer
func insertCustomer(w http.ResponseWriter, r *http.Request) {
	var customer m.Customer
	var _ = json.NewDecoder(r.Body).Decode(&customer)
	db.Create(&customer)
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Response-Code", "00")
	w.Header().Set("Response-Desc", "Success")
	json.NewEncoder(w).Encode(&customer)
}

// Update customer
func updateCustomer(w http.ResponseWriter, r *http.Request) {
	var customer m.Customer
	param := mux.Vars(r)
	if e := db.Where("customer_id = ?", param["id"]).Preload("Contacts").First(&customer).Error; e != nil {
		w.Header().Set("Content-Type", "application-json")
		w.Header().Set("Response-Code", "06")
		w.Header().Set("Response-Desc", "Data Not Found")
		w.WriteHeader(404)
		w.Write([]byte(`{"message":"data not found"}`))
	} else {
		_ = json.NewDecoder(r.Body).Decode(&customer)
		db.Save(&customer)
		w.Header().Set("Content-Type", "application-json")
		w.Header().Set("Response-Code", "00")
		w.Header().Set("Response-Desc", "Success")
		json.NewEncoder(w).Encode(&customer)
	}
}

// Delete customer
func deleteCustomer(w http.ResponseWriter, r *http.Request) {
	var customer m.Customer
	param := mux.Vars(r)
	if e := db.Where("customer_id = ?", param["id"]).Preload("Contacts").First(&customer).Error; e != nil {
		w.Header().Set("Content-Type", "application-json")
		w.Header().Set("Response-Code", "06")
		w.Header().Set("Response-Desc", "Data Not Found")
		w.WriteHeader(404)
		w.Write([]byte(`{"message":"data not found"}`))
	} else {
		db.Where("customer_id=?", param["id"]).Preload("Contacts").Delete(&customer)
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Response-Code", "00")
		w.Header().Set("Response-Desc", "Success")
	}
}
