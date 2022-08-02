package handlers

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"

	"github.com/upgradeskill/fp2022-crm-j-team/helpers"
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
		helpers.APIResponse(c, err.Error(), http.StatusBadRequest, nil)
		return nil
	}

	_, err := h.product.Create(product)
	if err.Code != 0 {
		helpers.APIResponse(c, "Failed to create product", err.Code, nil)
		return nil
	}

	helpers.APIResponse(c, "Success", http.StatusCreated, nil)
	return nil
}

func (h *handleProduct) Get(c echo.Context) error {
	product := new(schemas.Product)
	product.ID = c.Param("id")

	if product.ID == "" {
		helpers.APIResponse(c, "Id is required", http.StatusBadRequest, nil)
		return nil
	}

	res, err := h.product.Get(product)
	if err.Code != 0 {
		helpers.APIResponse(c, "Failed to get product", err.Code, nil)
		return nil
	}

	if res == nil {
		helpers.APIResponse(c, "Product not found", http.StatusNotFound, nil)
		return nil
	}

	helpers.APIResponse(c, "Success", http.StatusOK, res)
	return nil
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
		helpers.APIResponse(c, "Failed to get products", err.Code, nil)
		return nil
	}

	helpers.APIResponse(c, "Success", http.StatusOK, res)
	return nil
}

func (h *handleProduct) Update(c echo.Context) error {
	product := new(schemas.Product)
	if err := c.Bind(product); err != nil {
		return c.JSON(http.StatusBadRequest, "update failed")
	}

	product.ID = c.Param("id")
	if product.ID == "" {
		helpers.APIResponse(c, "Id is required", http.StatusBadRequest, nil)
		return nil
	}

	_, err := h.product.Update(product)
	if err.Code != 0 {
		helpers.APIResponse(c, "Failed to get products", err.Code, nil)
		return nil
	}

	helpers.APIResponse(c, "Success", http.StatusOK, nil)
	return nil
}

func (h *handleProduct) Delete(c echo.Context) error {
	product := new(schemas.Product)
	product.ID = c.Param("id")
	if product.ID == "" {
		helpers.APIResponse(c, "Id is required", http.StatusBadRequest, nil)
		return nil
	}

	_, err := h.product.Delete(product)
	if err.Code != 0 {
		helpers.APIResponse(c, "Failed to get products", err.Code, nil)
		return nil
	}

	helpers.APIResponse(c, "Success", http.StatusOK, nil)
	return nil
}
