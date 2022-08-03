package schemas

import "github.com/golang-jwt/jwt"

type User struct {
	ID       string `json:"id"`
	Name     string `json:"name" validate:"required"`
	Email    string `json:"email" validate:"required"`
	Password string `json:"password"`
	Role     string `json:"role"`
	OutletId string `json:"outlet_id"`
	Page     int    `json:"page"`
	PageSize int    `json:"page_size"`
}

type UserLogin struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type LoggedUser struct {
	ID       string `json:"id"`
	Email    string `json:"email"`
	Role     string `json:"role"`
	OutletId string `json:"outlet_id"`
}

type JwtUserClaims struct {
	ID       string `json:"id"`
	Email    string `json:"email"`
	Role     string `json:"role"`
	OutletId string `json:"outlet_id"`
	jwt.StandardClaims
}
