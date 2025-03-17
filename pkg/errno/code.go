package errno

//服务端错误码
const (
	SuccessCode = 10000
	SuccessMsg  = "success"
)

//config错误
const(
	ErrConfigParseCode = 20001 //config解析失败
	ErrConfigEmptyCode = 20002 //config为空
	RedisConnectFailed = 20003 //redis连接失败
)



//鉴权错误
const (
	ErrTokenGenCode       = 30001 //token生成失败
	ErrTokenExpiredCode   = 30002 //token过期
	ErrTokenInvalidCode   = 30003 //token无效
	ErrTokenParseCode     = 30004 //token解析失败
	ErrTokenVerifyCode    = 30005 //token校验失败
)

