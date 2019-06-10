package e

import "fmt"

type ECODE int

const (
	E_SUCCESS = 0
	E_UNKNOWN = 500

	// 内部错误
	E_CRON          = 900 // cron进程错误
	E_RPC           = 901 // rpc进程错误
	E_WEBSOCKET     = 902 // websocket进程错误
	E_POLL          = 903 // poll进程错误
	E_PROTO         = 904 // proto协议错误
	E_JSON          = 905 // json解析错误
	E_UUID          = 906 // uuid生成错误
	E_TYPE_ASSERT   = 907 // 类型断言错误
	E_REPEAT_REGIST = 908 // 重复注册

	// 错误
	E_INVALID_PARAM     = 1000
	E_SERVICE_ERROR     = 1001
	E_DB_ERROR          = 1002
	E_INVALID_CODE      = 1003
	E_AUTH_FAIL         = 1004
	E_AUTH_TIMEOUT      = 1005
	E_ROLE_EXIST        = 1006
	E_CACHE             = 1007
	E_CREATE_ROLE       = 1008
	E_ALREADY_ONLINE    = 1009 // 重复登录
	E_GOODS_NOT_EXIST   = 1010 // 道具不存在
	E_GOODS_NOT_ENOUGH  = 1011 // 道具不够
	E_ANOTHER_ONLINE    = 1012 // 异地登录
	E_ROLE_NOT_EXIST    = 1013 // 用户不存在
	E_SERVER_SHUTDOWN   = 1014 // 服务器关机准备
	E_SMS               = 1015 // 短信错误
	E_CAPTCHA_ERROR     = 1016 // 验证码错误
	E_REGIST_ERROR      = 1017 // 注册失败
	E_ACCOUNT_NOT_EXIST = 1018 // 账号不存在
	E_AICAPTCHA_ERROR   = 1019 // 人机验证失败
	E_INVALID_VERSION   = 1020 // 版本号不匹配
)

const (
	E_MSG_SUCCESS = "ok"
	E_MSG_UNKNOWN = "unknown fail"

	// 内部错误
	E_MSG_PROTO         = "proto匹配失败"
	E_MSG_RPC           = "rpc请求错误"
	E_MSG_JSON          = "json解析错误"
	E_MSG_UUID          = "uuid生成错误"
	E_MSG_TYPE_ASSERT   = "类型断言错误"
	E_MSG_REPEAT_REGIST = "重复注册"

	// 错误
	E_MSG_INVALID_PARAM     = "参数错误"
	E_MSG_SERVICE_ERROR     = "服务错误"
	E_MSG_DB_ERROR          = "DB错误"
	E_MSG_INVALID_CODE      = "代码错误"
	E_MSG_AUTH_FAIL         = "验证失败"
	E_MSG_AUTH_TIMEOUT      = "验证超时"
	E_MSG_ROLE_EXIST        = "用户已存在"
	E_MSG_CACHE             = "内部缓存错误"
	E_MSG_CREATE_ROLE       = "创建账号失败"
	E_MSG_ALREADY_ONLINE    = "重复登录"
	E_MSG_GOODS_NOT_EXIST   = "道具不存在"
	E_MSG_GOODS_NOT_ENOUGH  = "道具不够"
	E_MSG_ANOTHER_ONLINE    = "异地登录"
	E_MSG_ROLE_NOT_EXIST    = "用户不存在"
	E_MSG_SERVER_SHUTDOWN   = "服务器关机准备"
	E_MSG_SMS               = "短信服务错误"
	E_MSG_CAPTCHA_ERROR     = "验证码错误"
	E_MSG_REGIST_ERROR      = "注册失败"
	E_MSG_ACCOUNT_NOT_EXIST = "账号不存在"
	E_MSG_AICAPTCHA_ERROR   = "人机验证失败"
	E_MSG_INVALID_VERSION   = "版本号不匹配"
)

var (
	ErrSuccess = StandardError{E_SUCCESS, E_MSG_SUCCESS}
	ErrUnknown = StandardError{E_UNKNOWN, E_MSG_UNKNOWN}

	// 内部错误
	ErrProto        = StandardError{E_PROTO, E_MSG_PROTO}
	ErrRpc          = StandardError{E_RPC, E_MSG_RPC}
	ErrJson         = StandardError{E_JSON, E_MSG_JSON}
	ErrUUID         = StandardError{E_UUID, E_MSG_UUID}
	ErrTypeAssert   = StandardError{E_TYPE_ASSERT, E_MSG_TYPE_ASSERT}
	ErrRepeatRegist = StandardError{E_REPEAT_REGIST, E_MSG_REPEAT_REGIST}

	// 前端错误
	ErrInvalidParam    = StandardError{E_INVALID_PARAM, E_MSG_INVALID_PARAM}
	ErrServiceError    = StandardError{E_SERVICE_ERROR, E_MSG_SERVICE_ERROR}
	ErrDBError         = StandardError{E_DB_ERROR, E_MSG_DB_ERROR}
	ErrInvalidCode     = StandardError{E_INVALID_CODE, E_MSG_INVALID_CODE}
	ErrAuthFail        = StandardError{E_AUTH_FAIL, E_MSG_AUTH_FAIL}
	ErrAuthTimeout     = StandardError{E_AUTH_TIMEOUT, E_MSG_AUTH_TIMEOUT}
	ErrRoleExist       = StandardError{E_ROLE_EXIST, E_MSG_ROLE_EXIST}
	ErrCache           = StandardError{E_CACHE, E_MSG_CACHE}
	ErrCreateRole      = StandardError{E_CREATE_ROLE, E_MSG_CREATE_ROLE}
	ErrAlreadyOnline   = StandardError{E_ALREADY_ONLINE, E_MSG_ALREADY_ONLINE}
	ErrGoodsNotExist   = StandardError{E_GOODS_NOT_EXIST, E_MSG_GOODS_NOT_EXIST}
	ErrGoodsNotEnough  = StandardError{E_GOODS_NOT_ENOUGH, E_MSG_GOODS_NOT_ENOUGH}
	ErrAnotherOnline   = StandardError{E_ANOTHER_ONLINE, E_MSG_ANOTHER_ONLINE}
	ErrRoleNotExist    = StandardError{E_ROLE_NOT_EXIST, E_MSG_ROLE_NOT_EXIST}
	ErrServerShutdown  = StandardError{E_SERVER_SHUTDOWN, E_MSG_SERVER_SHUTDOWN}
	ErrSms             = StandardError{E_SMS, E_MSG_SMS}
	ErrCaptcha         = StandardError{E_CAPTCHA_ERROR, E_MSG_CAPTCHA_ERROR}
	ErrRegist          = StandardError{E_REGIST_ERROR, E_MSG_REGIST_ERROR}
	ErrAccountNotExist = StandardError{E_ACCOUNT_NOT_EXIST, E_MSG_ACCOUNT_NOT_EXIST}
	ErrAICaptchaError  = StandardError{E_AICAPTCHA_ERROR, E_MSG_AICAPTCHA_ERROR}
	ErrInvalidVersion  = StandardError{E_INVALID_VERSION, E_MSG_INVALID_VERSION}
)

type StandardError struct {
	ErrCode uint32 `json:"errCode"`
	ErrMsg  string `json:"errMsg"`
}

func (e StandardError) Error() string {
	return fmt.Sprintf("err code: %d msg: %s\n", e.ErrCode, e.ErrMsg)
}
