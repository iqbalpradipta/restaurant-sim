package resto

import (
	"github.com/iqbalpradipta/restaurant-sim/internal/model"
	"github.com/iqbalpradipta/restaurant-sim/internal/repository/repoMenu"
)
type restoUsecase struct {
	menuRepo repomenu.Repository
}

func GetUseCase(menuRepo repomenu.Repository) Usecase {
	return &restoUsecase{
		menuRepo : menuRepo,
	}
}

func (r *restoUsecase) GetMenu(menuType string) ([]model.MenuItem, error)  {
	return r.menuRepo.GetMenu(menuType)
}