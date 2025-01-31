package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	dbAddress = "host=localhost user=postgres password=postgres dbname=go_resto_app port=5432 sslmode=disable"
)

type MenuType string

const (
	MenuTypeFood  MenuType = "food"
	MenuTypeDrink MenuType = "drink"
)

type MenuItem struct {
	Name      string
	OrderCode string
	Price     int
	Type      MenuType
}

func seadDB() {
	foodMenu := []MenuItem{
		{
			Name:      "Nasi Goreng",
			OrderCode: "nasi-goreng",
			Price:     20000,
			Type:      MenuTypeFood,
		},
		{
			Name:      "Ayam Goreng",
			OrderCode: "ayam-goreng",
			Price:     15000,
			Type:      MenuTypeFood,
		},
		{
			Name:      "Mie Goreng",
			OrderCode: "mie-goreng",
			Price:     10000,
			Type:      MenuTypeFood,
		},
	}
	drinkMenu := []MenuItem{
		{
			Name:      "Es Teh",
			OrderCode: "es-teh",
			Price:     5000,
			Type:      MenuTypeDrink,
		},
		{
			Name:      "Es Jeruk",
			OrderCode: "es-jeruk",
			Price:     8000,
			Type:      MenuTypeDrink,
		},
		{
			Name:      "Es Kelapa",
			OrderCode: "es-kelapa",
			Price:     10000,
			Type:      MenuTypeDrink,
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
func getMenu(c echo.Context) error {
	menuType := c.FormValue("menu_type")

	db, err := gorm.Open(postgres.Open(dbAddress))
	if err != nil {
		panic(err)
	}
	var menuData []MenuItem
	db.Where(MenuItem{Type: MenuType(menuType)}).Find(&menuData)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"data":      menuData,
		"menu_type": menuType,
	})
}

func main() {

	e := echo.New()
	e.GET("/menu", getMenu)
	seadDB()
	e.Logger.Fatal(e.Start(":3000"))
}
