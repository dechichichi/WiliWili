package errno

//常用错误

var (
	Success = NewErrNo(SuccessCode, "ok")
	InternalServiceError = NewErrNo(500, "Internal Service Error")
)

