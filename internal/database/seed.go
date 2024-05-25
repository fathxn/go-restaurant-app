package database

import (
	"go-restaurant-app/internal/model"
	"go-restaurant-app/internal/model/constant"
	"gorm.io/gorm"
)

func seedDB(db *gorm.DB) {
	db.AutoMigrate(&model.MenuItem{}, &model.Order{}, &model.ProductOrder{}, &model.User{})

	foodMenu := []model.MenuItem{
		{Name: "Bakmie Jawa", OrderCode: "BKM", Price: 12000, Type: constant.MenuTypeFood},
		{Name: "Nasi Goreng", OrderCode: "NGR", Price: 10000, Type: constant.MenuTypeFood},
		{Name: "Capcay Goreng", OrderCode: "CCG", Price: 8000, Type: constant.MenuTypeFood},
	}
	drinkMenu := []model.MenuItem{
		{Name: "Es Teh", OrderCode: "EST", Price: 3000, Type: constant.MenuTypeDrink},
		{Name: "Es Jeruk", OrderCode: "ESJ", Price: 5000, Type: constant.MenuTypeDrink},
		{Name: "Jus Melon", OrderCode: "JSM", Price: 6000, Type: constant.MenuTypeDrink},
	}

	if err := db.First(&model.MenuItem{}).Error; err == gorm.ErrRecordNotFound {
		db.Create(&foodMenu)
		db.Create(&drinkMenu)
	}
}
