package middleware

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
)

func TestAuthMiddleware(t *testing.T) {
	gin.SetMode(gin.TestMode)

	token, _ := GenerateToken(1, "admin")

	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)
	ctx.Request = &http.Request{
		Header: make(http.Header),
		URL:    &url.URL{},
	}
	ctx.Request.Header.Set("Authorization", "Bearer "+token)

	AuthMiddleware(ctx)

	assert.Equal(t, 1, ctx.MustGet("currentUser").(int))
	assert.Equal(t, "admin", ctx.MustGet("roleUser").(string))
}
