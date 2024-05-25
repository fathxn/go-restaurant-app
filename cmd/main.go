package main

import (
	"github.com/labstack/echo/v4"
	"go-restaurant-app/internal/database"
	"go-restaurant-app/internal/delivery/rest"
	"go-restaurant-app/internal/repository/menu"
	"go-restaurant-app/internal/repository/order"
	"go-restaurant-app/internal/repository/user"
	"go-restaurant-app/internal/usecase/restaurant"
)

const (
	dsn = "host=127.0.0.1 port=5432 user=postgres password=root dbname=go-restaurant-app sslmode=disable"
)

func main() {
	e := echo.New()

	db := database.GetDB(dsn)
	secret := "AES256Key-32Characters1234567890"

	menuRepo := menu.GetRepository(db)
	orderRepo := order.GetRepository(db)
	userRepo, err := user.GetRepository(db, secret, 1, 64*1024, 4, 32)
	if err != nil {
		panic(err)
	}

	restaurantUsecase := restaurant.GetUsecase(menuRepo, orderRepo, userRepo)

	h := rest.NewHandler(restaurantUsecase)

	rest.LoadMiddlware(e)
	rest.LoadRoutes(e, h)

	e.Logger.Fatal(e.Start(":8080"))
}
