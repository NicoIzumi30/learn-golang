package resto

import (
	"github.com/NicoIzumi30/go-restaurant-app/internal/model"
	"github.com/NicoIzumi30/go-restaurant-app/internal/repository/menu"
)

type restoUsecase struct {
	menuRepo menu.Repository
}

func GetUsecase(menuRepo menu.Repository) Usecase {
	return &restoUsecase{
		menuRepo: menuRepo,
	}
}

func (u *restoUsecase) GetMenu(menuType string) ([]model.MenuItem, error) {
	return u.menuRepo.GetMenu(menuType)
}
