package handler

import (
	"fullstack_api_test/model"
	"fullstack_api_test/service"
	"github.com/labstack/echo/v4"
)

type Handler struct {
	orderService service.OrderService
}

func NewHandler(
	orderService service.OrderService,
) *Handler {
	return &Handler{
		orderService: orderService,
	}
}

func defaultPageRequest(pr *model.PageRequest) {
	if pr.PageNum == 0 {
		pr.PageNum = 1
	}
	if pr.PageSize == 0 {
		pr.PageSize = 5
	}
}

func RegisterHandlers(e *echo.Echo, h *Handler) {

	e.GET("/orders", h.GetOrders)
	e.GET("/orders/:id", h.GetOrderByID)
	e.POST("/orders", h.AddOrder)
	e.POST("/orders-bulk", h.BulkOrdersFromCSV)
	e.PUT("/orders/:id", h.EditOrder)
	e.DELETE("/orders/:id", h.DeleteOrderByID)

}
