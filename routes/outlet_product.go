package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/upgradeskill/fp2022-crm-j-team/handlers"
	"github.com/upgradeskill/fp2022-crm-j-team/middlewares"
	"github.com/upgradeskill/fp2022-crm-j-team/repositories"
	"github.com/upgradeskill/fp2022-crm-j-team/services"
	"gorm.io/gorm"
)

func OutletProduct(db *gorm.DB, router *echo.Echo) {
	repository := repositories.OutletProduct(db)
	service := services.OutletProduct(repository)
	handler := handlers.OutletProduct(service)

	route := router.Group("/api/v1/outlet-product")
	route.POST("/create", handler.Create, middlewares.Auth("owner"))
}
