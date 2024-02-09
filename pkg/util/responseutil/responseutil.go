package responseutil

import (
	"fullstack_api_test/model"
	"fullstack_api_test/pkg/config"
	"github.com/labstack/echo/v4"
	"net/http"

	pkgerror "fullstack_api_test/pkg/error"
)

func CreateSuccessResponse(data interface{}, pagination *model.Pagination) model.ResponseBody {
	return model.ResponseBody{
		Status:     "SUCCESS",
		Code:       "0000",
		Data:       data,
		Pagination: pagination,
	}
}

func CreateErrorResponse(err pkgerror.CustomError) model.ResponseBody {
	body := model.ResponseBody{
		Status:       "ERROR",
		Code:         err.Code,
		ErrorMessage: &err.Msg,
	}
	if !config.Data.IsEnvProduction() && err.Err != nil {
		e := err.Err.Error()
		body.ErrorRemark = &e
	}
	return body
}

func SendSuccessReponse(ctx echo.Context, data interface{}, pagination *model.Pagination) error {
	return ctx.JSON(http.StatusOK, CreateSuccessResponse(data, pagination))
}

func SendErrorResponse(ctx echo.Context, err pkgerror.CustomError) error {
	return ctx.JSON(err.HttpCode, CreateErrorResponse(err))
}
