package controller

type ResCode int

const (
	CodeInvalidParam ResCode = 1000 + iota
	CodeInvalidPassword
	CodeUserExist
	CodeUserNotExist
	CodeServerBusy
	CodeSuccess
	CodeNeedLogin
	CodeInvalidToken
	CodeInvalidParamToken
)

var resCodeMap = map[ResCode]string{
	CodeInvalidParam:      "请求参数错误",
	CodeInvalidPassword:   "用户名或密码错误",
	CodeUserExist:         "用户已存在",
	CodeUserNotExist:      "用户不存在",
	CodeServerBusy:        "服务繁忙",
	CodeSuccess:           "响应成功",
	CodeNeedLogin:         "请重新登录",
	CodeInvalidToken:      "无效的Token",
	CodeInvalidParamToken: "Token参数有误",
}

func (code ResCode) Msg() string {
	msg, ok := resCodeMap[code]
	if !ok {
		return resCodeMap[CodeServerBusy]
	}
	return msg
}
