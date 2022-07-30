package routes

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func NewProduct(db *gorm.DB, router *echo.Echo) {
	fmt.Println("hello")
}
