package service

//import (
//	"context"
//	"errors"
//	"fullstack_api_test/entity"
//	mocks "fullstack_api_test/mocks/repository"
//	"fullstack_api_test/model"
//	pkgerror "fullstack_api_test/pkg/error"
//	"fullstack_api_test/pkg/util/copyutil"
//	"github.com/labstack/echo/v4"
//	"github.com/stretchr/testify/assert"
//	"github.com/stretchr/testify/mock"
//	"gorm.io/gorm"
//	"testing"
//)
//
//func getExpectedOrdersResult() *[]model.GetOrdersResult {
//	productResult := model.GetOrdersResult{
//		ID:        1,
//		OrderName: "First Order 0",
//		CustomerName:  "Last Name 0",
//		Email:     "employee@email.com",
//	}
//	return &[]model.GetOrdersResult{productResult}
//}
//
//func TestGetOrders(t *testing.T) {
//	testCases := []struct {
//		Name           string
//		InitService    func(r *mocks.Repository) OrderService
//		Context        echo.Context
//		RequestParam   model.GetOrdersFilter
//		ExpectedCount  int
//		ExpectedResult *[]model.GetOrdersResult
//		ExpectedError  pkgerror.CustomError
//	}{
//		{
//			Name: "FindAllOrdersError",
//			InitService: func(r *mocks.Repository) OrderService {
//				r.On("FindAllOrders", context.Background(), model.GetOrdersFilter{}).Return([]entity.Order{}, errors.New("database error"))
//				return NewOrderService(r)
//			},
//			Context:        createEchoContext(false),
//			RequestParam:   model.GetOrdersFilter{},
//			ExpectedCount:  len(*getExpectedOrdersResult()),
//			ExpectedResult: nil,
//			ExpectedError:  pkgerror.ErrSystemError,
//		},
//		{
//			Name: "FindAllOrdersSuccess",
//			InitService: func(r *mocks.Repository) OrderService {
//				employee := entity.Order{
//					ID:        1,
//					OrderName: "First Order 0",
//					CustomerName:  "Last Name 0",
//					Email:     "employee@email.com",
//				}
//				expectedReturn := []entity.Order{employee}
//				r.On("FindAllOrders", context.Background(), model.GetOrdersFilter{}).Return(expectedReturn, nil)
//				return NewOrderService(r)
//			},
//			Context:        createEchoContext(false),
//			RequestParam:   model.GetOrdersFilter{},
//			ExpectedCount:  len(*getExpectedOrdersResult()),
//			ExpectedResult: getExpectedOrdersResult(),
//			ExpectedError:  pkgerror.NoError,
//		},
//	}
//	for _, tc := range testCases {
//		t.Run(tc.Name, func(t *testing.T) {
//			r := new(mocks.Repository)
//			s := tc.InitService(r)
//			results, err := s.GetOrders(tc.Context, tc.RequestParam)
//			assert.Equal(t, tc.ExpectedResult, results)
//			assert.Equal(t, tc.ExpectedError.Code, err.Code)
//			assert.Equal(t, tc.ExpectedError.HttpCode, err.HttpCode)
//			assert.Equal(t, tc.ExpectedError.Msg, err.Msg)
//			if tc.ExpectedResult != nil {
//				assert.NotNil(t, results)
//				assert.Equal(t, tc.ExpectedResult, results)
//			}
//			r.AssertExpectations(t)
//		})
//	}
//}
//
//func TestGetOrderByID(t *testing.T) {
//	employee := entity.Order{
//		ID: 1,
//	}
//	result := model.GetOrderByIDResult{}
//	copyutil.Copy(&employee, &result)
//	testCases := []struct {
//		Name           string
//		InitService    func(r *mocks.Repository) OrderService
//		Context        echo.Context
//		ExpectedResult *model.GetOrderByIDResult
//		ExpectedError  pkgerror.CustomError
//	}{
//		{
//			Name: "SystemError",
//			InitService: func(r *mocks.Repository) OrderService {
//				r.On("FindOrderByID", mock.Anything, mock.Anything).Return(entity.Order{}, errors.New("database error"))
//				return NewOrderService(r)
//			},
//			Context:       createEchoContext(true),
//			ExpectedError: pkgerror.ErrSystemError,
//		},
//		{
//			Name: "OrderNotFound",
//			InitService: func(r *mocks.Repository) OrderService {
//				r.On("FindOrderByID", mock.Anything, mock.Anything).Return(entity.Order{}, gorm.ErrRecordNotFound)
//				return NewOrderService(r)
//			},
//			Context:       createEchoContext(true),
//			ExpectedError: pkgerror.ErrOrderNotFound,
//		},
//		{
//			Name: "Success",
//			InitService: func(r *mocks.Repository) OrderService {
//				r.On("FindOrderByID", mock.Anything, mock.Anything).Return(employee, nil)
//				return NewOrderService(r)
//			},
//			Context:        createEchoContext(true),
//			ExpectedError:  pkgerror.NoError,
//			ExpectedResult: &result,
//		},
//	}
//	for _, tc := range testCases {
//		t.Run(tc.Name, func(t *testing.T) {
//			r := new(mocks.Repository)
//			s := tc.InitService(r)
//			result, err := s.GetOrderByID(tc.Context, model.GetOrderByIDRequest{})
//			assert.Equal(t, tc.ExpectedError.Code, err.Code)
//			assert.Equal(t, tc.ExpectedError.HttpCode, err.HttpCode)
//			assert.Equal(t, tc.ExpectedError.Msg, err.Msg)
//			if tc.ExpectedResult != nil {
//				assert.NotNil(t, result)
//				assert.Equal(t, tc.ExpectedResult, result)
//			}
//			r.AssertExpectations(t)
//		})
//	}
//}
//
//func TestCreateOrder(t *testing.T) {
//	testCases := []struct {
//		Name           string
//		InitService    func(r *mocks.Repository) OrderService
//		Context        echo.Context
//		Request        model.CreateOrderRequest
//		ExpectedResult *model.CreateOrderResult
//		ExpectedError  pkgerror.CustomError
//	}{
//		{
//			Name: "FindOrderByEmailErrorSystem",
//			InitService: func(r *mocks.Repository) OrderService {
//				r.On("FindOrderByName", context.Background(), "employee@email.com").Return(entity.Order{}, errors.New("database error"))
//				return NewOrderService(r)
//			},
//			Context: createEchoContext(false),
//			Request: model.CreateOrderRequest{
//				Email: "employee@email.com",
//			},
//			ExpectedResult: nil,
//			ExpectedError:  pkgerror.ErrSystemError,
//		},
//		{
//			Name: "FindOrderByEmailErrorExisted",
//			InitService: func(r *mocks.Repository) OrderService {
//				r.On("FindOrderByName", context.Background(), "employee@email.com").Return(entity.Order{ID: uint(1), Email: "employee@email.com"}, nil)
//				return NewOrderService(r)
//			},
//			Context: createEchoContext(false),
//			Request: model.CreateOrderRequest{
//				Email: "employee@email.com",
//			},
//			ExpectedResult: nil,
//			ExpectedError:  pkgerror.ErrOrderIsExist,
//		},
//		{
//			Name: "CreateOrderError",
//			InitService: func(r *mocks.Repository) OrderService {
//				r.On("FindOrderByName", context.Background(), "employee@email.com").Return(entity.Order{}, nil)
//				r.On("TxBegin").Return(nil)
//				r.On("CreateOrder", context.Background(), mock.Anything).Return(errors.New("database error"))
//				r.On("TxRollback").Return(nil)
//				return NewOrderService(r)
//			},
//			Context: createEchoContext(true),
//			Request: model.CreateOrderRequest{
//				OrderName: "First Order 0",
//				CustomerName:  "Last Name 0",
//				Email:     "employee@email.com",
//				HireDate:  "2023-09-20",
//			},
//			ExpectedResult: nil,
//			ExpectedError:  pkgerror.ErrSystemError,
//		},
//		{
//			Name: "CreateOrderSuccess",
//			InitService: func(r *mocks.Repository) OrderService {
//				r.On("FindOrderByName", context.Background(), "employee@email.com").Return(entity.Order{}, nil)
//				r.On("TxBegin").Return(nil)
//				r.On("CreateOrder", context.Background(), mock.Anything).Return(nil)
//				r.On("TxCommit").Return(nil)
//				return NewOrderService(r)
//			},
//			Context: createEchoContext(true),
//			Request: model.CreateOrderRequest{
//				OrderName: "First Order 0",
//				CustomerName:  "Last Name 0",
//				Email:     "employee@email.com",
//			},
//			ExpectedResult: &model.CreateOrderResult{
//				OrderName: "First Order 0",
//				CustomerName:  "Last Name 0",
//				Email:     "employee@email.com",
//			},
//			ExpectedError: pkgerror.NoError,
//		},
//	}
//	for _, tc := range testCases {
//		t.Run(tc.Name, func(t *testing.T) {
//			r := new(mocks.Repository)
//			s := tc.InitService(r)
//			result, err := s.CreateOrder(tc.Context, tc.Request)
//			assert.Equal(t, tc.ExpectedError.Code, err.Code)
//			assert.Equal(t, tc.ExpectedError.HttpCode, err.HttpCode)
//			assert.Equal(t, tc.ExpectedError.Msg, err.Msg)
//			if tc.ExpectedResult != nil {
//				assert.NotNil(t, result)
//				assert.Equal(t, tc.ExpectedResult, result)
//			}
//			r.AssertExpectations(t)
//		})
//	}
//}
