package app

import "time"

type (
	Menu struct {
		Name   string `json:"name"`
		Price  int    `json:"price"`
		Status int    `json:"status"`
	}

	Order struct {
		CustomerName string `json:"customer_name"`
		Qty          int    `json:"qty"`
		Menu         int    `json:"menu"`
		TableNumber  int    `json:"table_number"`
		Status       int    `json:"status"`
	}

	Payment struct {
		Type            string `json:"type"`
		ReferenceNumber int    `json:"ref_no"`
		OrderId         int    `json:"order_id"`
		Amount          int    `json:"amount"`
		Status          int    `json:"status"`
	}

	User struct {
		Username string `json:"username"`
		Password int    `json:"Password"`
		Name     string `json:"name"`
		Level    int    `json:"level"`
		Status   int    `json:"status"`
	}

	ResponseToken struct {
		Token string `json:"token"`
	}

	ResultData struct {
		Code    int         `json:"code"`
		Success bool        `json:"success"`
		Data    interface{} `json:"data"`
	}

	ShowPayment struct {
		ID           int       `json:"id"`
		CreatedAt    time.Time `json:"created_at"`
		UpdatedAt    time.Time `json:"updated_at"`
		Qty          int       `json:"qty"`
		CustomerName string    `json:"customer_name"`
		TableNumber  int       `json:"table_number"`
		MenuName     string    `json:"menu_name"`
		Price        int       `json:"price"`
		SubTotal     int       `json:"sub_total"`
	}
)
