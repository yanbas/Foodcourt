package app

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
		Type            string `json:"customer_name"`
		ReferenceNumber int    `json:"qty"`
		OrderId         string `json:"menu"`
		Amount          int    `json:"amount"`
		Status          string `json:"status"`
	}

	ResultData struct {
		Code    int         `json:"code"`
		Success bool        `json:"success"`
		Data    interface{} `json:"data"`
	}
)
