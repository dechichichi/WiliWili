package errno

//业务错误码

//User
const (
	ErrCodeDBerror           = 40001  //数据库错误
	ErrCodeUserExist         = 40002  //用户已存在
	ErrCodeGetUserID         = 40003  //获取用户ID失败
	ErrCodeCreateUser        = 40004  //创建用户失败
	ErrCodeUserNotExist      = 40005  //用户不存在
	ErrUserNotHavePermission = 40006  //用户没有权限
	ErrCodePasswordIncorrect = 40007  //密码错误
	ErrInvalidParams         = 40008  //参数错误
	ParamMissingErrorCode    = 40009  //参数缺失
	OSOperateErrorCode       = 400010 //操作系统错误
)
