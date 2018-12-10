package utils

import (
	"bytes"
	"fmt"
	"strings"
)

// ExceptionDebug 表示全局默认是否是调试状态，默认：true
var ExceptionDebug = true

type Error interface {
	error
	Debug() bool
	SetDebug(b bool) Error
	Code() string
	Message() string
}

// Exception 封装一个带有层级和 Code 的错误。
type Exception struct {
	inner Error
	code  string
	msg   interface{}
	debug int
}

func (e *Exception) SetDebug(b bool) Error {
	if !b {
		e.debug = -1
	} else {
		e.debug = 1
	}
	return e
}

func (e *Exception) Debug() bool {
	if e.debug < 0 {
		return false
	} else if e.debug > 0 {
		return true
	}
	return ExceptionDebug
}

func (e *Exception) Code() string {
	return e.code
}

func (e *Exception) Error() string {
	msg := "code: " + e.code
	if e.msg != nil {
		msg += " msg: " + e.Message()
	}
	if e.inner != nil {
		msg = e.inner.Error() + "\n" + msg
	}
	return msg
}

func (e *Exception) Message() string {
	msg := ""
	if e.msg != nil {
		if err, ok := e.msg.(error); ok {
			msg += err.Error()
		} else {
			msg += fmt.Sprintf("%v", e.msg)
		}
	} else if e.inner != nil {
		msg = e.inner.Error() + "\n" + msg
	}
	return msg
}

func (e *Exception) MarshalJSON() ([]byte, error) {
	code := e.Code()
	msg := strings.Replace(e.Message(), `"`, `\"`, -1)
	buf := new(bytes.Buffer)
	buf.WriteString("{")
	if e.Debug() {
		buf.WriteString(`"`)
		buf.WriteString(`error`)
		buf.WriteString(`":`)
		buf.WriteString(`"`)
		buf.WriteString(code)
		buf.WriteString(`",`)
		buf.WriteString(`"`)
		buf.WriteString(`errorMsg`)
		buf.WriteString(`":`)
		buf.WriteString(`"`)
		buf.WriteString(strings.Replace(msg, `"`, `\"`, -1))
		buf.WriteString(`"`)
	} else {
		if code == "" {
			code = "INNER.ERROR"
		}
		buf.WriteString(`"`)
		buf.WriteString(`error`)
		buf.WriteString(`":`)
		buf.WriteString(`"`)
		buf.WriteString(code)
		buf.WriteString(`"`)
	}
	buf.WriteString("}")
	return buf.Bytes(), nil
}

// func (e *Exception) code(inner interface{}) string {
// 	c := e.Code
// 	if c == "" && inner != nil {
// 		if err, ok := inner.(*Exception); ok {
// 			c = err.code(err.Err)
// 		} else if err, ok := inner.(Exception); ok {
// 			c = err.code(err.Err)
// 		}
// 	}
// 	return c
// }
// func (e *Exception) msg(inner interface{}) string {
// 	c := ""
// 	if inner != nil {
// 		if err, ok := inner.(*Exception); ok {
// 			c = err.msg(err.Err)
// 		} else if err, ok := inner.(Exception); ok {
// 			c = err.msg(err.Err)
// 		} else if err, ok := inner.(error); ok {
// 			c = err.Error()
// 		} else {
// 			c = fmt.Sprintf("%v", inner)
// 		}
// 	} else {
// 		c = e.code(inner)
// 	}
// 	return c
// }

func Err(code string, innerOrMsg ...interface{}) *Exception {
	ret := &Exception{
		code: code,
	}
	if len(innerOrMsg) == 1 {
		m := innerOrMsg[0]
		if err, ok := m.(Error); ok {
			if ret.code == "" {
				ret.code = err.Code()
			}
			ret.inner = err
		} else {
			ret.msg = m
		}
	} else {
		ret.msg = fmt.Sprint(innerOrMsg...)
	}
	if ret.code == "" {
		panic("无效的参数: code")
	}

	return ret
}
