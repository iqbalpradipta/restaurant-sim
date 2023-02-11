package main

import (
	"github.com/iqbalpradipta/restaurant-sim/internal/database"
	"github.com/iqbalpradipta/restaurant-sim/internal/delivery/rest"
	repomenu "github.com/iqbalpradipta/restaurant-sim/internal/repository/repoMenu"
	restoUsecase "github.com/iqbalpradipta/restaurant-sim/internal/usecase/resto"
	"github.com/labstack/echo/v4"
)

const (
	dbAddress = "host=localhost port=5432 user=postgres password=mbangg12 dbname=go_resto_sim sslmode=disable"
)

func main() {
	e := echo.New()
	db := database.GetDB(dbAddress)
	menuRepo := repomenu.GetRepository(db)
	restoUsecase := restoUsecase.GetUseCase(menuRepo)

	h := rest.NewHandler(restoUsecase)
	rest.LoadRoutes(e, h)
	e.Logger.Fatal(e.Start((":14045")))
}

