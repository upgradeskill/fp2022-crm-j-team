package middlewares

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/upgradeskill/fp2022-crm-j-team/helpers"
	"github.com/upgradeskill/fp2022-crm-j-team/pkg"
)

func Auth(allowedRole string) echo.MiddlewareFunc {
	return echo.MiddlewareFunc(func(next echo.HandlerFunc) echo.HandlerFunc {
		return echo.HandlerFunc(func(c echo.Context) error {
			if len(c.Request().Header["Authorization"]) > 0 {
				user, err := pkg.VerifyToken(c.Request().Header["Authorization"][0])
				if err != nil {
					helpers.APIResponse(c, "Cannot claim JWT Token", http.StatusUnauthorized, nil)
					return nil
				}

				var token string = user["id"].(string)
				var email string = user["email"].(string)
				var role string = user["role"].(string)
				var outletId string = ""

				fmt.Print(role)
				helpers.SetSession(token, email, role, outletId)

				if allowedRole == "*" {
					return next(c)
				}

				isAllow := checkIsAllowed(allowedRole, role)
				if isAllow {
					return next(c)
				}

				helpers.APIResponse(c, "Forbidden access", http.StatusForbidden, nil)
				return nil
			}

			helpers.APIResponse(c, "Unauthorized access", http.StatusUnauthorized, nil)
			return nil
		})
	})
}

func checkIsAllowed(allowedRoles string, userRole string) bool {
	resp := strings.Split(allowedRoles, "|")
	for _, val := range resp {
		if val == userRole {
			return true
		}
	}

	return false
}
