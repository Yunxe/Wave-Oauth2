package util

import (
	"errors"

	"github.com/golang-jwt/jwt/v4"
)

var (
	CommonErr = errors.New("common error")

	ReqParamInvalidErr = errors.New("request param invalid error")

	AuthRequire                 = errors.New("authorization required")
	AuthTokenTypeErr            = errors.New("authorization type error")
	AuthParseTokenErr           = errors.New("authorization parse token error")
	AuthTokenExpired            = jwt.ErrTokenExpired
	AuthTokenInvalidIssuer      = jwt.ErrTokenInvalidIssuer
	AuthTokenInvalidInBlackList = errors.New("authorization token invalid in black list")

	ClientNotFound               = errors.New("client not found")
	ClientInfoRedirectUrlInvalid = errors.New("client info redirect url invalid")
	ClientInfoSecretErr          = errors.New("client info secret error")
	CodeNotFound                 = errors.New("code not found")
	TokenTypeErr                 = errors.New("token type error")
)
