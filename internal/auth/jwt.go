package authJwt

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"goserver.com/internal/utils"
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
	secretKey := []byte(utils.GetEnv("JWT_SECRET_KEY", "default_secret"))
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

func VerifyToken(tokenStr string) (jwt.MapClaims, error) {
	secretKey := []byte(utils.GetEnv("JWT_SECRET_KEY", "default_secret")) // Consistent secret key

	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	}, jwt.WithValidMethods([]string{jwt.SigningMethodHS512.Alg()}))
	if err != nil || !token.Valid {
		fmt.Printf("VerifyToken: Invalid token: %v\n", err)
		return nil, errors.New("invalid token")
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		fmt.Println("VerifyToken: Invalid claims")
		return nil, errors.New("invalid claims")
	}
	return claims, nil
}

