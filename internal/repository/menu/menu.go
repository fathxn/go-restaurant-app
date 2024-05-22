package menu

import (
	"go-restaurant-app/internal/model"
	"gorm.io/gorm"
)

type menuRepo struct {
	db *gorm.DB
}

func GetRepository(db *gorm.DB) Repository {
	return &menuRepo{db: db}
}

func (m menuRepo) GetMenu(menuType string) ([]model.MenuItem, error) {
	var menuData []model.MenuItem
	//err := m.db.Where("type = ?", menuType).Find(&menuData).Error
	if err := m.db.Where(model.MenuItem{Type: model.MenuType(menuType)}).Find(&menuData).Error; err != nil {
		return nil, err
	}
	return menuData, nil
}
