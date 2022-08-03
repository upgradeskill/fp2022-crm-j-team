package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/upgradeskill/fp2022-crm-j-team/handlers"
	"github.com/upgradeskill/fp2022-crm-j-team/middlewares"
	"github.com/upgradeskill/fp2022-crm-j-team/repositories"
	"github.com/upgradeskill/fp2022-crm-j-team/services"
	"gorm.io/gorm"
)

func NewRouteUser(db *gorm.DB, router *echo.Echo) {
	repository := repositories.NewRepositoryUser(db)
	service := services.NewServiceUser(repository)
	handler := handlers.NewHandlerUser(service)

	route := router.Group("/api/v1/")
	route.POST("users", handler.Create, middlewares.Auth("owner"))
	route.GET("users", handler.GetAll, middlewares.Auth("staff|owner"))
	route.GET("users/:id", handler.Get, middlewares.Auth("staff|owner"))
	route.PUT("users/:id", handler.Update, middlewares.Auth("owner"))
	route.DELETE("users/:id", handler.Delete, middlewares.Auth("owner"))

	route.POST("login", handler.Login)
}
