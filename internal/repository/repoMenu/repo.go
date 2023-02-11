package repomenu

import (
	"github.com/iqbalpradipta/restaurant-sim/internal/model"
)

type Repository interface {
	GetMenu(menuType string) ([]model.MenuItem, error)
}