package models

type User struct{
	ID string `json:"id"  gorm:"primary_key"`
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	Email string `json:"email"`
	PasswordHash string `json:"password"`
	Address []Address `json:"addresses" grom:"foreignkey:UserID"`
}

type Address struct{
	ID string `json:"id" gorm:"primary_key"`
	UserID string `json:"user_id" gorm:"index"`
	Street string `json:"street"`
	City string `json:"city"`
	State string `json:"state"`
	PinCode string `json:"pin_code"`
}