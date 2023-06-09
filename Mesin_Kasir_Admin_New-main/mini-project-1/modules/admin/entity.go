package admin

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type User struct {
	ID       float64    `gorm:"primaryKey"`
	Name     string `gorm:"varchar" json:"name"`
	Username string `gorm:"varchar" json:"username"`
	Password string `gorm:"varchar" json:"password"`
}

type MyClaims struct {
	ID    float64 `json:"id"`
	Name string `json:"nama_lengkap"`
	jwt.StandardClaims
}

type AdminResponse struct {
	Message string
	Data    jwt.MapClaims
}

var APPLICATION_NAME = "yaza barudak"
var LOGIN_EXPIRATION_DURATION = time.Duration(1) * time.Hour
var JWT_SIGNATURE_KEY = []byte("opewfjdi3f84f339fu3")