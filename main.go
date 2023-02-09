package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main()  {
	seedDB()
	e := echo.New()
	e.GET("/menu/food", getFoodMenu)
	e.GET("/menu/drink", getDrinkMenu)
	e.Logger.Fatal(e.Start(":8080"))
}

type MenuItem struct {
	Name		string
	OrderCode 	string
	Price		int64
}


func seedDB()  {
	foodMenu := []MenuItem{
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

	drinkMenu := []MenuItem{
		{
			Name: "Bobba Gum",
			OrderCode: "Bukan Soda",
			Price: 1020004,
		},

		{
			Name: "Orang Tua",
			OrderCode: "Alkohol",
			Price: 90000,
		},
	}

	fmt.Println(foodMenu,drinkMenu)

	dbAddress := "host=localhost port=5432 user=iqbalp12 password=mbangg12 dbname=go_resto_sim sslmode=disable"
	db, err := gorm.Open(postgres.Open(dbAddress))
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&MenuItem{})
}

func getFoodMenu(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		// "data": foodMenu,
	})
}

func getDrinkMenu(c echo.Context) error{
	return c.JSON(http.StatusOK, map[string]interface{}{
		// "data": drinkMenu,
	})
}