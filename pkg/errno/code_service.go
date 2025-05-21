package errno

//业务错误码

//User
const (
	ErrCodeDBerror           = 40001 //数据库错误
	ErrCodeUserExist         = 40002 //用户已存在
	ErrCodeGetUserID         = 40003 //获取用户ID失败
	ErrCodeCreateUser        = 40004 //创建用户失败
	ErrCodeUserNotExist      = 40005 //用户不存在
	ErrCodeImageNotExist     = 40006 //图片不存在
	ErrUserNotHavePermission = 40007 //用户没有权限
	ErrCodePasswordIncorrect = 40008 //密码错误
	ErrInvalidParams         = 40009 //参数错误
	ParamMissingErrorCode    = 40010 //参数缺失
	OSOperateErrorCode       = 40011 //操作系统错误
)
