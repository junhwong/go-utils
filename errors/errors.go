package errors

import (
	"fmt"
	"io"
	"net"
	"runtime"
	"runtime/debug"
)

const (
	NoError      = ""
	ErrUnknown   = "unknown-error"
	ErrIOTimeout = "io.timeout"
	ErrIOClosed  = "io.closed"
	ErrIOEOF     = "io.EOF"
)

type TraceError struct {
	Raise  error
	code   string
	format string
	args   []interface{}
	File   string
	Method string
	Line   int
	stack  []byte
}

func (err *TraceError) fillSource() {
	pc, file, line, ok := runtime.Caller(3)
	if ok {
		err.File = file
		err.Line = line
		err.Method = runtime.FuncForPC(pc).Name()
	}
}

func (err *TraceError) Tracing() *TraceError {
	err.stack = debug.Stack()
	return err
}

func (err *TraceError) Code() string {
	return err.code
}

func (err *TraceError) Error() string {
	return fmt.Sprintf(err.format, err.args...)
}

// func (err *TraceError) Format(f fmt.State, c rune) {
// 	x := errors.Formatter{}
// }

func New(code, format string, args ...interface{}) *TraceError {
	err := &TraceError{
		code:   code,
		format: format,
		args:   args,
	}
	err.fillSource()
	return err
}

func Raise(raise error, code, format string, args ...interface{}) *TraceError {
	err := &TraceError{
		Raise:  raise,
		code:   code,
		format: format,
		args:   args,
	}
	err.fillSource()
	return err
}

func Wrap(raise error) *TraceError {
	err := &TraceError{
		Raise: raise,
	}
	err.fillSource()
	return err
}

func Code(err error) string {
	if err == nil {
		return NoError
	} else if err == io.EOF {
		return ErrIOEOF
	}

	var code string
	switch x := err.(type) {
	case *TraceError:
		code = x.Code()
	case *net.OpError:
		if x.Timeout() {
			code = ErrIOTimeout
		}
		if x.Error() == "use of closed network connection" {
			code = ErrIOClosed
		}
	}
	if code == "" {
		code = err.Error()
	}
	if code == "" {
		code = ErrUnknown
	}
	return code
}
