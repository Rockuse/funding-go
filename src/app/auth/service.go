package auth

import (
	"errors"

	"github.com/dgrijalva/jwt-go"
)

type Service interface {
	GenerateToken(data interface{}) (string, error)
	ValidateToken(token string) (*jwt.Token, error)
}

type jwtService struct {
}

var SECRET_KEY = []byte("Fahmi Muzakki")

func NewService() *jwtService {
	return &jwtService{}
}

func (j *jwtService) GenerateToken(data interface{}) (string, error) {
	claim := jwt.MapClaims{}
	claim["data"] = data
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	signedToken, err := token.SignedString(SECRET_KEY)
	if err != nil {
		return signedToken, err
	}
	return signedToken, nil
}

func (j *jwtService) ValidateToken(token string) (*jwt.Token, error) {
	toke, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		_, ok := t.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("invalid token")
		}
		return []byte(SECRET_KEY), nil
	})
	if err != nil {
		return toke, errors.New("validate Error")
	}
	return toke, nil
}
