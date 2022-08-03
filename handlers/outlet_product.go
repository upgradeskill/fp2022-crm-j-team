package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/upgradeskill/fp2022-crm-j-team/helpers"
	"github.com/upgradeskill/fp2022-crm-j-team/ports"
	"github.com/upgradeskill/fp2022-crm-j-team/schemas"
)

type handleOutletProduct struct {
	outletProduct ports.OutletProduct
}

func OutletProduct(outletProduct ports.OutletProduct) *handleOutletProduct {
	return &handleOutletProduct{outletProduct: outletProduct}
}

/**
* =====================================
* Handler Create New OutletProduct
*======================================
 */

func (h *handleOutletProduct) Create(c echo.Context) error {
	outletProduct := new(schemas.OutletProduct)

	if err := c.Bind(outletProduct); err != nil {
		helpers.APIResponse(c, err.Error(), http.StatusBadRequest, nil)
		return nil
	}

	_, err := h.outletProduct.Create(outletProduct)
	if err.Code != 0 {
		helpers.APIResponse(c, "Failed to create outlet product", err.Code, nil)
		return nil
	}

	helpers.APIResponse(c, "Success", http.StatusCreated, nil)
	return nil
}

func (h *handleOutletProduct) Update(c echo.Context) error {
	outletProduct := new(schemas.OutletProduct)

	if err := c.Bind(outletProduct); err != nil {
		helpers.APIResponse(c, err.Error(), http.StatusBadRequest, nil)
		return nil
	}

	outletProduct.ID = c.Param("id")

	if outletProduct.ID == "" {
		helpers.APIResponse(c, "Id is required", http.StatusBadRequest, nil)
		return nil
	}

	_, err := h.outletProduct.Update(outletProduct)
	if err.Code != 0 {
		helpers.APIResponse(c, "Failed to update outlet product", err.Code, nil)
		return nil
	}

	helpers.APIResponse(c, "Success", http.StatusOK, nil)
	return nil
}

func (h *handleOutletProduct) Delete(c echo.Context) error {
	outletProduct := new(schemas.OutletProduct)
	outletProduct.ID = c.Param("id")
	if outletProduct.ID == "" {
		helpers.APIResponse(c, "Id is required", http.StatusBadRequest, nil)
		return nil
	}

	_, err := h.outletProduct.Delete(outletProduct)
	if err.Code != 0 {
		helpers.APIResponse(c, "Failed to delete outlet product", err.Code, nil)
		return nil
	}

	helpers.APIResponse(c, "Success", http.StatusOK, nil)
	return nil
}
