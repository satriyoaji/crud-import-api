package service

import (
	"encoding/csv"
	"errors"
	"fullstack_api_test/entity"
	"fullstack_api_test/model"
	"fullstack_api_test/repository"
	"github.com/gocarina/gocsv"
	"github.com/labstack/echo/v4"
	"io"
	"strings"
	"time"

	pkgerror "fullstack_api_test/pkg/error"
	"fullstack_api_test/pkg/util/copyutil"

	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
)

type OrderService interface {
	GetOrders(ctx echo.Context, filter model.GetOrdersFilter) (*[]model.GetOrdersResult, pkgerror.CustomError)
	CreateOrder(ctx echo.Context, req model.CreateOrderRequest) (*model.CreateOrderResult, pkgerror.CustomError)
	BulkOrdersFromCSV(ctx echo.Context, reqFile model.CSVUploadInput) pkgerror.CustomError
	GetOrderByID(ctx echo.Context, req model.GetOrderByIDRequest) (*model.GetOrderByIDResult, pkgerror.CustomError)
	EditOrder(ctx echo.Context, req model.EditOrderRequest) (*model.EditOrderResult, pkgerror.CustomError)
	DeleteOrderByID(ctx echo.Context, req model.DeleteOrderByIDRequest) pkgerror.CustomError
}

type OrderServiceImpl struct {
	repo repository.Repository
}

func NewOrderService(
	repo repository.Repository) *OrderServiceImpl {
	return &OrderServiceImpl{
		repo: repo,
	}
}

func (s OrderServiceImpl) GetOrders(ctx echo.Context, filter model.GetOrdersFilter) (*[]model.GetOrdersResult, pkgerror.CustomError) {
	rctx := ctx.Request().Context()
	results := []model.GetOrdersResult{}
	orders, err := s.repo.FindAllOrders(rctx, filter)
	if err != nil {
		log.Error("Find orders error: ", err)
		return nil, pkgerror.ErrSystemError
	}
	copyutil.Copy(&orders, &results)
	return &results, pkgerror.NoError
}

func (s *OrderServiceImpl) GetOrderByID(ctx echo.Context, req model.GetOrderByIDRequest) (*model.GetOrderByIDResult, pkgerror.CustomError) {
	rctx := ctx.Request().Context()
	order, err := s.repo.FindOrderByID(rctx, uint(req.OrderID))
	if err != nil {
		log.Error("Find order by ID error: ", err)
		if errors.Is(gorm.ErrRecordNotFound, err) {
			return nil, pkgerror.ErrOrderNotFound.WithError(err)
		}
		return nil, pkgerror.ErrSystemError.WithError(err)
	}
	result := model.GetOrderByIDResult{}
	copyutil.Copy(&order, &result)
	return &result, pkgerror.NoError
}

func (s *OrderServiceImpl) BulkOrdersFromCSV(ctx echo.Context, reqFile model.CSVUploadInput) pkgerror.CustomError {
	rctx := ctx.Request().Context()

	//f, err := os.OpenFile("orders_example_test.csv", os.O_RDWR|os.O_CREATE, os.ModePerm)
	formFile, err := ctx.FormFile("file_csv")
	if err != nil {
		return pkgerror.ErrSystemError.WithError(err)
	}
	f, err := formFile.Open()
	if err != nil {
		return pkgerror.ErrSystemError.WithError(err)
	}
	defer f.Close()

	fileBytes, err := io.ReadAll(f)
	if err != nil {
		return pkgerror.ErrSystemError.WithError(err)
	}

	gocsv.SetCSVReader(func(in io.Reader) gocsv.CSVReader {
		r := csv.NewReader(in)
		r.LazyQuotes = true
		r.Comma = ';'
		return r
	})

	var orders []*model.BulkInsertOrderRequestCSV
	if err := gocsv.UnmarshalBytes(fileBytes, &orders); err != nil { // Load clients from file
		return pkgerror.ErrSystemError.WithError(err)
	}

	txSuccess := false
	err = s.repo.TxBegin()
	if err != nil {
		log.Error("Start db transaction error: ", err)
		return pkgerror.ErrSystemError.WithError(err)
	}
	defer func() {
		if r := recover(); r != nil || !txSuccess {
			err = s.repo.TxRollback()
			if err != nil {
				log.Error("Rollback db transaction error: ", err)
			}
		}
	}()

	for _, order := range orders {
		var orderEntity entity.Order
		copyutil.Copy(&order, &orderEntity)
		orderEntity.OrderName = strings.ToUpper(orderEntity.OrderName)
		orderEntity.OrderDate = time.Now()
		err = s.repo.CreateOrder(rctx, &orderEntity)
		if err != nil {
			log.Error("Create order error: ", err)
			return pkgerror.ErrSystemError.WithError(err)
		}
	}
	err = s.repo.TxCommit()
	txSuccess = true

	return pkgerror.NoError
}

