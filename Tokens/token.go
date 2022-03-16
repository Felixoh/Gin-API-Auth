package Tokens

import (
	"fmt"
	"time"

	"github.com/Felixoh/models"
	jwt "github.com/dgrijalva/jwt-go"
)

const (
	jWTPrivateToken = "SecretTokenSecretToken"
	ip              = "192.168.0.107"
)

func GenerateToken(claims *models.JwtClaims, expirationTime time.Time) (string, error) {

	claims.ExpiresAt = expirationTime.Unix()
	claims.IssuedAt = time.Now().UTC().Unix()
	claims.Issuer = ip

	//pass in the algorithm you want to use to generate the token:
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	//sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(jWTPrivateToken))

	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func VerifyToken(tokenString, origin string) (bool, *models.JwtClaims) {
	claims := &models.JwtClaims{}

	token, _ := getTokenFromString(tokenString, claims)
	if token.Valid {
		if e := claims.Valid(); e == nil {
			return true, claims
		}
	}
	return false, claims
}

func getTokenFromString(tokenString string, claims *models.JwtClaims) (*jwt.Token, error) {
	return jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {

		// Don't forget to validate the alg if its  what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		//hmacSampleSecret is a []byte containing your secret, e.g []byte("my_secret_key")
		return []byte(jWTPrivateToken), nil
	})
}

func GetClaims(tokenString string) models.JwtClaims {
	claims := &models.JwtClaims{}
	_, err := getTokenFromString(tokenString, claims)
	if err == nil {
		return *claims
	}
	return *claims
}
