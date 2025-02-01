package order

import "github.com/NicoIzumi30/go-restaurant-app/internal/model"

type Repository interface {
	CreateOrder(order model.Order) (model.Order, error)
	GetOrderInfo(orderID string) (model.Order, error)
}
