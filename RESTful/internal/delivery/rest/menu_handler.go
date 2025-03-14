package rest

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *handler) GetMenuList(c echo.Context) error {
	menuType := c.FormValue("menu_type")
	menuData, err := h.restoUsecase.GetMenuList(menuType)
	if err != nil {
		fmt.Printf("got error: %s\n", err)
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"data": menuData,
	})
}
