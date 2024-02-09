package constant

import "errors"

type OrderColumn string

const (
	OrderColumnOrderName       OrderColumn = "order_name"
	OrderColumnCustomerName    OrderColumn = "customer_name"
	OrderColumnCustomerCompany OrderColumn = "customer_company"
	OrderColumnOrderDate       OrderColumn = "order_date"
	OrderColumnDeliveredAmount OrderColumn = "delivered_amount"
	OrderColumnTotalAmount     OrderColumn = "total_amount"
)

var OrderColumns = []OrderColumn{
	OrderColumnOrderName,
	OrderColumnCustomerName,
	OrderColumnCustomerCompany,
	OrderColumnOrderDate,
	OrderColumnDeliveredAmount,
	OrderColumnTotalAmount,
}

func ParseOrderColumnName(str string) (OrderColumn, error) {
	for _, t := range OrderColumns {
		if str == string(t) {
			return t, nil
		}
	}
	return "", errors.New(str)
}
