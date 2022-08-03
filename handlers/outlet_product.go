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
