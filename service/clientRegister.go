package service

import (
	"fmt"
	"oauth2/database"
	"oauth2/model"
	"oauth2/util"

	"github.com/gin-gonic/gin"
)

type ClientRegisterParam struct {
	ClientName  string `json:"client_name" form:"client_name" binding:"required"`
	HomepageURL string `json:"homepage_url" form:"homepage_url" binding:"required,url"`
	RedirectURL string `json:"redirect_url" form:"redirect_url" binding:"required,url"`
	Description string `json:"description" form:"description"`
}

func ClientRegister(c *gin.Context) (data any, err error) {
	var (
		clientRegisterParam ClientRegisterParam
	)
	if err = c.ShouldBind(&clientRegisterParam); err != nil {
		return nil, util.ReqParamInvalidErr
	}
	fmt.Println(clientRegisterParam)
	client := &model.Client{
		ClientName:   clientRegisterParam.ClientName,
		ClientId:     util.GetMD5String([]byte(clientRegisterParam.HomepageURL))[:6],
		ClientSecret: util.GetSHAStringPass([]byte(clientRegisterParam.HomepageURL)),
		HomepageURL:  clientRegisterParam.HomepageURL,
		RedirectURL:  clientRegisterParam.RedirectURL,
		Description:  clientRegisterParam.Description,
	}

	res := database.DB.Create(&client)
	if res.Error != nil {
		return res.Error, nil
	}

	//TODO 应该做一个客户端审核

	return &util.StatusWithData{
		Code:    0,
		Message: "成功",
		Data: struct {
			ClientId     string `json:"client_id"`
			ClientSecret string `json:"client_secret"`
		}{
			client.ClientId, client.ClientSecret,
		},
	}, nil
}
