package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/upgradeskill/fp2022-crm-j-team/handlers"
	"github.com/upgradeskill/fp2022-crm-j-team/middlewares"
	"github.com/upgradeskill/fp2022-crm-j-team/repositories"
	"github.com/upgradeskill/fp2022-crm-j-team/services"
	"gorm.io/gorm"
)

func Outlet(db *gorm.DB, router *echo.Echo) {

	repository := repositories.Outlet(db)
	service := services.Outlet(repository)
	handler := handlers.Outlet(service)

	route := router.Group("/api/v1/outlet")
	route.GET("", handler.GetAll, middlewares.Auth("owner"))
	route.GET("/:id", handler.Get, middlewares.Auth("owner"))
	route.POST("", handler.Create, middlewares.Auth("owner"))
	route.PUT("/:id", handler.Update, middlewares.Auth("owner"))
	route.DELETE("/:id", handler.Delete, middlewares.Auth("owner"))
}
