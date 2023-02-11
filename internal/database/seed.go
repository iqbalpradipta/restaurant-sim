package database

import (
	"github.com/iqbalpradipta/restaurant-sim/internal/model"
	"github.com/iqbalpradipta/restaurant-sim/internal/model/constant"
	"gorm.io/gorm"
)
func seedDB(db *gorm.DB)  {
	foodMenu := []model.MenuItem{
		{
			Name: "Rendang",
			OrderCode: "Padang",
			Price: 1000000,
			Type: model.MenuType(constant.Food),
		},

		{
			Name: "AyamBakar",
			OrderCode: "Warteg",
			Price: 100000,
			Type: model.MenuType(constant.Food),
		},
	}

	drinkMenu := []model.MenuItem{
		{
			Name: "Bobba Gum",
			OrderCode: "Bukan Soda",
			Price: 1020004,
			Type: model.MenuType(constant.Drink),
		},

		{
			Name: "Orang Tua",
			OrderCode: "Alkohol",
			Price: 90000,
			Type: model.MenuType(constant.Drink),
		},
	}

	if err := db.First(&model.MenuItem{}).Error; err == gorm.ErrRecordNotFound {
		db.Create(&foodMenu)
		db.Create(&drinkMenu)
	}
}