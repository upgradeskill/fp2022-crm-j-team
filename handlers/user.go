package handlers

import (
	"net/http"
	"strings"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/upgradeskill/fp2022-crm-j-team/helpers"
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
		helpers.APIResponse(c, "Login failed", http.StatusUnprocessableEntity, err.Error())
		return nil
	}

	validate = validator.New()
	eValidation := validate.Struct(userLoginSchema)
	if eValidation != nil {
		helpers.APIResponse(c, "Error validation", http.StatusUnprocessableEntity, strings.Split(eValidation.Error(), "tag\n"))
		return nil
	}

	user, err := h.user.Login(userLoginSchema)
	if err.Code != 0 {
		helpers.APIResponse(c, err.Type, err.Code, nil)
		return nil
	}

	helpers.APIResponse(c, "Login success", http.StatusOK, user)
	return nil
}

/**
* =====================================
* Handler Create New User
*======================================
 */
func (h *handleUser) Create(c echo.Context) error {
	userSchema := new(schemas.User)
	if err := c.Bind(userSchema); err != nil {
		helpers.APIResponse(c, "Failed create user", http.StatusUnprocessableEntity, err.Error())
		return nil
	}

	validate = validator.New()
	eValidation := validate.Struct(userSchema)
	if eValidation != nil {
		helpers.APIResponse(c, "Error validation", http.StatusUnprocessableEntity, strings.Split(eValidation.Error(), "tag\n"))
		return nil
	}

	user, err := h.user.Create(userSchema)
	if err.Code != 0 {
		helpers.APIResponse(c, err.Type, err.Code, strings.Split(eValidation.Error(), "tag\n"))
		return nil
	}

	helpers.APIResponse(c, "Success create user", http.StatusCreated, user)
	return nil
}

/**
* =====================================
* Handler Update User
*======================================
 */
func (h *handleUser) Update(c echo.Context) error {
	if c.Param("id") == "" {
		helpers.APIResponse(c, "param id is required", http.StatusUnprocessableEntity, nil)
		return nil
	}

	userSchema := new(schemas.User)
	userSchema.ID = c.Param("id")
	if err := c.Bind(userSchema); err != nil {
		helpers.APIResponse(c, "Failed update user", http.StatusUnprocessableEntity, err.Error())
		return nil
	}

	validate = validator.New()
	eValidation := validate.Struct(userSchema)
	if eValidation != nil {
		helpers.APIResponse(c, "Error validation", http.StatusUnprocessableEntity, strings.Split(eValidation.Error(), "tag\n"))
		return nil
	}

	user, err := h.user.Update(userSchema)
	if err.Code != 0 {
		helpers.APIResponse(c, err.Type, err.Code, strings.Split(eValidation.Error(), "tag\n"))
		return nil
	}

	helpers.APIResponse(c, "Success update user", http.StatusOK, user)
	return nil
}

/**
* =====================================
* Handler Get All User
*======================================
 */
func (h *handleUser) GetAll(c echo.Context) error {
	users, err := h.user.GetAll()
	if err.Code != 0 {
		helpers.APIResponse(c, err.Type, err.Code, nil)
		return nil
	}

	helpers.APIResponse(c, "Success get users", http.StatusOK, users)
	return nil
}

/**
* =====================================
* Handler Get Spesific User
*======================================
 */
func (h *handleUser) Get(c echo.Context) error {
	if c.Param("id") == "" {
		helpers.APIResponse(c, "Param id is required", http.StatusUnprocessableEntity, nil)
		return nil
	}

	userSchema := new(schemas.User)
	userSchema.ID = c.Param("id")

	user, err := h.user.Get(userSchema)
	if err.Code != 0 {
		helpers.APIResponse(c, err.Type, err.Code, nil)
		return nil
	}

	helpers.APIResponse(c, "Success get user", http.StatusOK, user)
	return nil
}

/**
* =====================================
* Handler Delete User
*======================================
 */
func (h *handleUser) Delete(c echo.Context) error {
	if c.Param("id") == "" {
		helpers.APIResponse(c, "Param id is required", http.StatusUnprocessableEntity, nil)
		return nil
	}

	userSchema := new(schemas.User)
	userSchema.ID = c.Param("id")

	user, err := h.user.Delete(userSchema)
	if err.Code != 0 {
		helpers.APIResponse(c, err.Type, err.Code, nil)
		return nil
	}

	helpers.APIResponse(c, "Success delete user", http.StatusOK, user)
	return nil
}
