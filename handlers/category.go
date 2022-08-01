package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/upgradeskill/fp2022-crm-j-team/entities"
	"github.com/upgradeskill/fp2022-crm-j-team/schemas"
)

type handleCategory struct {
	category entities.Category
}

func Category(category entities.Category) *handleCategory {
	return &handleCategory{category: category}
}

/**
* =====================================
* Handler Create New Product
*======================================
 */

func (h *handleCategory) Create(c echo.Context) error {
	category := new(schemas.Category)

	if err := c.Bind(category); err != nil {
		return c.JSON(http.StatusBadRequest, "create failed")
	}

	_, err := h.category.Create(category)

	if err.Code != 0 {
		return c.JSON(http.StatusBadRequest, "create failed")
	}

	return c.JSON(http.StatusCreated, "success")
}
