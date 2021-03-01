package auth

import (
	jwt "github.com/dgrijalva/jwt-go"
)

var (
	secret = "SecretKey01234567890123456789012345678901234567890123456789012345678"
)

func Verify(accessToken string) (bool, string, error) {
	if accessToken == "" {
		return false, "", nil
	}
	token, err := jwt.ParseWithClaims(accessToken, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		data, err := jwt.DecodeSegment(secret)
		return data, err
	})
	if err != nil {
		return false, "", err
	}

	claims, ok := token.Claims.(*jwt.StandardClaims)
	if !ok {
		return false, "", nil
	}
	if err := token.Claims.Valid(); err != nil {
		return false, "", nil
	}
	return true, claims.Subject, nil
}
