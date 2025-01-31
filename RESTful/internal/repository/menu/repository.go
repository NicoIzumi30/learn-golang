package menu

import "github.com/NicoIzumi30/go-restaurant-app/internal/model"

type Repository interface {
	GetMenu(menuType string) ([]model.MenuItem, error)
}
