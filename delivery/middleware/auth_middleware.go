package middleware

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/kelompok-2/ilmu-padi/repository"
)

var jwtKey = []byte("my_secret_key")

type Claims struct {
	UserID uint `json:"user_id"`
	jwt.StandardClaims
}

func GenerateJWT(userID uint) (string, error) {
	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &Claims{
		UserID: userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}

func JWTAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {

			tokenString, _ = c.Cookie("jwt")
			if tokenString == "" {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header required"})
				c.Abort()
				return
			}
		} else {
			tokenString = strings.TrimPrefix(tokenString, "Bearer ")
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte("YOUR_SECRET_KEY"), nil
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		userID := uint(claims["user_id"].(float64))
		c.Set("user_id", userID)
		c.Next()
	}
}

// func RoleMiddleware(roles ...string) gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		userID, exists := c.Get("user_id")
// 		if !exists {
// 			c.JSON(http.StatusUnauthorized, gin.H{"error": "User ID not found"})
// 			c.Abort()
// 			return
// 		}

// 		userRepo := c.MustGet("userRepo").(*repository.UserRepository)
// 		user, err := userRepo.FindByID(userID.(uint))
// 		if err != nil {
// 			c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
// 			c.Abort()
// 			return
// 		}

// 		// Assuming user has a field Roles which is a slice of strings
// 		userRoles := map[string]struct{}{}
// 		for _, role := range user.Roles {
// 			userRoles[role] = struct{}{}
// 		}

// 		authorized := false
// 		for _, role := range roles {
// 			if _, ok := userRoles[role]; ok {
// 				authorized = true
// 				break
// 			}
// 		}

// 		if !authorized {
// 			c.JSON(http.StatusForbidden, gin.H{"error": "Forbidden"})
// 			c.Abort()
// 			return
// 		}

// 		c.Next()
// 	}
// }

func RoleMiddleware(allowedRoles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := c.MustGet("user_id").(uint)
		userRepo := repository.NewUserRepository(nil) // Pass actual DB connection

		roles, err := userRepo.GetRolesByUserID(userID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
			c.Abort()
			return
		}

		roleMap := make(map[string]bool)
		for _, role := range roles {
			roleMap[role.Name] = true
		}

		for _, allowedRole := range allowedRoles {
			if roleMap[allowedRole] {
				c.Next()
				return
			}
		}

		c.JSON(http.StatusForbidden, gin.H{"error": "Forbidden"})
		c.Abort()
	}
}
