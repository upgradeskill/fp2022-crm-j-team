package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/upgradeskill/fp2022-crm-j-team/handlers"
	"github.com/upgradeskill/fp2022-crm-j-team/repositories"
	"github.com/upgradeskill/fp2022-crm-j-team/services"
	"gorm.io/gorm"
)

func Transaction(db *gorm.DB, router *echo.Echo) {

	repository := repositories.NewTransaction(db)
	service := services.NewTransaction(repository)
	handler := handlers.NewTransaction(service)

	route := router.Group("/api/v1/transaction")
	route.POST("", handler.Create)
	route.GET("", handler.FindAll)
	route.GET("/:id", handler.FindById)
	route.PUT("/:id", handler.Update)
	route.DELETE("/:id", handler.Delete)
}
