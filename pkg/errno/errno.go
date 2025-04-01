package errno

import (
	"errors"
	"fmt"
)

type ErrNo struct {
	ErrorCode int64
	ErrorMsg  string
}

func (e ErrNo) Error() string {
	return fmt.Sprintf("[%d] %s", e.ErrorCode, e.ErrorMsg)
}

func NewErrNo(code int64, msg string) ErrNo {
	return ErrNo{
		ErrorCode: code,
		ErrorMsg:  msg,
	}
}
func Errorf(code int64, template string, args ...interface{}) ErrNo {
	return ErrNo{
		ErrorCode: code,
		ErrorMsg:  fmt.Sprintf(template, args...),
	}
}

func (e ErrNo) WithMessage(msg string) ErrNo {
	e.ErrorMsg = msg
	return e
}

func (e ErrNo) WithError(err error) ErrNo {
	e.ErrorMsg = e.ErrorMsg + ": " + err.Error()
	return e
}

func ConvertErr(err error) ErrNo {
	if err == nil {
		return Success
	}
	errno := ErrNo{}
	if errors.As(err, &errno) {
		return errno
	}
	s := InternalServiceError
	s.ErrorMsg = err.Error()
	return s
}
