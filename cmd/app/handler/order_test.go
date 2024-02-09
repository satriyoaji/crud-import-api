package handler

//import (
//	mocks "fullstack_api_test/mocks/service"
//	"fullstack_api_test/model"
//	pkgerror "fullstack_api_test/pkg/error"
//	"fullstack_api_test/pkg/util/jsonutil"
//	"fullstack_api_test/pkg/util/responseutil"
//	pkgvalidator "fullstack_api_test/pkg/validator"
//	"github.com/go-playground/validator/v10"
//	"github.com/go-playground/validator/v10/non-standard/validators"
//	"github.com/labstack/echo/v4"
//	"github.com/shopspring/decimal"
//	"github.com/stretchr/testify/assert"
//	"github.com/stretchr/testify/mock"
//	"net/http"
//	"net/http/httptest"
//	"strings"
//	"testing"
//)
//
////func TestGetEmployees(t *testing.T) {
////	getOneEmployeeResult := model.GetOrdersResult{
////		ID:        1,
////		OrderName: "First Employee 0",
////		CustomerName:  "Last Name 0",
////		Email:     "employee@email.com",
////	}
////	getEmployeesResult := []model.GetOrdersResult{
////		getOneEmployeeResult,
////	}
////	testCases := []struct {
////		Name                 string
////		InitHandler          func(ctx echo.Context, ps *mocks.EmployeeService) *Handler
////		Json                 string
////		ExpectedHttpCode     int
////		ExpectedResponseBody model.ResponseBody
////	}{
////		{
////			Name: "GetEmployees/SystemError",
////			InitHandler: func(ctx echo.Context, ps *mocks.EmployeeService) *Handler {
////				reqParam := model.GetOrdersFilter{} // if needed
////				ps.On("GetEmployees", ctx, reqParam).Return(nil, pkgerror.ErrSystemError)
////				return NewHandler(ps)
////			},
////			ExpectedHttpCode:     http.StatusInternalServerError,
////			ExpectedResponseBody: responseutil.CreateErrorResponse(pkgerror.ErrSystemError),
////		},
////		{
////			Name: "GetEmployees/Success",
////			InitHandler: func(ctx echo.Context, ps *mocks.EmployeeService) *Handler {
////				reqParam := model.GetOrdersFilter{} // if needed
////				ps.On("GetEmployees", ctx, reqParam).Return(&getEmployeesResult, pkgerror.NoError)
////				return NewHandler(ps)
////			},
////			ExpectedHttpCode:     http.StatusOK,
////			ExpectedResponseBody: responseutil.CreateSuccessResponse(&getEmployeesResult, nil),
////		},
////	}
////	for _, tc := range testCases {
////		t.Run(tc.Name, func(t *testing.T) {
////			e := echo.New()
////			e.Validator = pkgvalidator.New(validator.New())
////			q := make(url.Values)
////
////			req := httptest.NewRequest(http.MethodGet, "/v1/employees?"+q.Encode(), nil)
////			res := httptest.NewRecorder()
////			c := e.NewContext(req, res)
////			ps := new(mocks.EmployeeService)
////			h := tc.InitHandler(c, ps)
////			if assert.NoError(t, h.GetEmployees(c)) {
////				assert.Equal(t, tc.ExpectedHttpCode, res.Code)
////				expected := tc.ExpectedResponseBody
////				jsonpath, err := jsonutil.NewJsonPath(res.Body.String())
////				assert.Nil(t, err)
////				assert.Equal(t, expected.Status, jsonpath.GetString("status"))
////				assert.Equal(t, expected.Code, jsonpath.GetString("code"))
////				assert.Equal(t, expected.ErrorMessage, jsonpath.GetStringPtr("error_message"))
////				//assert.Equal(t, expected.ErrorDebug, jsonpath.GetStringPtr("error_debug"))
////				if expected.Data != nil {
////					data := expected.Data.(*[]model.GetOrdersResult)
////					for i, m := range *data {
////						assert.Equal(t, m.ID, jsonpath.GetIntf("data[%d].id", i))
////						assert.Equal(t, m.OrderName, jsonpath.GetStringf("data[%d].first_name", i))
////						assert.Equal(t, m.CustomerName, jsonpath.GetStringf("data[%d].last_name", i))
////					}
////				}
////			}
////			ps.AssertExpectations(t)
////		})
////	}
////}
//
//func TestGetEmployeeByID(t *testing.T) {
//	result := model.GetOrderByIDResult{
//		ID: 1,
//	}
//	testCases := []struct {
//		Name                 string
//		InitHandler          func(ctx echo.Context, s *mocks.EmployeeService) *Handler
//		PathEmployeeID       string
//		ExpectedHttpCode     int
//		ExpectedResponseBody model.ResponseBody
//	}{
//		{
//			Name: "InvalidParams",
//			InitHandler: func(ctx echo.Context, s *mocks.EmployeeService) *Handler {
//				return NewHandler(s)
//			},
//			ExpectedHttpCode:     http.StatusBadRequest,
//			ExpectedResponseBody: responseutil.CreateErrorResponse(pkgerror.ErrInvalidParams),
//		},
//		{
//			Name: "ServiceError",
//			InitHandler: func(ctx echo.Context, s *mocks.EmployeeService) *Handler {
//				s.On("GetEmployeeByID", mock.Anything, mock.Anything).Return(nil, pkgerror.ErrSystemError)
//				return NewHandler(s)
//			},
//			PathEmployeeID:       "1",
//			ExpectedHttpCode:     http.StatusInternalServerError,
//			ExpectedResponseBody: responseutil.CreateErrorResponse(pkgerror.ErrSystemError),
//		},
//		{
//			Name: "Success",
//			InitHandler: func(ctx echo.Context, s *mocks.EmployeeService) *Handler {
//				s.On("GetEmployeeByID", mock.Anything, mock.Anything).Return(&result, pkgerror.NoError)
//				return NewHandler(s)
//			},
//			PathEmployeeID:       "1",
//			ExpectedHttpCode:     http.StatusOK,
//			ExpectedResponseBody: responseutil.CreateSuccessResponse(&result, nil),
//		},
//	}
//	for _, tc := range testCases {
//		t.Run(tc.Name, func(t *testing.T) {
//			e := echo.New()
//			e.Validator = pkgvalidator.New(validator.New())
//			req := httptest.NewRequest(http.MethodGet, "/", nil)
//			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
//			res := httptest.NewRecorder()
//			c := e.NewContext(req, res)
//			c.SetPath("/v1/employees/:id")
//			c.SetParamNames("id")
//			c.SetParamValues(tc.PathEmployeeID)
//			s := new(mocks.EmployeeService)
//			h := tc.InitHandler(c, s)
//			if assert.NoError(t, h.GetEmployeeByID(c)) {
//				assert.Equal(t, tc.ExpectedHttpCode, res.Code)
//				expected := tc.ExpectedResponseBody
//				jsonpath, err := jsonutil.NewJsonPath(res.Body.String())
//				assert.Nil(t, err)
//				assert.Equal(t, expected.Status, jsonpath.GetString("status"))
//				assert.Equal(t, expected.Code, jsonpath.GetString("code"))
//				assert.Equal(t, expected.ErrorMessage, jsonpath.GetStringPtr("error_message"))
//				if expected.Data != nil {
//					data := expected.Data.(*model.GetOrderByIDResult)
//					assert.Equal(t, data.ID, jsonpath.GetInt("data.id"))
//				}
//			}
//			s.AssertExpectations(t)
//		})
//	}
//}
//
//func TestAddEmployee(t *testing.T) {
//	validJson := `{
//		"first_name": "Ryo",
//		"last_name": "Ajeee",
//		"email": "ryoaji27@gmail.com",
//		"hire_date": "2023-06-27"
//	}`
//	result := model.CreateOrderResult{
//		ID: 1,
//	}
//	testCases := []struct {
//		Name                 string
//		InitHandler          func(ctx echo.Context, s *mocks.EmployeeService) *Handler
//		Json                 string
//		ExpectedHttpCode     int
//		ExpectedResponseBody model.ResponseBody
//	}{
//		{
//			Name: "InvalidParams",
//			InitHandler: func(ctx echo.Context, s *mocks.EmployeeService) *Handler {
//				return NewHandler(s)
//			},
//			ExpectedHttpCode:     http.StatusBadRequest,
//			ExpectedResponseBody: responseutil.CreateErrorResponse(pkgerror.ErrInvalidParams),
//		},
//		{
//			Name: "ServiceError",
//			InitHandler: func(ctx echo.Context, s *mocks.EmployeeService) *Handler {
//				s.On("CreateOrder", mock.Anything, mock.Anything).Return(nil, pkgerror.ErrSystemError)
//				return NewHandler(s)
//			},
//			Json:                 validJson,
//			ExpectedHttpCode:     http.StatusInternalServerError,
//			ExpectedResponseBody: responseutil.CreateErrorResponse(pkgerror.ErrSystemError),
//		},
//		{
//			Name: "Success",
//			InitHandler: func(ctx echo.Context, s *mocks.EmployeeService) *Handler {
//				s.On("CreateOrder", mock.Anything, mock.Anything).Return(&result, pkgerror.NoError)
//				return NewHandler(s)
//			},
//			Json:                 validJson,
//			ExpectedHttpCode:     http.StatusOK,
//			ExpectedResponseBody: responseutil.CreateSuccessResponse(&result, nil),
//		},
//	}
//	for _, tc := range testCases {
//		t.Run(tc.Name, func(t *testing.T) {
//			v := validator.New()
//			v.RegisterCustomTypeFunc(pkgvalidator.DecimalValidator, decimal.Decimal{})
//			v.RegisterValidation("notblank", validators.NotBlank)
//			e := echo.New()
//			e.Validator = pkgvalidator.New(v)
//			req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(tc.Json))
//			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
//			res := httptest.NewRecorder()
//			c := e.NewContext(req, res)
//			c.SetPath("/v1/employees")
//			s := new(mocks.EmployeeService)
//			h := tc.InitHandler(c, s)
//			if assert.NoError(t, h.AddEmployee(c)) {
//				assert.Equal(t, tc.ExpectedHttpCode, res.Code)
//				expected := tc.ExpectedResponseBody
//				jsonpath, err := jsonutil.NewJsonPath(res.Body.String())
//				assert.Nil(t, err)
//				assert.Equal(t, expected.Status, jsonpath.GetString("status"))
//				assert.Equal(t, expected.Code, jsonpath.GetString("code"))
//				assert.Equal(t, expected.ErrorMessage, jsonpath.GetStringPtr("error_message"))
//				if expected.Data != nil {
//					data := expected.Data.(*model.CreateOrderResult)
//					assert.Equal(t, data.ID, jsonpath.GetInt("data.id"))
//				}
//			}
//			s.AssertExpectations(t)
//		})
//	}
//}
