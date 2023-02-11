package resto

import "github.com/iqbalpradipta/restaurant-sim/internal/model"

type Usecase interface {
	GetMenu(menuType string) ([]model.MenuItem, error)
}