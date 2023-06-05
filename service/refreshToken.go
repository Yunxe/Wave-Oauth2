package service

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"oauth2/proto/rpc"
	pb "oauth2/proto/sso_client"
	"oauth2/util"
	"os"
	"strings"
	"time"
)

type RefreshTokenParam struct {
	Authorization string `header:"Authorization" binding:"required"`
}

func RefreshToken(c *gin.Context) (data any, err error) {
	var (
		h       *RefreshTokenParam
		arr     []string
		signKey = os.Getenv("SIGNING_KEY_REFRESH")
	)
	if err := c.ShouldBindHeader(&h); err != nil {
		return nil, util.AuthRequire
	}

	if arr = strings.Fields(h.Authorization); strings.ToLower(arr[0]) != "bearer" {
		return nil, util.AuthTokenTypeErr
	}
	token, err := jwt.ParseWithClaims(arr[1], &util.UserClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(signKey), nil
	})

	claims, ok := token.Claims.(*util.UserClaims)
	if !ok || !token.Valid {
		return nil, err
	}
	if claims.Type != "refresh_token" {
		return nil, util.TokenTypeErr
	}
	audience := claims.Audience
	slice := strings.Split(audience[0], "-")
	ctx, cancel := context.WithTimeout(c, 3*time.Second)
	defer cancel()
	rpcRes, err := rpc.Grpc.GetExternalToken(ctx, &pb.TokenRequest{
		UserId:   claims.Uid,
		ClientId: slice[1],
	})
	if err != nil {
		return nil, err
	}
	return &util.StatusWithData{
		Code:    0,
		Message: "成功",
		Data: &AccessTokenRes{
			AccessToken: rpcRes.Token,
			ExpiresIn:   rpcRes.ExpiresIn,
		},
	}, nil

}
