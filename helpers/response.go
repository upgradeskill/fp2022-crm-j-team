package helpers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/upgradeskill/fp2022-crm-j-team/schemas"
)

func APIResponse(c echo.Context, Message string, StatusCode int, Data interface{}) {
	jsonResponse := schemas.Responses{
		Message: Message,
		Data:    Data,
	}

	c.JSON(StatusCode, jsonResponse)
}

func ErrorResponse(c echo.Context, Error interface{}) {
	err := schemas.ErrorResponse{
		StatusCode: http.StatusBadRequest,
		Error:      Error,
	}

	c.JSON(err.StatusCode, err)
}
