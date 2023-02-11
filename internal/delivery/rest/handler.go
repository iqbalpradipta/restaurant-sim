package rest

import "github.com/iqbalpradipta/restaurant-sim/internal/usecase/resto"

type handler struct {
	restoUsecase resto.Usecase
}

func NewHandler(restoUsecase resto.Usecase) *handler  {
	return &handler{
		restoUsecase: restoUsecase,
	}
}