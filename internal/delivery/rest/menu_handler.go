package rest

import (
	"github.com/labstack/echo/v4"
	"go-restaurant-app/internal/tracing"
	"net/http"
)

func (h *handler) GetMenuList(c echo.Context) error {
	ctx, span := tracing.CreateSpan(c.Request().Context(), "GetMenuList")
	defer span.End()

	menuType := c.FormValue("menu_type")
	menuData, err := h.restaurantUsecase.GetMenuList(ctx, menuType)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{"data": menuData})
}
