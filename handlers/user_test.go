package handlers

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/upgradeskill/fp2022-crm-j-team/helpers"
	"github.com/upgradeskill/fp2022-crm-j-team/models"
	"github.com/upgradeskill/fp2022-crm-j-team/repositories"
	"github.com/upgradeskill/fp2022-crm-j-team/services"
)

var (
	userCreate    = models.User{}
	payloadCreate = `{"name":"Wahyu Agung","email":"wahyu.agung@majoo.id","password":"12345","role":"owner","outlet_id":"1"}`
	payloadUpdate = `{"name":"Wahyu Agung","email":"wahyuagung@majoo.id","password":"12345","role":"owner","outlet_id":"1"}`
)

/**
* ==========================================
* Test Handler Create User
*===========================================
 */
func TestSuccessCreateUser(t *testing.T) {
	// Setup
	app := echo.New()
	db := helpers.SetupDatabaseTesting()
	repository := repositories.NewRepositoryUser(db)
	service := services.NewServiceUser(repository)
	handler := NewHandlerUser(service)

	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(payloadCreate))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	rec := httptest.NewRecorder()

	context := app.NewContext(req, rec)
	context.SetPath("/users")

	// Assertions
	if assert.NoError(t, handler.Create(context)) {
		assert.Equal(t, http.StatusCreated, rec.Code)
	}
}

/**
* ==========================================
* Test Handler Update User
*===========================================
 */
func TestSuccessUpdate(t *testing.T) {
	// Setup
	app := echo.New()
	db := helpers.SetupDatabaseTesting()
	repository := repositories.NewRepositoryUser(db)
	service := services.NewServiceUser(repository)
	handler := NewHandlerUser(service)

	user, _ := repository.Create(&models.User{
		Name:     "Wahyu",
		Email:    "wahyuagung@gmail.com",
		Password: "12345",
		Role:     "owner",
		OutletId: "123456",
	})
	userCreate = *user

	req := httptest.NewRequest(http.MethodPut, "/", strings.NewReader(payloadUpdate))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	rec := httptest.NewRecorder()

	context := app.NewContext(req, rec)
	context.SetPath("/users/:id")
	context.SetParamNames("id")
	context.SetParamValues(user.ID)

	// Assertions
	if assert.NoError(t, handler.Update(context)) {
		assert.Equal(t, http.StatusOK, rec.Code, rec.Body.String())
	}
}

/**
* ==========================================
* Test Handler Get Spesific User
*===========================================
 */
func TestSuccessGetUser(t *testing.T) {
	// Setup
	app := echo.New()
	db := helpers.SetupDatabaseTesting()
	repository := repositories.NewRepositoryUser(db)
	service := services.NewServiceUser(repository)
	handler := NewHandlerUser(service)

	req := httptest.NewRequest(http.MethodGet, "/", strings.NewReader(payloadCreate))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	rec := httptest.NewRecorder()

	context := app.NewContext(req, rec)
	context.SetPath("/users/:id")
	context.SetParamNames("id")
	context.SetParamValues(userCreate.ID)

	// Assertions
	if assert.NoError(t, handler.Get(context)) {
		assert.Equal(t, http.StatusOK, rec.Code, rec.Body.String())
	}
}

/**
* ==========================================
* Test Handler Get All User
*===========================================
 */
func TestSuccessGetAllUser(t *testing.T) {
	// Setup
	app := echo.New()
	db := helpers.SetupDatabaseTesting()
	repository := repositories.NewRepositoryUser(db)
	service := services.NewServiceUser(repository)
	handler := NewHandlerUser(service)

	req := httptest.NewRequest(http.MethodGet, "/", strings.NewReader(payloadCreate))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	rec := httptest.NewRecorder()

	context := app.NewContext(req, rec)
	context.SetPath("/users")

	// Assertions
	if assert.NoError(t, handler.GetAll(context)) {
		assert.Equal(t, http.StatusOK, rec.Code, rec.Body.String())
	}
}

/**
* ==========================================
* Test Handler Delete User
*===========================================
 */
func TestSuccessDeleteUser(t *testing.T) {
	// Setup
	app := echo.New()
	db := helpers.SetupDatabaseTesting()
	repository := repositories.NewRepositoryUser(db)
	service := services.NewServiceUser(repository)
	handler := NewHandlerUser(service)

	req := httptest.NewRequest(http.MethodGet, "/", strings.NewReader(payloadCreate))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	rec := httptest.NewRecorder()

	context := app.NewContext(req, rec)
	context.SetPath("/users/:id")
	context.SetParamNames("id")
	context.SetParamValues(userCreate.ID)

	// Assertions
	if assert.NoError(t, handler.Delete(context)) {
		assert.Equal(t, http.StatusOK, rec.Code, rec.Body.String())
	}
}
