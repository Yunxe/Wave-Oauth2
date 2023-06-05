package service

import (
	"encoding/json"
	"net/http"
	"oauth2/database"
	"oauth2/util"
	"time"

	"github.com/gin-gonic/gin"
)

type CodeParam struct {
	ResponseType string `json:"response_type" form:"response_type" binding:"required"`
	ClientId     string `json:"client_id" form:"client_id" binding:"required"`
	State        string `json:"state" form:"state" binding:"required"`
	RedirectURL  string `json:"redirect_url" form:"redirect_url" binding:"required"`
	Uid          string `json:"uid" form:"uid" binding:"required"`
}

type CodeInfo struct {
	Code     string `json:"code"`
	State    string `json:"state"`
	UserId   string `json:"userId"`
	ClientId string `json:"client_id"`
}

func Code(c *gin.Context) {
	var (
		codeParam *CodeParam
		codeInfo  *CodeInfo
		htmlMap   *HTMLMap
		layout    string = "layout/index.html"
	)

	err := c.ShouldBind(&codeParam)
	if err != nil {
		htmlMap.HtmlReturnError(util.ReqParamInvalidErr)
		c.HTML(util.ErrorMap[util.ReqParamInvalidErr].HttpCode, layout, htmlMap)
		return
	}

	codeInfo = &CodeInfo{
		Code:     util.GetMD5String([]byte(codeParam.ClientId + time.Now().String()))[:8],
		State:    codeParam.State,
		UserId:   codeParam.Uid,
		ClientId: codeParam.ClientId,
	}

	jsonObj, _ := json.Marshal(codeInfo)
	database.RDB.Set(c, codeInfo.Code, jsonObj, 10*time.Minute)
	url := codeParam.RedirectURL + "?code=" + codeInfo.Code + "&state=" + codeParam.State
	c.Redirect(http.StatusFound, url)
}
