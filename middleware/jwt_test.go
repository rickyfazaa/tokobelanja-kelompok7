package middleware

import (
	"strings"
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestValidateToken(t *testing.T) {
	expected := 1
	expectedRole := "admin"

	token, _ := GenerateToken(1, "admin")

	// Test GenerateToken()
	arr := strings.Split(token, ".")
	if len(arr) != 3 {
		t.Errorf("GenerateToken() not valid")
	}

	result, resultRole, _ := ValidateToken(token)

	// Test ValidateToken()
	assert.Equal(t, expected, result)
	assert.Equal(t, expectedRole, resultRole)
}
