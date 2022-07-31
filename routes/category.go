package routes

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func NewRouteCategory(router *echo.Echo) {

	router.GET("/api/v1/category", func(c echo.Context) error {
		return c.String(http.StatusOK, "Category Get")
	})

	router.GET("/api/v1/category/:id", func(c echo.Context) error {
		return c.String(http.StatusOK, "/category/:id")
	})

	router.POST("/api/v1/category", func(c echo.Context) error {
		return c.String(http.StatusOK, "Category Post")
	})

	router.PUT("/api/v1/category/:id", func(c echo.Context) error {
		return c.String(http.StatusOK, "category Put")
	})

	router.DELETE("/api/v1/category/:id", func(c echo.Context) error {
		return c.String(http.StatusOK, "Category Delete")
	})
}
