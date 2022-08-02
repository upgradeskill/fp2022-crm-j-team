package pkg

import (
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/sirupsen/logrus"
	"github.com/upgradeskill/fp2022-crm-j-team/schemas"
)

var SecretPublicKey = "tes"

func GenerateToken(user *schemas.LoggedUser) (string, error) {
	// Set custom claims
	claims := schemas.JwtUserClaims{
		user.ID,
		user.Email,
		user.Role,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		},
	}

	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(SecretPublicKey))
	if err != nil {
		logrus.Error(err.Error())
		return t, err
	}

	return t, nil
}

func VerifyToken(accessToken string) (*jwt.Token, error) {

	token, err := jwt.Parse(accessToken, func(token *jwt.Token) (interface{}, error) {
		return []byte(SecretPublicKey), nil
	})

	if err != nil {
		logrus.Error(err.Error())
		return nil, err
	}

	return token, nil
}
