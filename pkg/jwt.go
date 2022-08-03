package pkg

import (
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/upgradeskill/fp2022-crm-j-team/schemas"
)

var SecretPublicKey = "tes"

func GenerateTokenJwt(ID string, Email string, Role string, OutletId string) (string, error) {
	// Set custom claims
	claims := schemas.JwtUserClaims{
		ID,
		Email,
		Role,
		OutletId,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		},
	}

	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(GodotEnv("JWT_SECRET_KEY")))
	if err != nil {
		return t, err
	}

	return t, nil
}

func VerifyToken(accessToken string) (jwt.MapClaims, error) {

	token, err := jwt.Parse(accessToken, func(token *jwt.Token) (interface{}, error) {
		return []byte(GodotEnv("JWT_SECRET_KEY")), nil
	})

	if err != nil {
		return nil, err
	}

	claims := token.Claims.(jwt.MapClaims)
	return claims, nil
}
