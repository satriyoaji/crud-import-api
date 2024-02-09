package repository

import (
	"context"
	"fullstack_api_test/constant"
	"fullstack_api_test/entity"
	"fullstack_api_test/model"

	"gorm.io/gorm"
)

func whereOrderNameContains(name string, alias string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if name == "" {
			return db
		}
		sql := withAlias(string(constant.OrderColumnOrderName), alias) + " ilike ?"
		return db.Where(sql, withPercentAround(name))
	}
}
func whereOrderCustomerNameContains(name string, alias string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if name == "" {
			return db
		}
		sql := withAlias(string(constant.OrderColumnCustomerName), alias) + " ilike ?"
		return db.Where(sql, withPercentAround(name))
	}
}

func whereOrderIDIn(ids []int, alias string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if len(ids) == 0 {
			return db
		}
		sql := withAlias("id", alias) + " in ?"
		return db.Where(sql, ids)
	}
}

func (d DefaultRepository) FindOrders(ctx context.Context, filter model.GetOrdersFilter) ([]entity.Order, error) {
	employeeIDs := []int{}
	if filter.ID != nil {
		employeeIDs = append(employeeIDs, *filter.ID)
	}
	shops := []entity.Order{}
	err := d.handler.Tx.WithContext(ctx).
		Scopes(
			whereOrderNameContains(filter.OrderName, ""),
			whereOrderCustomerNameContains(filter.CustomerName, ""),
			whereOrderIDIn(employeeIDs, ""),
			paginate(filter.PageRequest.PageNum, filter.PageRequest.PageSize)).
		Order("order_date desc").Find(&shops).Error
	return shops, err
}

func (d DefaultRepository) FindAllOrders(ctx context.Context, filter model.GetOrdersFilter) ([]entity.Order, error) {
	employeeIDs := []int{}
	if filter.ID != nil {
		employeeIDs = append(employeeIDs, *filter.ID)
	}
	shops := []entity.Order{}
	err := d.handler.Tx.WithContext(ctx).
		Scopes(
			whereOrderNameContains(filter.OrderName, ""),
			whereOrderCustomerNameContains(filter.CustomerName, ""),
			whereOrderIDIn(employeeIDs, "")).
		Order("order_date desc").Find(&shops).Error
	return shops, err
}

func (d DefaultRepository) CountOrders(ctx context.Context, filter model.GetOrdersFilter) (int, error) {
	employeeIDs := []int{}
	if filter.ID != nil {
		employeeIDs = append(employeeIDs, *filter.ID)
	}
	var count int64
	err := d.handler.Tx.WithContext(ctx).Model(&entity.Order{}).
		Scopes(
			whereOrderNameContains(filter.OrderName, ""),
			whereOrderCustomerNameContains(filter.CustomerName, ""),
			whereOrderIDIn(employeeIDs, "")).
		Count(&count).Error
	return int(count), err
}

func (d DefaultRepository) CreateOrder(ctx context.Context, employee *entity.Order) error {
	return d.handler.Tx.Create(employee).Error
}

func (d DefaultRepository) FindOrderByID(ctx context.Context, id uint) (entity.Order, error) {
	employee := entity.Order{}
	err := d.handler.Tx.WithContext(ctx).Where("id=?", id).First(&employee).Error
	return employee, err
}

func (d DefaultRepository) FindOrderByName(ctx context.Context, name string) (entity.Order, error) {
	employee := entity.Order{}
	err := d.handler.Tx.WithContext(ctx).Where("order_name=?", name).First(&employee).Error
	return employee, err
}

func (d DefaultRepository) UpdateOrder(ctx context.Context, employee *entity.Order) error {
	return d.handler.Tx.WithContext(ctx).Save(employee).Error
}

func (d DefaultRepository) DeleteOrder(ctx context.Context, id uint) error {
	return d.handler.Tx.WithContext(ctx).Delete(&entity.Order{}, id).Error
}
