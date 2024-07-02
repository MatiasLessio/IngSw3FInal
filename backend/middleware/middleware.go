package middleware

import (
	"fmt"
	"net/http"
	"strings"

	client "backend/clients/users"
	e "backend/utilities"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var jwtSecret = []byte("secret_key")

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		userID, err := GetUserIdByToken(c)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			c.Abort()
			return
		}

		// You can add userID to the context if needed
		c.Set("userID", userID)
		c.Next()
	}
}

func GetUserIdByToken(c *gin.Context) (int, e.ApiError) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		return 0, e.NewBadRequestApiError("no token provided")
	}

	// Extraer el token eliminando el prefijo "Bearer "
	tokenString := strings.TrimPrefix(authHeader, "Bearer ")
	if tokenString == authHeader {
		return 0, e.NewBadRequestApiError("invalid toekn format")
	}

	// Decodificar el token JWT verificando la firma
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return jwtSecret, nil
	})

	if err != nil {
		return 0, e.NewBadRequestApiError("invalid token")
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if idUser, ok := claims["user_id"].(float64); ok {
			user, err := client.GetUserById(int(idUser))
			if err != nil || user.UserId == 0 {
				return 0, e.NewBadRequestApiError("invalid token claims")
			}
			return int(idUser), nil
		}
	}

	return 0, e.NewBadRequestApiError("invalid token claims")
}
