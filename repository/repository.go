package repository

import (
	"context"
	"fullstack_api_test/entity"
	"fullstack_api_test/model"
	"fullstack_api_test/pkg/db"
	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
	"strings"
)

type Repository interface {
	TxBegin() error
	TxCommit() error
	TxRollback() error

	// Order
	FindOrders(ctx context.Context, filter model.GetOrdersFilter) ([]entity.Order, error)
	FindAllOrders(ctx context.Context, filter model.GetOrdersFilter) ([]entity.Order, error)
	CountOrders(ctx context.Context, filter model.GetOrdersFilter) (int, error)
	CreateOrder(ctx context.Context, merchant *entity.Order) error
	FindOrderByID(ctx context.Context, id uint) (entity.Order, error)
	FindOrderByName(ctx context.Context, name string) (entity.Order, error)
	UpdateOrder(ctx context.Context, merchant *entity.Order) error
	DeleteOrder(ctx context.Context, id uint) error
}

type DefaultRepository struct {
	handler *db.Handler
}

func Default(handler *db.Handler) *DefaultRepository {
	return &DefaultRepository{
		handler: handler,
	}
}

func (d DefaultRepository) TxBegin() error {
	log.Debug("Start db transaction")
	d.handler.Tx = d.handler.DB.Begin()
	return d.handler.Tx.Error
}

func (d DefaultRepository) TxCommit() error {
	log.Debug("Commit db transaction")
	err := d.handler.Tx.Commit().Error
	d.handler.Tx = d.handler.DB
	return err
}

func (d DefaultRepository) TxRollback() error {
	log.Debug("Rollback db transaction")
	err := d.handler.Tx.Rollback().Error
	d.handler.Tx = d.handler.DB
	return err
}

func paginate(pageNum, pageSize int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if pageNum <= 0 {
			pageNum = 1
		}
		if pageSize <= 0 {
			pageSize = 10
		}
		offset := (pageNum - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}

func withAlias(column, alias string) string {
	if alias == "" {
		return column
	}
	return alias + "." + column
}

func withPercentAround(val string) string {
	return "%" + val + "%"
}

func withPercentAfter(val string) string {
	return val + "%"
}

func withPercentBefore(val string) string {
	return "%" + val
}

func getSortDir(sortDir string) string {
	if strings.ToLower(sortDir) == "asc" {
		return strings.ToLower(sortDir)
	}
	return "desc"
}
