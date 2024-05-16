package authJwt

import (
	"testing"

	"github.com/golang-jwt/jwt/v5"
)

//-------------------------- UNIT TESTS-------------------------------
func TestCreateTokenSuccess(t *testing.T) {
	username := "testuser"
	role := "admin"
	tokenString, err := CreateToken(username, role)
	if err != nil {
		t.Errorf("CreateToken() error = %v, wantErr %v", err, nil)
	}

	// Parse the token to check if the claims are correct
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if claims["username"] != username || claims["role"] != role {
			t.Errorf("Claims do not match expected values")
		}
	} else {
		t.Errorf("Failed to parse token or token is invalid")
	}
}
