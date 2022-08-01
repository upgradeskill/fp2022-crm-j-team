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
* Handler Create New Category
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

/**
* ======================================
* Handler Update Category By ID Category
*=======================================
 */

func (h *handleCategory) Update(c echo.Context) error {
	category := new(schemas.Category)
	id := c.Param("id")
	category.ID = id

	if err := c.Bind(category); err != nil {
		return c.JSON(http.StatusBadRequest, "update failed")
	}

	_, err := h.category.Update(category)

	if err.Code != 0 {
		return c.JSON(http.StatusBadRequest, "update failed")
	}

	return c.JSON(http.StatusCreated, "success")
}

/**
* =====================================
* Handler Result Category by ID Category
*======================================
 */

func (h *handleCategory) Get(c echo.Context) error {
	category := new(schemas.Category)
	id := c.Param("id")
	category.ID = id

	if err := c.Bind(category); err != nil {
		return c.JSON(http.StatusBadRequest, "get failed")
	}

	res, err := h.category.Get(category)

	if err.Code != 0 {
		return c.JSON(http.StatusBadRequest, "create failed")
	}

	return c.JSON(http.StatusOK, res)
}
