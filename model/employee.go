package model

import "time"

type GetOrdersFilter struct {
	OrderName    string `query:"order_name"`
	CustomerName string `query:"customer_name"`
	ID           *int   `query:"id"`
	PageRequest  PageRequest
}

type GetOrdersResult struct {
	ID              int       `json:"id"`
	OrderDate       time.Time `json:"created_at"`
	OrderName       string    `json:"order_name"`
	CustomerName    string    `json:"customer_name"`
	CustomerCompany string    `json:"customer_company"`
	DeliveredAmount *uint     `json:"delivered_amount"`
	TotalAmount     uint      `json:"total_amount"`
}

type CreateOrderRequest struct {
	OrderName       string `json:"order_name" validate:"required,notblank,alphanum,min=3,max=60"`
	CustomerName    string `json:"customer_name" validate:"required,notblank,min=3,max=60"`
	CustomerCompany string `json:"customer_company" validate:"required,notblank,min=3,max=60"`
	TotalAmount     uint   `json:"total_amount" validate:"required,numeric"`
	DeliveredAmount uint   `json:"delivered_amount" validate:"omitempty,numeric"`
}

type CreateOrderResult struct {
	ID              int       `json:"id"`
	OrderDate       time.Time `json:"created_at"`
	OrderName       string    `json:"order_name"`
	CustomerName    string    `json:"customer_name"`
	CustomerCompany string    `json:"customer_company"`
	DeliveredAmount *uint     `json:"delivered_amount"`
	TotalAmount     uint      `json:"total_amount"`
}

type GetOrderByIDRequest struct {
	OrderID int `param:"id" validate:"required"`
}

type DeleteOrderByIDRequest struct {
	OrderID int `param:"id" validate:"required"`
}

type GetOrderByIDResult struct {
	ID              int       `json:"id"`
	OrderDate       time.Time `json:"created_at"`
	OrderName       string    `json:"order_name"`
	CustomerName    string    `json:"customer_name"`
	CustomerCompany string    `json:"customer_company"`
	DeliveredAmount *uint     `json:"delivered_amount"`
	TotalAmount     uint      `json:"total_amount"`
}

type EditOrderRequest struct {
	OrderID int `param:"id" validate:"required"` // Path variable

	OrderName       string `json:"order_name" validate:"required,notblank,alphanum,min=3,max=60"`
	CustomerName    string `json:"customer_name" validate:"required,notblank,min=3,max=60"`
	CustomerCompany string `json:"customer_company" validate:"required,notblank,min=3,max=60"`
	TotalAmount     uint   `json:"total_amount" validate:"required,numeric"`
	DeliveredAmount uint   `json:"delivered_amount" validate:"omitempty,numeric"`
}

type EditOrderResult struct {
	ID              int       `json:"id"`
	OrderDate       time.Time `json:"created_at"`
	OrderName       string    `json:"order_name"`
	CustomerName    string    `json:"customer_name"`
	CustomerCompany string    `json:"customer_company"`
	DeliveredAmount *uint     `json:"delivered_amount"`
	TotalAmount     uint      `json:"total_amount"`
}
