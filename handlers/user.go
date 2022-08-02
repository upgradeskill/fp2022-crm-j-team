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
* Handler Login
*======================================
 */
func (h *handleUser) Login(c echo.Context) error {
	userLoginSchema := new(schemas.UserLogin)
	if err := c.Bind(userLoginSchema); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, schemas.Responses{
			StatusCode: http.StatusUnprocessableEntity,
			Message:    "failed login user",
		})
	}

	validate = validator.New()
	eValidation := validate.Struct(userLoginSchema)
	if eValidation != nil {
		return c.JSON(http.StatusUnprocessableEntity, schemas.Responses{
			StatusCode: http.StatusUnprocessableEntity,
			Message:    "failed login",
			Data:       strings.Split(eValidation.Error(), "tag\n"),
		})
	}

	user, err := h.user.Login(userLoginSchema)
	if err.Code != 0 {
		return c.JSON(err.Code, schemas.Responses{
			StatusCode: err.Code,
			Message:    err.Type,
		})
	}

	return c.JSON(http.StatusOK, schemas.Responses{
		StatusCode: http.StatusOK,
		Message:    "success login",
		Data:       user,
	})
}

/**
* =====================================
* Handler Create New User
*======================================
 */
func (h *handleUser) Create(c echo.Context) error {
	userSchema := new(schemas.User)
	if err := c.Bind(userSchema); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, schemas.Responses{
			StatusCode: http.StatusUnprocessableEntity,
			Message:    "failed create user",
		})
	}

	validate = validator.New()
	eValidation := validate.Struct(userSchema)
	if eValidation != nil {
		return c.JSON(http.StatusUnprocessableEntity, schemas.Responses{
			StatusCode: http.StatusUnprocessableEntity,
			Message:    "failed create user",
			Data:       strings.Split(eValidation.Error(), "tag\n"),
		})
	}

	user, err := h.user.Create(userSchema)
	if err.Code != 0 {
		return c.JSON(err.Code, schemas.Responses{
			StatusCode: err.Code,
			Message:    err.Type,
		})
	}

	return c.JSON(http.StatusCreated, schemas.Responses{
		StatusCode: http.StatusCreated,
		Message:    "Success create user",
		Data:       user,
	})
}

/**
* =====================================
* Handler Update User
*======================================
 */
func (h *handleUser) Update(c echo.Context) error {
	if c.Param("id") == "" {
		return c.JSON(http.StatusOK, schemas.Responses{
			StatusCode: http.StatusUnprocessableEntity,
			Message:    "param id is required",
		})
	}

	userSchema := new(schemas.User)
	userSchema.ID = c.Param("id")
	if err := c.Bind(userSchema); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, schemas.Responses{
			StatusCode: http.StatusUnprocessableEntity,
			Message:    "failed update user",
		})
	}

	validate = validator.New()
	eValidation := validate.Struct(userSchema)
	if eValidation != nil {
		return c.JSON(http.StatusUnprocessableEntity, schemas.Responses{
			StatusCode: http.StatusUnprocessableEntity,
			Message:    "failed update user",
			Data:       strings.Split(eValidation.Error(), "tag\n"),
		})
	}

	user, err := h.user.Update(userSchema)
	if err.Code != 0 {
		return c.JSON(err.Code, schemas.Responses{
			StatusCode: err.Code,
			Message:    err.Type,
		})
	}

	return c.JSON(http.StatusOK, schemas.Responses{
		StatusCode: http.StatusOK,
		Message:    "Success update user",
		Data:       user,
	})
}

/**
* =====================================
* Handler Get All User
*======================================
 */
func (h *handleUser) GetAll(c echo.Context) error {
	users, err := h.user.GetAll()
	if err.Code != 0 {
		return c.JSON(err.Code, schemas.Responses{
			StatusCode: err.Code,
			Message:    err.Type,
		})
	}

	return c.JSON(http.StatusOK, schemas.Responses{
		StatusCode: http.StatusOK,
		Message:    "Success update user",
		Data:       users,
	})
}

/**
* =====================================
* Handler Get Spesific User
*======================================
 */
func (h *handleUser) Get(c echo.Context) error {
	if c.Param("id") == "" {
		return c.JSON(http.StatusOK, schemas.Responses{
			StatusCode: http.StatusUnprocessableEntity,
			Message:    "param id is required",
		})
	}

	userSchema := new(schemas.User)
	userSchema.ID = c.Param("id")

	user, err := h.user.Get(userSchema)
	if err.Code != 0 {
		return c.JSON(err.Code, schemas.Responses{
			StatusCode: err.Code,
			Message:    err.Type,
		})
	}

	return c.JSON(http.StatusOK, schemas.Responses{
		StatusCode: http.StatusOK,
		Message:    "found 1 user",
		Data:       user,
	})
}

/**
* =====================================
* Handler Delete User
*======================================
 */
func (h *handleUser) Delete(c echo.Context) error {
	if c.Param("id") == "" {
		return c.JSON(http.StatusOK, schemas.Responses{
			StatusCode: http.StatusUnprocessableEntity,
			Message:    "param id is required",
		})
	}

	userSchema := new(schemas.User)
	userSchema.ID = c.Param("id")

	user, err := h.user.Delete(userSchema)
	if err.Code != 0 {
		return c.JSON(err.Code, schemas.Responses{
			StatusCode: err.Code,
			Message:    err.Type,
		})
	}

	return c.JSON(http.StatusOK, schemas.Responses{
		StatusCode: http.StatusOK,
		Message:    "success delete user",
		Data:       user,
	})
}
