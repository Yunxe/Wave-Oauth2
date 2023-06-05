package util

import (
	"net/http"
)

type Status struct {
	Code    uint   `json:"code"`
	Message string `json:"message"`
}

type StatusWithData struct {
	Code    uint        `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

type Response struct {
	HttpCode int
	Status   *Status
}

var (
	// SUCCESS 0 通用成功
	SUCCESS = &Response{HttpCode: http.StatusOK, Status: &Status{Code: 0, Message: "成功"}}
)

var ErrorMap = map[error]*Response{

	// 99 通用错误
	CommonErr: {HttpCode: http.StatusInternalServerError, Status: &Status{Code: 99, Message: "服务器内部错误"}},

	// 100xx 基本常用错误码
	ReqParamInvalidErr: {HttpCode: http.StatusBadRequest, Status: &Status{Code: 10001, Message: "请求参数无效"}},

	// 200xx 鉴定权限相关错误
	AuthRequire:                 {HttpCode: http.StatusUnauthorized, Status: &Status{Code: 20001, Message: "需要认证"}},
	AuthTokenTypeErr:            {HttpCode: http.StatusForbidden, Status: &Status{Code: 20002, Message: "认证类型错误"}},
	AuthParseTokenErr:           {HttpCode: http.StatusInternalServerError, Status: &Status{Code: 20003, Message: "解析令牌错误"}},
	AuthTokenExpired:            {HttpCode: http.StatusForbidden, Status: &Status{Code: 20004, Message: "令牌已过期"}},
	AuthTokenInvalidIssuer:      {HttpCode: http.StatusForbidden, Status: &Status{Code: 20005, Message: "令牌签发者无效"}},
	AuthTokenInvalidInBlackList: {HttpCode: http.StatusForbidden, Status: &Status{Code: 20006, Message: "令牌无效,已在黑名单"}},

	ClientNotFound:               {HttpCode: http.StatusForbidden, Status: &Status{Code: 30201, Message: "无此客户端信息,请进行认证"}},
	ClientInfoRedirectUrlInvalid: {HttpCode: http.StatusForbidden, Status: &Status{Code: 30202, Message: "客户端信息错误,回调地址无效"}},
	ClientInfoSecretErr:          {HttpCode: http.StatusForbidden, Status: &Status{Code: 30203, Message: "客户端信息错误,密码错误"}},
	CodeNotFound:                 {http.StatusInternalServerError, &Status{Code: 30204, Message: "无授权码信息，请重新授权"}},
	TokenTypeErr:                 {HttpCode: http.StatusForbidden, Status: &Status{Code: 30205, Message: "令牌类型错误"}},
}
