package handlers

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/upgradeskill/fp2022-crm-j-team/helpers"
	"github.com/upgradeskill/fp2022-crm-j-team/repositories"
	"github.com/upgradeskill/fp2022-crm-j-team/services"
)

var (
	createUserSuccess = `{"name":"Wahyu Agung","email":"wahyu.agung@majoo.id","password":"12345","role":"owner","outlet_id":"1"}`
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

	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(createUserSuccess))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	context := app.NewContext(req, rec)
	handler := NewHandlerUser(service)

	// Assertions
	if assert.NoError(t, handler.Create(context)) {
		assert.Equal(t, http.StatusCreated, rec.Code)
	}
}

func TestFailedCreateUserErrorValidation(t *testing.T) {
	// Setup
	app := echo.New()
	db := helpers.SetupDatabaseTesting()
	repository := repositories.NewRepositoryUser(db)
	service := services.NewServiceUser(repository)

	req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(createUserSuccess))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	context := app.NewContext(req, rec)
	handler := NewHandlerUser(service)

	// Assertions
	if assert.NoError(t, handler.Create(context)) {
		assert.Equal(t, http.StatusUnprocessableEntity, rec.Code)
	}
}
