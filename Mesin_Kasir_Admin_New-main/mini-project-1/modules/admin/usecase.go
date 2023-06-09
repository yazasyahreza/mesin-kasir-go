package admin

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type Usecase struct {
	Repo Repository
}

func (uc Usecase) Login(username, password string) (string, error) {
	admin, err := uc.Repo.CheckUsername(username, password)
	if err != nil {
		return "", err
	}

	claims := MyClaims{
		ID:   admin.ID,
		Name: admin.Name,
		StandardClaims: jwt.StandardClaims{
			Issuer:    APPLICATION_NAME,
			ExpiresAt: time.Now().Add(LOGIN_EXPIRATION_DURATION).Unix(),
		},
	}

	// mendeklarasikan algoritma yang akan digunakan untuk signing
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// signed token
	signedToken, err := token.SignedString(JWT_SIGNATURE_KEY)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}
