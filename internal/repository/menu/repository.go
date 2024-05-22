package menu

import "go-restaurant-app/internal/model"

type Repository interface {
	GetMenu(menuType string) ([]model.MenuItem, error)
}
