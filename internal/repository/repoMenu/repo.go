package repomenu

import (
	"github.com/iqbalpradipta/restaurant-sim/internal/model"
)

type Repository interface {
	GetMenuList(menuType string) ([]model.MenuItem, error)
	GetMenu(orderCode string) (model.MenuItem, error)
}