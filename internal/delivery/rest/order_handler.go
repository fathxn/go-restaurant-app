package rest

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
	"go-restaurant-app/internal/model"
	"net/http"
)

func (h *handler) Order(c echo.Context) error {
	var request model.OrderMenuRequest
	err := json.NewDecoder(c.Request().Body).Decode(&request)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
	}

	orderData, err := h.restaurantUsecase.Order(request)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": orderData,
	})
}

func (h *handler) GetOrderInfo(c echo.Context) error {
	orderID := c.Param("orderID")
	orderData, err := h.restaurantUsecase.GetOrderInfo(model.GetOrderInfoRequest{
		OrderID: orderID,
	})
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": orderData,
	})
}
