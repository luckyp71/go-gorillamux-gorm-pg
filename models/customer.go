package models

type Customer struct {
	CustomerID   int       `gorm:"primary_key" json:"customer_id"`
	CustomerName string    `json:"customer_name"`
	Contacts     []Contact `gorm:"ForeignKey:CustId" json:"contacts"`
}
