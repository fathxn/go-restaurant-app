package main

import (
	"github.com/labstack/echo/v4"
	"go-restaurant-app/internal/database"
	"go-restaurant-app/internal/delivery/rest"
	"go-restaurant-app/internal/repository/menu"
	"go-restaurant-app/internal/usecase/restaurant"
)

const (
	dsn = "host=127.0.0.1 port=5432 user=postgres password=root dbname=go-restaurant-app sslmode=disable"
)

func main() {
	e := echo.New()

	db := database.GetDB(dsn)
	menuRepo := menu.GetRepository(db)
	restaurantUsecase := restaurant.GetUsecase(menuRepo)
	h := rest.NewHandler(restaurantUsecase)
	rest.LoadRoutes(e, h)
	e.Logger.Fatal(e.Start(":8080"))
}
