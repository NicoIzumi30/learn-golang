package resto

import (
	"github.com/NicoIzumi30/go-restaurant-app/internal/model"
	"github.com/NicoIzumi30/go-restaurant-app/internal/model/constant"
	"github.com/NicoIzumi30/go-restaurant-app/internal/repository/menu"
	"github.com/NicoIzumi30/go-restaurant-app/internal/repository/order"
	"github.com/google/uuid"
)

type restoUsecase struct {
	menuRepo  menu.Repository
	orderRepo order.Repository
}

func GetUsecase(menuRepo menu.Repository, orderRepo order.Repository) Usecase {
	return &restoUsecase{
		menuRepo:  menuRepo,
		orderRepo: orderRepo,
	}
}

func (u *restoUsecase) GetMenuList(menuType string) ([]model.MenuItem, error) {
	return u.menuRepo.GetMenuList(menuType)
}

func (r *restoUsecase) Order(request model.OrderMenuRequest) (model.Order, error) {
	productOrderData := make([]model.ProductOrder, len(request.OrderProducts))

	for i, orderProduct := range request.OrderProducts {
		menuData, err := r.menuRepo.GetMenu(orderProduct.OrderCode)
		if err != nil {
			return model.Order{}, err
		}

		productOrderData[i] = model.ProductOrder{
			ID:         uuid.New().String(),
			OrderCode:  menuData.OrderCode,
			Quantity:   orderProduct.Quantity,
			TotalPrice: int64(menuData.Price * orderProduct.Quantity),
			Status:     constant.ProductOrderStatusPreparing,
		}
	}
	orderData := model.Order{
		ID:            uuid.New().String(),
		Status:        constant.OrderStatusProcessed,
		ProductOrders: productOrderData,
		ReferenceID:   request.ReferenceID,
	}
	createdOrderData, err := r.orderRepo.CreateOrder(orderData)
	if err != nil {
		return model.Order{}, err
	}
	return createdOrderData, nil
}

func (r *restoUsecase) GetOrderInfo(request model.GetOrderInfoRequest) (model.Order, error) {
	return r.orderRepo.GetOrderInfo(request.OrderID)
}
