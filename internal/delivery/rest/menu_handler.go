package rest

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func (h *handler) GetMenu(c echo.Context) error {
	menuType := c.FormValue("menu_type")
	menuData, err := h.restaurantUsecase.GetMenu(menuType)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{"data": menuData})
}
