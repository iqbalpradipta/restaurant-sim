package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type FoodMenu struct {
	Name		string
	OrderCode 	string
	Price		int
}

func getFoodMenu(c echo.Context) error {
	foodMenu := []FoodMenu{
		{
			Name: "Rendang",
			OrderCode: "Padang",
			Price: 1000000,
		},

		{
			Name: "AyamBakar",
			OrderCode: "Warteg",
			Price: 100000,
		},
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": foodMenu,
	})
}

func main()  {
	e := echo.New()
	e.GET("/menu/food", getFoodMenu)
	e.Logger.Fatal(e.Start(":8080"))
}