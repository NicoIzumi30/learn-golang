package main

import (
	"github.com/NicoIzumi30/go-restaurant-app/internal/database"
	"github.com/NicoIzumi30/go-restaurant-app/internal/delivery/rest"
	mRepo "github.com/NicoIzumi30/go-restaurant-app/internal/repository/menu"
	oRepo "github.com/NicoIzumi30/go-restaurant-app/internal/repository/order"
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
	orderRepo := oRepo.GetRepository(db)
	restoUsecase := rUsecase.GetUsecase(menuRepo, orderRepo)
	h := rest.NewHandler(restoUsecase)
	rest.LoadMiddlewares(e)
	rest.LoadRoutes(e, h)
	e.Logger.Fatal(e.Start(":3000"))
}
