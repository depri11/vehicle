package helper

import (
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type claims struct {
	Email string `json"email"`
	jwt.StandardClaims
}

func NewToken(email string) *claims {
	return &claims{
		Email: email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 5).Unix(),
		},
	}
}

func (c *claims) Create() string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	tokenString, _ := token.SignedString([]byte(os.Getenv("JWT_SECRET_KEY")))
	return tokenString
}

func CheckToken(token string) (bool, error) {
	tokens, err := jwt.ParseWithClaims(token, &claims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET_KEY")), nil
	})

	if err != nil {
		return false, err
	}

	return tokens.Valid, nil
}
