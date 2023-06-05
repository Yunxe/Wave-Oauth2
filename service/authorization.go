package service

import (
	"net/http"
	"oauth2/database"
	"oauth2/model"
	"oauth2/util"

	"github.com/gin-gonic/gin"
)

type AuthorizationParam struct {
	ResponseType string `json:"response_type" form:"response_type" binding:"required"`
	ClientId     string `json:"client_id" form:"client_id" binding:"required"`
	State        string `json:"state" form:"state" binding:"required"`
	RedirectURL  string `json:"redirect_url" form:"redirect_url" binding:"required"`
}

type HTMLMap struct {
	Title      string       `json:"title"`
	HasError   bool         `json:"hasError"`
	Error      *util.Status `json:"error"`
	ClientName string       `json:"client_name"`
}

func (htmlMap *HTMLMap) Default(hasError bool) {
	if hasError {
		htmlMap.Title = "error"
		htmlMap.HasError = true
	} else {
		htmlMap.Title = "auth"
		htmlMap.HasError = false
		htmlMap.Error = nil
	}
}

func (htmlMap *HTMLMap) HtmlReturnError(err error) {
	htmlMap.Default(true)
	htmlMap.Error = util.ErrorMap[err].Status
}

func Authorization(c *gin.Context) {
	var (
		authorizationParam AuthorizationParam
		client             *model.Client
		htmlMap            HTMLMap
		layout             string = "layout/index.html"
	)
	// var client model.Client
	err := c.ShouldBind(&authorizationParam)
	if err != nil || authorizationParam.ResponseType != "code" {
		htmlMap.HtmlReturnError(util.ReqParamInvalidErr)
		c.HTML(util.ErrorMap[util.ReqParamInvalidErr].HttpCode, layout, htmlMap)
		return
	}

	database.DB.Where("client_id = ?", authorizationParam.ClientId).First(&client)

	if *client == *model.NewClient() {
		htmlMap.HtmlReturnError(util.ClientNotFound)
		c.HTML(util.ErrorMap[util.ClientNotFound].HttpCode, layout, htmlMap)
		return
	}
	if client.RedirectURL != authorizationParam.RedirectURL {
		htmlMap.HtmlReturnError(util.ClientInfoRedirectUrlInvalid)
		c.HTML(util.ErrorMap[util.ClientInfoRedirectUrlInvalid].HttpCode, layout, htmlMap)
		return
	}

	htmlMap.Default(false)
	htmlMap.ClientName = client.ClientName
	c.HTML(http.StatusOK, layout, htmlMap)
}
