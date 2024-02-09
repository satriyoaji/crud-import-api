package handler

import (
	"fullstack_api_test/model"
	"fullstack_api_test/pkg/util/responseutil"
	"fullstack_api_test/pkg/validator"
	"github.com/labstack/echo/v4"
)

func (h *Handler) GetOrders(ctx echo.Context) error {
	var filter model.GetOrdersFilter
	if err := validator.BindAndValidate(ctx, &filter); !err.IsNoError() {
		return responseutil.SendErrorResponse(ctx, err)
	}
	defaultPageRequest(&filter.PageRequest)
	results, err := h.orderService.GetOrders(ctx, filter)
	if err.IsNoError() {
		return responseutil.SendSuccessReponse(ctx, results, nil)
	}
	return responseutil.SendErrorResponse(ctx, err)
}

func (h *Handler) BulkOrdersFromCSV(ctx echo.Context) error {
	req := model.CSVUploadInput{}
	if err := validator.BindAndValidate(ctx, &req); !err.IsNoError() {
		return responseutil.SendErrorResponse(ctx, err)
	}
	//if err := ctx.Bind(&req); err != nil {
	//	return responseutil.SendErrorResponse(ctx, pkgerror.ErrInvalidParams.WithError(err))
	//}
	ce := h.orderService.BulkOrdersFromCSV(ctx, req)
	if ce.IsNoError() {
		return responseutil.SendSuccessReponse(ctx, nil, nil)
	}
	return responseutil.SendErrorResponse(ctx, ce)
}

func (h *Handler) AddOrder(ctx echo.Context) error {
	req := model.CreateOrderRequest{}
	if err := validator.BindAndValidate(ctx, &req); !err.IsNoError() {
		return responseutil.SendErrorResponse(ctx, err)
	}
	result, ce := h.orderService.CreateOrder(ctx, req)
	if ce.IsNoError() {
		return responseutil.SendSuccessReponse(ctx, result, nil)
	}
	return responseutil.SendErrorResponse(ctx, ce)
}

func (h *Handler) GetOrderByID(ctx echo.Context) error {
	req := model.GetOrderByIDRequest{}
	if err := validator.BindAndValidate(ctx, &req); !err.IsNoError() {
		return responseutil.SendErrorResponse(ctx, err)
	}
	result, ce := h.orderService.GetOrderByID(ctx, req)
	if ce.IsNoError() {
		return responseutil.SendSuccessReponse(ctx, result, nil)
	}
	return responseutil.SendErrorResponse(ctx, ce)
}

func (h *Handler) EditOrder(ctx echo.Context) error {
	req := model.EditOrderRequest{}
	if err := validator.BindAndValidate(ctx, &req); !err.IsNoError() {
		return responseutil.SendErrorResponse(ctx, err)
	}
	result, ce := h.orderService.EditOrder(ctx, req)
	if ce.IsNoError() {
		return responseutil.SendSuccessReponse(ctx, result, nil)
	}
	return responseutil.SendErrorResponse(ctx, ce)
}

func (h *Handler) DeleteOrderByID(ctx echo.Context) error {
	req := model.DeleteOrderByIDRequest{}
	if err := validator.BindAndValidate(ctx, &req); !err.IsNoError() {
		return responseutil.SendErrorResponse(ctx, err)
	}
	ce := h.orderService.DeleteOrderByID(ctx, req)
	if ce.IsNoError() {
		return responseutil.SendSuccessReponse(ctx, nil, nil)
	}
	return responseutil.SendErrorResponse(ctx, ce)
}
