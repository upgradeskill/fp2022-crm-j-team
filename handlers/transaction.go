package handlers

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"

	"github.com/upgradeskill/fp2022-crm-j-team/helpers"
	"github.com/upgradeskill/fp2022-crm-j-team/ports"
	"github.com/upgradeskill/fp2022-crm-j-team/schemas"
)

type Transaction struct {
	service ports.Transaction
}

func NewTransaction(service ports.Transaction) *Transaction {
	return &Transaction{service: service}
}

func (handler *Transaction) Create(c echo.Context) error {
	transaction := new(schemas.Transaction)

	if err := c.Bind(transaction); err != nil {
		helpers.APIResponse(c, err.Error(), http.StatusBadRequest, nil)
		return nil
	}

	_, err := handler.service.Create(transaction)
	if err.Code != 0 {
		helpers.APIResponse(c, "Failed to create transaction", err.Code, nil)
		return nil
	}

	helpers.APIResponse(c, "Success", http.StatusCreated, nil)
	return nil
}

func (handler *Transaction) FindById(c echo.Context) error {

	transactionId := c.Param("id")
	res, err := handler.service.FindById(transactionId)

	if err.Code != 0 {
		helpers.APIResponse(c, "Failed to get transaction", err.Code, nil)
		return nil
	}

	if res == nil {
		helpers.APIResponse(c, "Transaction not found", http.StatusNotFound, nil)
		return nil
	}

	helpers.APIResponse(c, "Success", http.StatusOK, res)
	return nil
}

func (handler *Transaction) FindAll(c echo.Context) error {

	page, _ := strconv.Atoi(c.QueryParam("page"))
	pageSize, _ := strconv.Atoi(c.QueryParam("page_size"))

	transaction := schemas.TransactionPage{
		Page:     page,
		PageSize: pageSize,
	}

	res, err := handler.service.FindAll(&transaction)
	if err.Code != 0 {
		helpers.APIResponse(c, "Failed to get transactions", err.Code, nil)
		return nil
	}

	helpers.APIResponse(c, "Success", http.StatusOK, res)
	return nil
}

func (handler *Transaction) Update(c echo.Context) error {
	transaction := new(schemas.Transaction)
	if err := c.Bind(&transaction); err != nil {
		return c.JSON(http.StatusBadRequest, "update failed")
	}

	transactionId := c.Param("id")

	_, err := handler.service.Update(transactionId, transaction)
	if err.Code != 0 {
		helpers.APIResponse(c, "Failed to get transactions", err.Code, nil)
		return nil
	}

	helpers.APIResponse(c, "Success", http.StatusOK, nil)
	return nil
}

func (handler *Transaction) Delete(c echo.Context) error {

	transactionId := c.Param("id")

	_, err := handler.service.Delete(transactionId)
	if err.Code != 0 {
		helpers.APIResponse(c, "Failed to get transactions", err.Code, nil)
		return nil
	}

	helpers.APIResponse(c, "Success", http.StatusOK, nil)
	return nil
}
