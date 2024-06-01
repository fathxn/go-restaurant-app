package rest

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
	"go-restaurant-app/internal/model"
	"go-restaurant-app/internal/tracing"
	"net/http"
)

func (h *handler) RegisterUser(c echo.Context) error {
	ctx, span := tracing.CreateSpan(c.Request().Context(), "RegisterUser")
	defer span.End()

	var request model.RegisterRequest
	err := json.NewDecoder(c.Request().Body).Decode(&request)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
	}

	userData, err := h.restaurantUsecase.RegisterUser(ctx, request)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": userData,
	})
}

func (h *handler) LoginUser(c echo.Context) error {
	ctx, span := tracing.CreateSpan(c.Request().Context(), "LoginUser")
	defer span.End()

	var request model.LoginRequest
	if err := json.NewDecoder(c.Request().Body).Decode(&request); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
	}

	userData, err := h.restaurantUsecase.LoginUser(ctx, request)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": userData,
	})
}
