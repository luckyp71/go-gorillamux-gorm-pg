package main

import (
	"net/http"
	"encoding/json"
	"fmt"
	"github.com/jinzhu/gorm"
	_"github.com/jinzhu/gorm/dialects/postgres"
	"github.com/gorilla/mux"
)

var db *gorm.DB
var e error

/*************** Models ****************/
// Customer 
type Customer struct {
	CustomerID int `gorm:"primary_key" json:"customer_id"`
	CustomerName string	`json:"customer_name"`
	Contacts []Contact `gorm:"ForeignKey:CustId" json:"contacts"`
}
// Contact
type Contact struct {
	ContactID int `gorm:"primary_key" json:"contact_id"`
	CountryCode int	`json:"country_code"`
	MobileNo uint `json:"mobile_no"`
	CustId int `json:"cust_id"`
}
/********************************************/

/************ Main method for our service **************/
func main(){	
	db, e = gorm.Open("postgres", "user=postgres password=pratama dbname=postgres sslmode=disable")
	if e != nil {
		fmt.Println(e)
	} else {
		fmt.Println("Connection Established")
	}
	defer db.Close()
	db.SingularTable(true)
	db.AutoMigrate(Customer{}, Contact{})
	db.Model(&Contact{}).AddForeignKey("cust_id", "customer(customer_id)","CASCADE","CASCADE")


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
func getCustomers(w http.ResponseWriter, r *http.Request){
	var customers []Customer
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
func getCustomersByName(w http.ResponseWriter, r *http.Request){
	var customers []Customer
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
		json.NewEncoder(w).Encode(customers)
	}
}

// Get customer by id
func getCustomerById(w http.ResponseWriter, r *http.Request){
	var customer Customer
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
		json.NewEncoder(w).Encode(customer)
	}
}

// Insert cusotmer
func insertCustomer(w http.ResponseWriter, r *http.Request){
	var customer Customer
	var
	_= json.NewDecoder(r.Body).Decode(&customer)
	db.Create(&customer)
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Response-Code", "00")
	w.Header().Set("Response-Desc", "Success")
	json.NewEncoder(w).Encode(&customer)
}

// Update customer
func updateCustomer(w http.ResponseWriter, r * http.Request){
	var customer Customer
	param := mux.Vars(r)
	if e := db.Where("customer_id = ?", param["id"]).Preload("Contacts").First(&customer).Error; e != nil {
		w.Header().Set("Content-Type", "application-json")
		w.Header().Set("Response-Code", "06")
		w.Header().Set("Response-Desc", "Data Not Found")
		w.WriteHeader(404)
		w.Write([]byte(`{"message":"data not found"}`))
	} else {
		_= json.NewDecoder(r.Body).Decode(&customer)
		db.Save(&customer)
		w.Header().Set("Content-Type", "application-json")
		w.Header().Set("Response-Code", "00")
		w.Header().Set("Response-Desc", "Success")
		json.NewEncoder(w).Encode(&customer)
	}
}

// Delete customer
func deleteCustomer(w http.ResponseWriter, r * http.Request){
	var customer Customer
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