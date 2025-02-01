package menu

import "github.com/NicoIzumi30/go-restaurant-app/internal/model"

type Repository interface {
	GetMenuList(menuType string) ([]model.MenuItem, error)
	GetMenu(orderCode string) (model.MenuItem, error)
}
