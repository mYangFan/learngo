package main

import (
	"crypto/rsa"
	"github.com/dgrijalva/jwt-go"
	"time"
)

func main()  {

}

type TokenGenerator interface {
	GenerateToken(accountID string, expire time.Duration)
}


type JWTTokenGen struct {
	privateKey *rsa.PrivateKey
	Issue string
	nowFunc func() time.Time
}

func NewJWTTokenGen(issuer string, privateKey *rsa.PrivateKey) *JWTTokenGen {
	return &JWTTokenGen{
		Issue: issuer,
		nowFunc: func() time.Time {
			return time.Now()
		},
		privateKey: privateKey,
	}
}

func (J *JWTTokenGen) GenerateToken(accountID string, expire time.Duration) (string, error){
	now := J.nowFunc().Unix()
	tkn := jwt.NewWithClaims(jwt.SigningMethodRS512, jwt.StandardClaims{
		Issuer: J.Issue,
		IssuedAt: now,
		ExpiresAt: now + int64(expire.Seconds()),
		Subject: accountID,
	})

	return tkn.SignedString(J.privateKey)
}
