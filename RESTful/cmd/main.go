package main

import (
	"github.com/NicoIzumi30/go-restaurant-app/internal/database"
	"github.com/NicoIzumi30/go-restaurant-app/internal/delivery/rest"
	mRepo "github.com/NicoIzumi30/go-restaurant-app/internal/repository/menu"
	rUsecase "github.com/NicoIzumi30/go-restaurant-app/internal/usecase/resto"
	"github.com/labstack/echo/v4"
)

const (
	dbAddress = "host=localhost user=postgres password=postgres dbname=go_resto_app port=5432 sslmode=disable"
)

func main() {

	e := echo.New()
	db := database.GetDB(dbAddress)
	menuRepo := mRepo.GetRepository(db)
	restoUsecase := rUsecase.GetUsecase(menuRepo)

	h := rest.NewHandler(restoUsecase)
	rest.LoadRoutes(e, h)
	e.Logger.Fatal(e.Start(":3000"))
}
