package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/upgradeskill/fp2022-crm-j-team/helpers"
	"github.com/upgradeskill/fp2022-crm-j-team/ports"
	"github.com/upgradeskill/fp2022-crm-j-team/schemas"
)

type handleCategory struct {
	category ports.Category
}

func Category(category ports.Category) *handleCategory {
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
		helpers.ErrorResponse(c, err)
	}

	_, err := h.category.Create(category)

	if err.Code != 0 {
		helpers.ErrorResponse(c, err)
	}

	helpers.APIResponse(c, "Success create category", http.StatusCreated, nil)
	return nil
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
		helpers.ErrorResponse(c, err)
	}

	_, err := h.category.Update(category)

	if err.Code != 0 {
		helpers.ErrorResponse(c, err)
	}

	helpers.APIResponse(c, "Success update category", http.StatusCreated, nil)
	return nil
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

	res, err := h.category.Get(category)

	if err.Code != 0 {
		helpers.ErrorResponse(c, err)
	}

	if err.Type == "error_result_01" {
		helpers.APIResponse(c, fmt.Sprintf("Category data not found for this id %s ", category.ID), err.Code, nil)
		return nil
	}

	helpers.APIResponse(c, "Success get data", http.StatusOK, res)
	return nil
}

/**
* =====================================
* Handler Result All Category
*======================================
 */

func (h *handleCategory) GetAll(c echo.Context) error {

	category := new(schemas.Category)
	category.Name = c.QueryParam("name")
	category.Page, _ = strconv.Atoi(c.QueryParam("page"))
	category.PageSize, _ = strconv.Atoi(c.QueryParam("page_size"))

	res, err := h.category.GetAll(category)

	if err.Code != 0 {
		helpers.ErrorResponse(c, err)
	}

	if err.Type == "error_result_01" {
		helpers.APIResponse(c, fmt.Sprintf("Category data not found"), err.Code, nil)
		return nil
	}

	helpers.APIResponse(c, "Success get category data", http.StatusOK, res)
	return nil
}

/**
* ======================================
* Handler Delete Category By ID Category
*=======================================
 */

func (h *handleCategory) Delete(c echo.Context) error {
	category := new(schemas.Category)
	id := c.Param("id")
	category.ID = id

	if err := c.Bind(category); err != nil {
		helpers.ErrorResponse(c, err)
	}

	_, err := h.category.Delete(category)

	if err.Code != 0 {
		helpers.ErrorResponse(c, err)
	}

	helpers.APIResponse(c, "Success delete category", http.StatusOK, nil)
	return nil
}
