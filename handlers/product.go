package handlers

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"

	"github.com/upgradeskill/fp2022-crm-j-team/ports"
	"github.com/upgradeskill/fp2022-crm-j-team/schemas"
)

type handleProduct struct {
	product ports.Product
}

func Product(product ports.Product) *handleProduct {
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

func (h *handleProduct) Get(c echo.Context) error {
	product := new(schemas.Product)
	product.ID = c.Param("id")

	res, err := h.product.Get(product)
	if err.Code != 0 {
		return c.JSON(http.StatusBadRequest, "get data failed")
	}

	return c.JSON(http.StatusCreated, res)
}

func (h *handleProduct) GetAll(c echo.Context) error {
	product := new(schemas.Product)
	product.ID = c.QueryParam("id")
	product.Name = c.QueryParam("name")
	product.SKU, _ = strconv.ParseUint(c.QueryParam("sku"), 10, 64)
	product.CategoryId = c.QueryParam("category_id")
	product.Page, _ = strconv.Atoi(c.QueryParam("page"))
	product.PageSize, _ = strconv.Atoi(c.QueryParam("page_size"))

	res, err := h.product.GetAll(product)
	if err.Code != 0 {
		return c.JSON(http.StatusBadRequest, "get data failed")
	}

	return c.JSON(http.StatusCreated, res)
}
