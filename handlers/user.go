package handlers

import (
	"net/http"
	"strings"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/upgradeskill/fp2022-crm-j-team/ports"
	"github.com/upgradeskill/fp2022-crm-j-team/schemas"
)

var validate *validator.Validate

type handleUser struct {
	user ports.UserServiceInterface
}

func NewHandlerUser(user ports.UserServiceInterface) *handleUser {
	return &handleUser{user: user}
}

/**
* =====================================
* Handler Create New User
*======================================
 */
func (h *handleUser) Create(c echo.Context) error {
	user := new(schemas.User)
	if err := c.Bind(user); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, schemas.Responses{
			StatusCode: http.StatusUnprocessableEntity,
			Message:    "failed create user",
		})
	}

	validate = validator.New()
	eValidation := validate.Struct(user)
	if eValidation != nil {
		return c.JSON(http.StatusUnprocessableEntity, schemas.Responses{
			StatusCode: http.StatusUnprocessableEntity,
			Message:    "failed create user",
			Data:       strings.Split(eValidation.Error(), "tag\n"),
		})
	}

	createdUser, err := h.user.Create(user)
	if err.Code != 0 {
		return c.JSON(err.Code, schemas.Responses{
			StatusCode: err.Code,
			Message:    err.Type,
		})
	}

	return c.JSON(http.StatusCreated, schemas.Responses{
		StatusCode: http.StatusCreated,
		Message:    "Success create user",
		Data:       createdUser,
	})
}

/**
* =====================================
* Handler Update User
*======================================
 */
// func (h *handleUser) Update(c echo.Context) {
// 	user := new(schemas.User)
// 	if err := c.Bind(user); err != nil {
// 		return c.JSON(http.StatusBadRequest, "create failed")
// 	}

// 	_, err := h.user.Update(user)
// 	if err.Code != 0 {
// 		return c.JSON(http.StatusBadRequest, "create failed")
// 	}

// 	return c.JSON(http.StatusCreated, "success")
// }

/**
* =====================================
* Handler Get All User
*======================================
 */
// func (h *handleUser) GetAll(c echo.Context) {
// 	user := new(schemas.User)
// 	if err := c.Bind(user); err != nil {
// 		return c.JSON(http.StatusBadRequest, "create failed")
// 	}

// 	_, err := h.user.GetAll()
// 	if err.Code != 0 {
// 		return c.JSON(http.StatusBadRequest, "create failed")
// 	}

// 	return c.JSON(http.StatusCreated, "success")
// }

/**
* =====================================
* Handler Get Spesific User
*======================================
 */
// func (h *handleUser) Get(c echo.Context) {
// 	user := new(schemas.User)
// 	if err := c.Bind(user); err != nil {
// 		return c.JSON(http.StatusBadRequest, "create failed")
// 	}

// 	_, err := h.user.Get(user)
// 	if err.Code != 0 {
// 		return c.JSON(http.StatusBadRequest, "create failed")
// 	}

// 	return c.JSON(http.StatusCreated, "success")
// }

/**
* =====================================
* Handler Delete User
*======================================
 */
// func (h *handleUser) Delete(c echo.Context) {
// 	user := new(schemas.User)
// 	if err := c.Bind(user); err != nil {
// 		return c.JSON(http.StatusBadRequest, "create failed")
// 	}

// 	_, err := h.user.Delete(user)
// 	if err.Code != 0 {
// 		return c.JSON(http.StatusBadRequest, "create failed")
// 	}

// 	return c.JSON(http.StatusCreated, "success")
// }
