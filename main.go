package main

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type MenuType string

const (
	Food  MenuType = "food"
	Drink MenuType = "drink"
)

const (
	dbAddress = "host=localhost port=5432 user=postgres password=mbangg12 dbname=go_resto_sim sslmode=disable"
)

type MenuItem struct {
	OrderCode string `gorm:"primaryKey"`
	Name      string
	Price     int64
	Type      MenuType
}

func seedDB() {
	dbAddress := "host=localhost port=5432 user=postgres password=mbangg12 dbname=go_resto_sim sslmode=disable"
	db, err := gorm.Open(postgres.Open(dbAddress), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&MenuItem{})

	foodMenu := []MenuItem{
		{
			Name: "Rendang",
			OrderCode: "Padang",
			Price: 1000000,
			Type: Food,
		},

		{
			Name: "AyamBakar",
			OrderCode: "Warteg",
			Price: 100000,
			Type: Food,
		},
	}

	drinkMenu := []MenuItem{
		{
			Name: "Bobba Gum",
			OrderCode: "Bukan Soda",
			Price: 1020004,
			Type: Drink,
		},

		{
			Name: "Orang Tua",
			OrderCode: "Alkohol",
			Price: 90000,
			Type: Drink,
		},
	}

	if err := db.First(&MenuItem{}).Error; err == gorm.ErrRecordNotFound {
		fmt.Println("Seeding db data...")
		db.Create(&foodMenu)
		db.Create(&drinkMenu)
	}
}

func main() {
	e := echo.New()

	seedDB()

	e.GET("/menu", GetMenu)

	e.Logger.Fatal(e.Start((":14045")))
}

func GetMenu(c echo.Context) error {
	menuType := c.FormValue("menu_type")

	db, err := gorm.Open(postgres.Open(dbAddress), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	menuData := make([]MenuItem, 0)

	if err := db.Where(MenuItem{Type: MenuType(menuType)}).Find(&menuData).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": menuData,
	})
}