package authJwt

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

//TODO: change secret key
var secretKey = []byte("secretSauce")

func generateTokenExpiration() int64 {
	expHours, err := strconv.Atoi(os.Getenv("TOKEN_EXPIRATION_HOURS"))
    if err != nil {
        expHours = 12
    }
    return time.Now().Add(time.Hour * time.Duration(expHours)).Unix()
}

func CreateToken(username string, role string) (string, error) {
	exp := generateTokenExpiration()
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, jwt.MapClaims{
		"username": username,
		"role": role,
		"exp": exp,
	})

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "There was an error generating the token", err
		}

	return tokenString, nil
}

func WithValidMethods(methods []string) jwt.ParserOption {
	return jwt.WithValidMethods(methods)
}

func VerifyToken(tokenString string) error {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	}, WithValidMethods([]string{"HS512"})) 

	if err != nil {
		return err
	}

	if !token.Valid {
		return fmt.Errorf("invalid token")
	}

	return nil
}

