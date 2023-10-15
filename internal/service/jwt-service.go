package service

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func GenerateTokenPair(userId uint) (string, string, error) {
	tokenLifespan, err := strconv.Atoi("1")
	if err != nil {
		return "", "", err
	}

	// Generate token
	tokenClaims := getTokenClaims(userId, tokenLifespan)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, tokenClaims)
	tokenString, err := token.SignedString([]byte("secret_no_sifat"))
	if err != nil {
		return "", "", err
	}

	// Generate refresh token
	refreshTokenLifespan, err := strconv.Atoi("2")
	if err != nil {
		return "", "", err
	}
	refreshTokenClaims := getTokenClaims(userId, refreshTokenLifespan)
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshTokenClaims)
	refreshTokenString, err := refreshToken.SignedString([]byte("secret_no_sifat"))
	if err != nil {
		return "", "", err
	}
	return tokenString, refreshTokenString, nil
}

func getTokenClaims(userId uint, tokenLifespan int) jwt.MapClaims {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["user_id"] = userId
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(tokenLifespan)).Unix()
	return claims
}

// TokenValid Check token valid or not
func TokenValid(ctx *gin.Context) error {
	tokenString := ExtractToken(ctx)
	_, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return []byte("secret_no_sifat"), nil
	})
	if err != nil {
		return err
	}
	return nil
}

// ExtractToken Extract token
func ExtractToken(ctx *gin.Context) string {
	token := ctx.Query("token")
	if token != "" {
		return token
	}
	bearerToken := ctx.Request.Header.Get("Authorization")
	if len(strings.Split(bearerToken, " ")) == 2 {
		return strings.Split(bearerToken, " ")[1]
	}
	return ""
}

// ExtractTokenID Extract User Id from token.
func ExtractTokenID(c *gin.Context) (uint, error) {

	tokenString := ExtractToken(c)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte("secret_no_sifat"), nil
	})
	if err != nil {
		return 0, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		uid, err := strconv.ParseUint(fmt.Sprintf("%.0f", claims["user_id"]), 10, 32)
		if err != nil {
			return 0, err
		}
		return uint(uid), nil
	}
	return 0, nil
}
