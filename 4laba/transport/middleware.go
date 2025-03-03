package transport

import (
    "net/http"

    "github.com/gin-gonic/gin"
    "github.com/golang-jwt/jwt/v5"
)

func AuthRequired(c *gin.Context) {
    tokenString := c.GetHeader("Authorization")
    if tokenString == "" {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Токен отсутствует"})
        c.Abort()
        return
    }

    token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
        return []byte("dc6bb2fb-6a41-42db-a32d-a8f7b5b86393"), nil
    })

    if err != nil || !token.Valid {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Неверный токен"})
        c.Abort()
        return
    }

    claims := token.Claims.(jwt.MapClaims)
    c.Set("customerID", int(claims["sub"].(float64)))
    c.Next()
}