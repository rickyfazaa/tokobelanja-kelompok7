package middleware

import (
	"net/http"
	"strings"
	"tokobelanja-kelompok7/helper"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")

	if !strings.Contains(authHeader, "Bearer") {
		c.AbortWithStatusJSON(
			http.StatusUnauthorized,
			helper.NewErrorResponse(
				http.StatusUnauthorized,
				"Unauthenticated",
				nil,
			),
		)
		return
	}

	tokenType, token := strings.Split(authHeader, " ")[0], strings.Split(authHeader, " ")[1]
	if tokenType == "" || tokenType != "Bearer" {
		c.AbortWithStatusJSON(
			http.StatusUnauthorized,
			helper.NewErrorResponse(
				http.StatusUnauthorized,
				"Unauthenticated",
				nil,
			),
		)
		return
	}

	if token == "" {
		c.AbortWithStatusJSON(
			http.StatusUnauthorized,
			helper.NewErrorResponse(
				http.StatusUnauthorized,
				"Unauthenticated",
				nil,
			),
		)
		return
	}

	userID, userRole, err := ValidateToken(token)
	if err != nil {
		c.AbortWithStatusJSON(
			http.StatusUnauthorized,
			helper.NewErrorResponse(
				http.StatusUnauthorized,
				"Unauthenticated",
				err.Error(),
			),
		)
		return
	}

	c.Set("currentUser", userID)
	c.Set("roleUser", userRole)
}
