package util

import (
	"github.com/golang-jwt/jwt/v4"
	"os"
	"time"
)

type TokenInfo struct {
	Token     string
	TokenType string
	ExpiresIn time.Duration
}

type UserClaims struct {
	Uid  string `json:"uid"`
	Type string `json:"type"`
	jwt.RegisteredClaims
}

func (c *UserClaims) CreateRefreshToken(clientId string) (string, error) {
	var (
		expiresTime *jwt.NumericDate
		dur         time.Duration
		signKey     = os.Getenv("SIGNING_KEY_REFRESH")
	)

	dur, _ = time.ParseDuration("24h")
	expiresTime = jwt.NewNumericDate(time.Now().Add(dur))

	c.RegisteredClaims = jwt.RegisteredClaims{
		Issuer:    "Wave-oauth2",
		Audience:  []string{"client-" + clientId},
		ExpiresAt: expiresTime,
		IssuedAt:  jwt.NewNumericDate(time.Now()),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	signedToken, err := token.SignedString([]byte(signKey))
	return signedToken, err
}
