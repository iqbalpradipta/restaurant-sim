package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	dbAddress = "host=localhost port=5432 user=postgres password=mbangg12 dbname=go_resto_sim sslmode=disable"
)

func main()  {
	seedDB()
	e := echo.New()
	e.GET("/menu/food", getFoodMenu)
	e.GET("/menu/drink", getDrinkMenu)
	e.Logger.Fatal(e.Start(":8080"))
}

type MenuType string

const (
	MenuTypeFood = "food"
	MenuTypeDrink = "drink"
)

type MenuItem struct {
	Name		string
	OrderCode 	string
	Price		int64
	Type		MenuType
}


func seedDB()  {
	foodMenu := []MenuItem{
		{
			Name: "Rendang",
			OrderCode: "Padang",
			Price: 1000000,
			Type: MenuTypeFood,
		},

		{
			Name: "AyamBakar",
			OrderCode: "Warteg",
			Price: 100000,
			Type: MenuTypeFood,
		},
	}

	drinkMenu := []MenuItem{
		{
			Name: "Bobba Gum",
			OrderCode: "Bukan Soda",
			Price: 1020004,
			Type: MenuTypeDrink,
		},

		{
			Name: "Orang Tua",
			OrderCode: "Alkohol",
			Price: 90000,
			Type: MenuTypeDrink,
		},
	}

	db, err := gorm.Open(postgres.Open(dbAddress))
	if err != nil {
		panic(err)
	}

	if err := db.First(&MenuItem{}).Error; err == gorm.ErrRecordNotFound {
		db.Create(&foodMenu)
		db.Create(&drinkMenu)
	}
}

func getFoodMenu(c echo.Context) error {
	db, err := gorm.Open(postgres.Open(dbAddress))
	if err != nil{
		panic(err)
	}

	var menuData []MenuItem
	db.Where(MenuItem{Type: MenuTypeFood}).Find(&menuData)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": menuData,
	})
}

func getDrinkMenu(c echo.Context) error{
	db, err := gorm.Open(postgres.Open(dbAddress))
	if err != nil {
		panic(err)
	}

	var menuData []MenuItem
	db.Where(MenuItem{Type: MenuTypeDrink}).Find(&menuData)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": menuData,
	})
}