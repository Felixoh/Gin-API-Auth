package models

import (
	"fmt"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

type JwtClaims struct {
	CompanyId string `json:"companyId,omitempty"`
	Username  string `json:"username,omitempty"`
	Roles     []int  `json:"roles,omitempty"`

	jwt.StandardClaims
}

const ip = "192.168.0.107"

//implementing claims interface.
func (claims JwtClaims) Valid() error {
	var now = time.Now().UTC().Unix()
	if claims.VerifyExpiresAt(now, true) && claims.VerifyIssuer(ip, true) {
		return nil
	}
	return fmt.Errorf("Token is invalid")
}
