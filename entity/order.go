package entity

import (
	"time"
)

type Order struct {
	ID              uint      `gorm:"primary_key"`
	OrderName       string    `json:"order_name" db:"order_name"`
	CustomerName    string    `json:"customer_name" db:"customer_name"`
	CustomerCompany string    `json:"customer_company" db:"customer_company"`
	DeliveredAmount *float32  `json:"delivered_amount" db:"delivered_amount"`
	TotalAmount     float32   `json:"total_amount" db:"total_amount"`
	OrderDate       time.Time `json:"order_date" db:"order_date"`
}

func (Order) TableName() string {
	return "orders"
}
