package main

import (
	"crypto/rand"
	"crypto/rsa"
	"github.com/labstack/echo/v4"
	"go-restaurant-app/internal/database"
	"go-restaurant-app/internal/delivery/rest"
	"go-restaurant-app/internal/logger"
	"go-restaurant-app/internal/repository/menu"
	"go-restaurant-app/internal/repository/order"
	"go-restaurant-app/internal/repository/user"
	"go-restaurant-app/internal/tracing"
	"go-restaurant-app/internal/usecase/restaurant"
	"time"
)

const (
	dsn = "host=127.0.0.1 port=5432 user=postgres password=root dbname=go-restaurant-app sslmode=disable"
)

func main() {
	logger.Init()
	tracing.Init("http://localhost:14268/api/traces")
	e := echo.New()

	db := database.GetDB(dsn)
	secret := "AES256Key-32Characters1234567890"
	signKey, err := rsa.GenerateKey(rand.Reader, 4096)
	if err != nil {
		panic(err)
	}

	menuRepo := menu.GetRepository(db)
	orderRepo := order.GetRepository(db)
	userRepo, err := user.GetRepository(db, secret, 1, 64*1024, 4, 32, 60*time.Second, signKey)
	if err != nil {
		panic(err)
	}

	restaurantUsecase := restaurant.GetUsecase(menuRepo, orderRepo, userRepo)

	h := rest.NewHandler(restaurantUsecase)

	rest.LoadMiddlware(e)
	rest.LoadRoutes(e, h)

	e.Logger.Fatal(e.Start(":8080"))
}
