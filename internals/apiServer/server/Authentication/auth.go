package Authentication

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"net/http"
	"time"
)

const (
	signingKey     = "my_secret_key"
	expirationTime = 3600
)

type JwtCustomClaims struct {
	Name string `json:"name"`
	Role bool   `json:"role"`
	jwt.RegisteredClaims
}

func CreateJwt(username string, role bool) (string, error) {
	claims := &JwtCustomClaims{
		Name: username,
		Role: role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(expirationTime * time.Second)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte(signingKey))
	if err != nil {
		return "", err
	}
	return t, nil
}

func AuthMiddleware(requiredRole bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is missing"})
			c.Abort()
			return
		}

		tokenString := authHeader[len("Bearer "):]
		token, err := jwt.ParseWithClaims(tokenString, &JwtCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(signingKey), nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
			c.Abort()
			return
		}

		claims, ok := token.Claims.(*JwtCustomClaims)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Failed to parse token claims"})
			c.Abort()
			return
		}

		if claims.Role != requiredRole {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized access"})
			c.Abort()
			return
		}

		c.Set("username", claims.Name)
	}
}