func (s *OrderServiceImpl) CreateOrder(ctx echo.Context, req model.CreateOrderRequest) (*model.CreateOrderResult, pkgerror.CustomError) {
	rctx := ctx.Request().Context()

	orderFound, err := s.repo.FindOrderByName(rctx, strings.ToUpper(req.OrderName))
	if err != nil {
		log.Error("Find user by Email error: ", err)
		if !errors.Is(gorm.ErrRecordNotFound, err) {
			return nil, pkgerror.ErrSystemError.WithError(err)
		}
	}
	if orderFound.OrderName != "" {
		return nil, pkgerror.ErrOrderIsExist.WithError(errors.New("Order `name` is already created."))
	}

	txSuccess := false
	err = s.repo.TxBegin()
	if err != nil {
		log.Error("Start db transaction error: ", err)
		return nil, pkgerror.ErrSystemError.WithError(err)
	}
	defer func() {
		if r := recover(); r != nil || !txSuccess {
			err = s.repo.TxRollback()
			if err != nil {
				log.Error("Rollback db transaction error: ", err)
			}
		}
	}()

	var order entity.Order
	copyutil.Copy(&req, &order)
	order.OrderName = strings.ToUpper(order.OrderName)
	order.OrderDate = time.Now()
	err = s.repo.CreateOrder(rctx, &order)
	if err != nil {
		log.Error("Create order error: ", err)
		return nil, pkgerror.ErrSystemError.WithError(err)
	}
	err = s.repo.TxCommit()
	if err != nil {
		log.Error("Commit db transaction error: ", err)
	}
	var result model.CreateOrderResult
	copyutil.Copy(&order, &result)
	txSuccess = true
	return &result, pkgerror.NoError
}

func (s *OrderServiceImpl) EditOrder(ctx echo.Context, req model.EditOrderRequest) (*model.EditOrderResult, pkgerror.CustomError) {
	rctx := ctx.Request().Context()
	order, err := s.repo.FindOrderByID(rctx, uint(req.OrderID))
	if err != nil {
		log.Error("Find order by ID error: ", err)
		if errors.Is(gorm.ErrRecordNotFound, err) {
			return nil, pkgerror.ErrOrderNotFound.WithError(err)
		}
		return nil, pkgerror.ErrSystemError.WithError(err)
	}

	// validate unique name on other orders
	orderByName, err := s.repo.FindOrderByName(rctx, strings.ToUpper(req.OrderName))
	if err != nil {
		log.Error("Find user by Email error: ", err)
		if !errors.Is(gorm.ErrRecordNotFound, err) {
			return nil, pkgerror.ErrSystemError.WithError(err)
		}
	}
	if orderByName.OrderName != "" && order.ID != orderByName.ID {
		return nil, pkgerror.ErrOrderIsExist.WithError(errors.New("Order `name` is already created."))
	}

	txSuccess := false
	err = s.repo.TxBegin()
	if err != nil {
		log.Error("Start db transaction error: ", err)
		return nil, pkgerror.ErrSystemError.WithError(err)
	}
	defer func() {
		if r := recover(); r != nil || !txSuccess {
			err = s.repo.TxRollback()
			if err != nil {
				log.Error("Rollback db transaction error: ", err)
			}
		}
	}()

	copyutil.Copy(&req, &order)
	order.OrderName = strings.ToUpper(order.OrderName)
	order.OrderDate = time.Now()
	err = s.repo.UpdateOrder(rctx, &order)
	if err != nil {
		return nil, pkgerror.ErrSystemError.WithError(err)
	}
	// Commit transaction
	err = s.repo.TxCommit()
	if err != nil {
		log.Error("Commit db transaction error: ", err)
	}
	result := model.EditOrderResult{}
	copyutil.Copy(&order, &result)
	txSuccess = true
	return &result, pkgerror.NoError
}

func (s *OrderServiceImpl) DeleteOrderByID(ctx echo.Context, req model.DeleteOrderByIDRequest) pkgerror.CustomError {
	rctx := ctx.Request().Context()
	err := s.repo.DeleteOrder(rctx, uint(req.OrderID))
	if err != nil {
		log.Error("Delete order by ID error: ", err)
		if errors.Is(gorm.ErrRecordNotFound, err) {
			return pkgerror.ErrOrderNotFound.WithError(err)
		}
		return pkgerror.ErrSystemError.WithError(err)
	}

	return pkgerror.NoError
}
