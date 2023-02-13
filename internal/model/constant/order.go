package constant

import (
	"github.com/iqbalpradipta/restaurant-sim/internal/model"
)

const (
	OrderStatusProcess		model.OrderStatus= "process"
	OrderStatusFinished		model.OrderStatus = "finished"
	OrderStatusFailed		model.OrderStatus = "failed"
)


const (
	ProductOrderStatusPreparing	model.ProductOrderStatus = "preparing"
	ProductOrderStatusFinishing model.ProductOrderStatus = "finishing"
)