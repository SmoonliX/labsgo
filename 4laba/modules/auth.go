package modules

import (
    "time"
	"golang.org/x/crypto/bcrypt"
    "github.com/golang-jwt/jwt/v5"
)

func GenerateJWT(customerID int) (string, error) {
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "sub": customerID,
        "exp": time.Now().Add(time.Hour * 24).Unix(),
    })

    secretKey := "dc6bb2fb-6a41-42db-a32d-a8f7b5b86393"
    return token.SignedString([]byte(secretKey))
}

func VerifyPassword(hashedPassword, password string) error {
    return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}