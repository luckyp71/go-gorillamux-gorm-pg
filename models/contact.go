package models

type Contact struct {
	ContactID   int  `gorm:"primary_key" json:"contact_id"`
	CountryCode int  `json:"country_code"`
	MobileNo    uint `json:"mobile_no"`
	CustId      int  `json:"cust_id"`
}
