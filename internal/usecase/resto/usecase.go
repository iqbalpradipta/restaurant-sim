package resto

import "github.com/iqbalpradipta/restaurant-sim/internal/model"

type Usecase interface {
	GetMenuLList(menuType string) ([]model.MenuItem, error)
	Order(request model.OrderMenuRequest) (model.Order, error)
	GetOrderInfo(request model.GetOrderInfoRequest) (model.Order, error)
}