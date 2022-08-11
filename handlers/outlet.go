package handlers

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/upgradeskill/fp2022-crm-j-team/helpers"
	"github.com/upgradeskill/fp2022-crm-j-team/ports"
	"github.com/upgradeskill/fp2022-crm-j-team/schemas"
)

type handleOutlet struct {
	outlet ports.Outlet
}

func Outlet(outlet ports.Outlet) *handleOutlet {
	return &handleOutlet{outlet: outlet}
}

/**
* =====================================
* Handler Create New Outlet
*======================================
 */

func (h *handleOutlet) Create(c echo.Context) error {
	outlet := new(schemas.Outlet)

	if err := c.Bind(outlet); err != nil {
		helpers.ErrorResponse(c, err)
	}

	_, err := h.outlet.Create(outlet)

	if err.Code != 0 {
		helpers.ErrorResponse(c, err)
	}

	helpers.APIResponse(c, "Success create outlet", http.StatusCreated, nil)
	return nil
}

/**
* ======================================
* Handler Update Outlet By ID Outlet
*=======================================
 */

func (h *handleOutlet) Update(c echo.Context) error {
	outlet := new(schemas.Outlet)
	id := c.Param("id")
	outlet.ID = id

	if err := c.Bind(outlet); err != nil {
		helpers.ErrorResponse(c, err)
	}

	_, err := h.outlet.Update(outlet)

	if err.Code != 0 {
		helpers.ErrorResponse(c, err)
	}

	helpers.APIResponse(c, "Success update outlet", http.StatusCreated, nil)
	return nil
}

/**
* =====================================
* Handler Result Outlet by ID Outlet
*======================================
 */

func (h *handleOutlet) Get(c echo.Context) error {
	outlet := new(schemas.Outlet)
	id := c.Param("id")
	outlet.ID = id

	res, err := h.outlet.Get(outlet)

	if err.Code != 0 {
		helpers.ErrorResponse(c, err)
	}

	helpers.APIResponse(c, "Success get data", http.StatusOK, res)
	return nil
}

/**
* =====================================
* Handler Result All Outlet
*======================================
 */

func (h *handleOutlet) GetAll(c echo.Context) error {

	outlet := new(schemas.Outlet)
	outlet.Name = c.QueryParam("name")
	outlet.Page, _ = strconv.Atoi(c.QueryParam("page"))
	outlet.PageSize, _ = strconv.Atoi(c.QueryParam("page_size"))

	res, err := h.outlet.GetAll(outlet)

	if err.Code != 0 {
		helpers.ErrorResponse(c, err)
	}

	helpers.APIResponse(c, "Success get outlet data", http.StatusOK, res)
	return nil
}

/**
* ======================================
* Handler Delete Outlet By ID Outlet
*=======================================
 */

func (h *handleOutlet) Delete(c echo.Context) error {
	outlet := new(schemas.Outlet)
	id := c.Param("id")
	outlet.ID = id

	if err := c.Bind(outlet); err != nil {
		helpers.ErrorResponse(c, err)
	}

	_, err := h.outlet.Delete(outlet)

	if err.Code != 0 {
		helpers.ErrorResponse(c, err)
	}

	helpers.APIResponse(c, "Success delete outlet", http.StatusOK, nil)
	return nil
}
