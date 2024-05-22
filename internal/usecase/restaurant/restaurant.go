package restaurant

import (
	"go-restaurant-app/internal/model"
	"go-restaurant-app/internal/repository/menu"
)

type restaurantUsecase struct {
	menuRepo menu.Repository
}

func GetUsecase(menuRepo menu.Repository) Usecase {
	return &restaurantUsecase{menuRepo: menuRepo}
}

func (r *restaurantUsecase) GetMenu(menuType string) ([]model.MenuItem, error) {
	return r.menuRepo.GetMenu(menuType)
}
