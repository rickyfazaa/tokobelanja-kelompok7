package helper

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	err = errors.New("you are not admin")
)

func TestGetErrorData(t *testing.T) {
	expected := "you are not admin"

	result := GetErrorData(err)

	assert.Equal(t, expected, result)
}
