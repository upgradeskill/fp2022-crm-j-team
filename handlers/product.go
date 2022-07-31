package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/upgradeskill/fp2022-crm-j-team/entities"
	"github.com/upgradeskill/fp2022-crm-j-team/schemas"
)

type handleProduct struct {
	product entities.Product
}

func Product(product entities.Product) *handleProduct {
	return &handleProduct{product: product}
}

/**
* =====================================
* Handler Create New Product
*======================================
 */

func (h *handleProduct) Create(c echo.Context) error {
	product := new(schemas.Product)
	if err := c.Bind(product); err != nil {
		return c.JSON(http.StatusBadRequest, "create failed")
	}

	_, err := h.product.Create(product)
	if err.Code != 0 {
		return c.JSON(http.StatusBadRequest, "create failed")
	}

	return c.JSON(http.StatusCreated, "success")
}
