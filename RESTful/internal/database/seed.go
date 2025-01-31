package database

import (
	"github.com/NicoIzumi30/go-restaurant-app/internal/model"
	"github.com/NicoIzumi30/go-restaurant-app/internal/model/constant"
	"gorm.io/gorm"
)

func seedDB(db *gorm.DB) {
	foodMenu := []model.MenuItem{
		{
			Name:      "Nasi Goreng",
			OrderCode: "nasi-goreng",
			Price:     20000,
			Type:      constant.MenuTypeFood,
		},
		{
			Name:      "Ayam Goreng",
			OrderCode: "ayam-goreng",
			Price:     15000,
			Type:      constant.MenuTypeFood,
		},
		{
			Name:      "Mie Goreng",
			OrderCode: "mie-goreng",
			Price:     10000,
			Type:      constant.MenuTypeFood,
		},
	}
	drinkMenu := []model.MenuItem{
		{
			Name:      "Es Teh",
			OrderCode: "es-teh",
			Price:     5000,
			Type:      constant.MenuTypeDrink,
		},
		{
			Name:      "Es Jeruk",
			OrderCode: "es-jeruk",
			Price:     8000,
			Type:      constant.MenuTypeDrink,
		},
		{
			Name:      "Es Kelapa",
			OrderCode: "es-kelapa",
			Price:     10000,
			Type:      constant.MenuTypeDrink,
		},
	}
	if err := db.First(&model.MenuItem{}).Error; err == gorm.ErrRecordNotFound {
		db.Create(&foodMenu)
		db.Create(&drinkMenu)
	}
}
