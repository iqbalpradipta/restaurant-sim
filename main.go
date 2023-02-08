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

type DrinkMenu struct {
	Nama		string
	TypeDrink	string
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

func getDrinkMenu(c echo.Context) error{
	drinkMenu := []DrinkMenu{
		{
			Nama: "Bobba Gum",
			TypeDrink: "Bukan Soda",
			Price: 1020004,
		},

		{
			Nama: "Orang Tua",
			TypeDrink: "Alkohol",
			Price: 90000,
		},
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": drinkMenu,
	})
}

func main()  {
	e := echo.New()
	e.GET("/menu/food", getFoodMenu)
	e.GET("/menu/drink", getDrinkMenu)
	e.Logger.Fatal(e.Start(":8080"))
}