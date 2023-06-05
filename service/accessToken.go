package service

import (
	"context"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v9"
	"oauth2/database"
	"oauth2/model"
	"oauth2/proto/rpc"
	pb "oauth2/proto/sso_client"
	"oauth2/util"
	"time"
)

type AccessTokenParam struct {
	ClientId     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	Code         string `json:"code"`
	GrantType    string `json:"grant_type"`
}

type AccessTokenRes struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token,omitempty"`
	ExpiresIn    int64  `json:"expires_in"`
}

func AccessToken(c *gin.Context) (data any, err error) {
	var (
		accessTokenParam AccessTokenParam
		client           *model.Client
		codeInfo         CodeInfo
	)
	err = c.ShouldBind(&accessTokenParam)
	if err != nil || accessTokenParam.GrantType != "authorization_code" {
		return nil, util.ReqParamInvalidErr
	}

	database.DB.Where("client_id = ?", accessTokenParam.ClientId).First(&client)
	if client.ClientSecret != accessTokenParam.ClientSecret {
		return nil, util.ClientInfoSecretErr
	}

	val, err := database.RDB.Get(c, accessTokenParam.Code).Result()
	if err == redis.Nil {
		return nil, util.CodeNotFound
	}
	err = json.Unmarshal([]byte(val), &codeInfo)
	if err != nil {
		return nil, err
	}

	//protobuf请求sso返回access_token && 自己返回refresh_token
	ctx, cancel := context.WithTimeout(c, 3*time.Second)
	defer cancel()
	rpcRes, err := rpc.Grpc.GetExternalToken(ctx, &pb.TokenRequest{
		UserId:   codeInfo.UserId,
		ClientId: accessTokenParam.ClientId,
	})
	if err != nil {
		return nil, err
	}

	claims := &util.UserClaims{
		Uid:  codeInfo.UserId,
		Type: "refresh_token",
	}
	refreshToken, err := claims.CreateRefreshToken(accessTokenParam.ClientId)
	if err != nil {
		return nil, err
	}

	//TODO Cache-Control: no-store;Pragma: no-cache

	return &util.StatusWithData{
		Code:    0,
		Message: "成功",
		Data: &AccessTokenRes{
			AccessToken:  rpcRes.Token,
			RefreshToken: refreshToken,
			ExpiresIn:    rpcRes.ExpiresIn,
		},
	}, nil

}
