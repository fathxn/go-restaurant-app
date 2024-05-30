package rest

import (
	"context"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go-restaurant-app/internal/model/constant"
	"go-restaurant-app/internal/usecase/restaurant"
	"net/http"
)

func LoadMiddlware(e *echo.Echo) {
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"https://restoku.com"},
	}))
}

func GetAuthMiddleware(restaurantUsecase restaurant.Usecase) *authMiddleware {
	return &authMiddleware{
		restaurantUsecase: restaurantUsecase,
	}
}

type authMiddleware struct {
	restaurantUsecase restaurant.Usecase
}

func (am *authMiddleware) CheckAuth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		sessionData, err := GetSessionData(c.Request())
		if err != nil {
			return &echo.HTTPError{
				Code:     http.StatusUnauthorized,
				Message:  err.Error(),
				Internal: err,
			}
		}

		userID, err := am.restaurantUsecase.CheckSession(sessionData)
		if err != nil {
			return &echo.HTTPError{
				Code:     http.StatusUnauthorized,
				Message:  err.Error(),
				Internal: err,
			}
		}

		authContext := context.WithValue(c.Request().Context(), constant.AuthContextKey, userID)
		c.SetRequest(c.Request().WithContext(authContext))

		return nil
	}
}
