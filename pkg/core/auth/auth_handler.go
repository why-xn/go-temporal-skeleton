package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/why-xn/go-temporal-skeleton/pkg/core/context"
	. "github.com/why-xn/go-temporal-skeleton/pkg/types"
	"net/http"
)

func TokenAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		usr := GetUserFromToken(token)
		if usr != nil {
			context.AddUserToContext(c, usr)
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"status": "401", "msg": "Authentication required"})
			c.Abort()
			return
		}
		c.Next()
	}
}

// Returning dummy auth if any token provided
func GetUserFromToken(token string) *User {
	if len(token) > 0 {
		return &User{
			Username: "shihab",
			Password: "-",
			Email:    "shihabhasan.official@gmail.com",
		}
	}
	return nil
}
