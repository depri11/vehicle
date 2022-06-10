package helper

import (
	"errors"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type claims struct {
	Id    int    `json:"id"`
	Email string `json:"email"`
	jwt.StandardClaims
}

func NewToken(id int, email string) *claims {
	return &claims{
		Id:    id,
		Email: email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 5).Unix(),
		},
	}
}

func (c *claims) Create() (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	return token.SignedString([]byte(os.Getenv("JWT_SECRET_KEY")))
}

func CheckToken(token string) (*claims, error) {
	tokens, err := jwt.ParseWithClaims(token, &claims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET_KEY")), nil
	})

	if err != nil {
		return nil, err
	}

	// claims, ok := tokens.Claims.(*claims)
	claims, ok := tokens.Claims.(*claims)
	if !ok {
		err = errors.New("couldn't parse claims")
		return nil, err
	}

	return claims, nil
}

// func CheckToken(token string) (bool, error) {
// 	tokens, err := jwt.ParseWithClaims(token, &claims{}, func(t *jwt.Token) (interface{}, error) {
// 		return []byte(os.Getenv("JWT_SECRET_KEY")), nil
// 	})

// 	if err != nil {
// 		return false, err
// 	}

// 	// claims, ok := tokens.Claims.(*claims)

// 	// if ok && tokens.Valid {
// 	// 	return uint(claims.ID), nil
// 	// }

// 	return tokens.Valid, nil
// }
