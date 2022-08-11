package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/upgradeskill/fp2022-crm-j-team/handlers"
	"github.com/upgradeskill/fp2022-crm-j-team/middlewares"
	"github.com/upgradeskill/fp2022-crm-j-team/repositories"
	"github.com/upgradeskill/fp2022-crm-j-team/services"
	"gorm.io/gorm"
)

func Product(db *gorm.DB, router *echo.Echo) {
	repository := repositories.Product(db)
	service := services.Product(repository)
	handler := handlers.Product(service)

	route := router.Group("/api/v1/product")
	route.POST("", handler.Create, middlewares.Auth("owner"))
	route.GET("", handler.GetAll, middlewares.Auth("staff|owner"))
	route.GET("/:id", handler.Get, middlewares.Auth("staff|owner"))
	route.PUT("/:id", handler.Update, middlewares.Auth("owner"))
	route.DELETE("/:id", handler.Delete, middlewares.Auth("owner"))
}
