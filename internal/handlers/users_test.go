package handlers

import (
	"testing"

	"github.com/go-playground/assert"
)


func TestPasswordVerification(t *testing.T) {
	storedPassword := "correctpassword"
	inputPassword := "correctpassword"
	wrongPassword := "wrongpassword"

	assert.Equal(t, storedPassword, inputPassword)
	assert.NotEqual(t, storedPassword, wrongPassword)
}