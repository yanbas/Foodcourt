package model

import "github.com/jinzhu/gorm"

type Menu struct {
	gorm.Model
	Name   string `json:"name"`
	Price  int    `json:"price"`
	Status int    `json:"status"`
}

type Order struct {
	gorm.Model
	MenuId       int    `json:"menu_id"`
	Qty          int    `json:"qty"`
	CustomerName string `json:"customer_name"`
	TableNumber  int    `json:"table_number"`
	Status       int    `json:"status"`
}

type Payment struct {
	gorm.Model
	Type            string `json:"type"`
	ReferenceNumber int    `json:"ref_no"`
	OrderId         int    `json:"order_id"`
	Amount          int    `json:"amount"`
	Status          int    `json:"status"`
}

type User struct {
	gorm.Model
	Username string `json:"username"`
	Level    int    `json:"level"`
	Password int    `json:"password"`
	Name     string `json:"name"`
	Status   int    `json:"status"`
}
