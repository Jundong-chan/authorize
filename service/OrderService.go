package service

import (
	"CommoditySpike/server/model"

	"github.com/gohouse/gorose"
)

type OrderService interface {
	ModifyOrder(data map[string]interface{}, field string, value interface{}) error
	QureyOrder(value interface{}) ([]gorose.Data, error)
}
type OrderServiceMidWare func(OrderService) OrderService

type OrderServiceImpl struct {
}

func NewOrderSvcimpl() *OrderServiceImpl {
	return &OrderServiceImpl{}
}

// func (osvc OrderServiceImpl) CreateOrder(order model.Order) error {
// 	omodel := model.NewOrderModel()
// 	err := omodel.CreateOrder(order)
// 	return err
// }

func (osvc OrderServiceImpl) QureyOrder(value interface{}) ([]gorose.Data, error) {
	omodel := model.NewOrderModel()
	data, err := omodel.QureyOrder(value)
	return data, err
}

func (osvc OrderServiceImpl) ModifyOrder(data map[string]interface{}, field string, value interface{}) error {
	omodel := model.NewOrderModel()
	err := omodel.ModifyOrderByfield(data, field, value)
	return err
}
