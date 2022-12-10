package helper

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewResponse(t *testing.T) {
	expected := Response{
		StatusCode: 200,
		Message:    "ok",
		Data:       nil,
	}

	result := NewResponse(200, "ok", nil)

	assert.Equal(t, expected, result)
}

func TestNewErrorResponse(t *testing.T) {
	expected := ErrorResponse{
		StatusCode: 400,
		Message:    "bad request",
		Errors:     nil,
	}

	result := NewErrorResponse(400, "bad request", nil)

	assert.Equal(t, expected, result)
}
