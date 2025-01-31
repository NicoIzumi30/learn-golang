package resto

import "github.com/NicoIzumi30/go-restaurant-app/internal/model"

type Usecase interface {
	GetMenu(menuType string) ([]model.MenuItem, error)
}
