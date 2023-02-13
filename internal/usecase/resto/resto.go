package resto

import (
	"github.com/google/uuid"
	"github.com/iqbalpradipta/restaurant-sim/internal/model"
	"github.com/iqbalpradipta/restaurant-sim/internal/model/constant"
	"github.com/iqbalpradipta/restaurant-sim/internal/repository/order"
	repomenu "github.com/iqbalpradipta/restaurant-sim/internal/repository/repoMenu"
)

type restoUsecase struct {
	menuRepo  repomenu.Repository
	orderRepo order.Repository
}

func GetUseCase(menuRepo repomenu.Repository, orderRepo order.Repository) Usecase {
	return &restoUsecase{
		menuRepo:  menuRepo,
		orderRepo: orderRepo,
	}
}

func (r *restoUsecase) GetMenuLList(menuType string) ([]model.MenuItem, error) {
	return r.menuRepo.GetMenuList(menuType)
}

func (r *restoUsecase) Order(request model.OrderMenuRequest) (model.Order, error) {
	productOrderData := make([]model.ProductOrders, len(request.OrderProducts))

	for i, orderProduct := range request.OrderProducts {
		menuData, err := r.menuRepo.GetMenu(orderProduct.OrderCode)
		if err != nil {
			return model.Order{}, err
		}
		productOrderData[i] = model.ProductOrders{
			ID:         uuid.New().String(),
			OrderCode:  menuData.OrderCode,
			Quantity:   orderProduct.Quantity,
			TotalPrice: menuData.Price * int64(orderProduct.Quantity),
			Status:     constant.ProductOrderStatusPreparing,
		}
	}

	orderData := model.Order{
		ID:            uuid.New().String(),
		Status:        constant.OrderStatusProcess,
		ProductOrders: productOrderData,
	}

	createOrderData, err := r.orderRepo.CreateOrder(orderData)
	if err != nil {
		return model.Order{}, err
	}
	return createOrderData, nil
}

func (r *restoUsecase) GetOrderInfo(reqeust model.GetOrderInfoRequest) (model.Order, error) {
	orderData, err := r.orderRepo.GetOrderInfo(reqeust.OrderID)
	if err != nil {
		return orderData, err
	}

	return orderData, nil
}
