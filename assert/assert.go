package assert

import (
	"fmt"
	"os"
	"reflect"
	"runtime/debug"
)

// DEBUG 指示所有断言是否开启。默认：true
//
// 推荐：在编译时使用 `go build -ldflags` 的 `-x github.com/junhwong/go-logs/assert.DEBUG=false` 来关闭它。
//
// 建议：如果你的断言语句复杂耗时可以使用下面的添加语句来优化它：
//
// if assert.DEBUG {
//
// ... // your code here
//
// }
//
var DEBUG = true

// Must 断言参数 `b` 必须是 `true` 否则 panic。
func Must(b bool, format string, args ...interface{}) {
	if !DEBUG {
		return
	}
	if !b {
		panic(fmt.Errorf(format, args...))
	}
}

// Assert 断言参数 `b` 必须是 `true` 否则 `panic`。
//
// assert.Assert(1==1,"...")
//
func Assert(b bool, format string, args ...interface{}) {
	if !DEBUG {
		return
	}
	if !b {
		fmt.Fprintf(os.Stderr, "=== Assert FAIL: "+format+"\n", args...)
		debug.PrintStack()
		os.Exit(1)
	}
}
func containsKind(kinds []reflect.Kind, kind reflect.Kind) bool {
	for i := 0; i < len(kinds); i++ {
		if kind == kinds[i] {
			return true
		}
	}

	return false
}
func IsNil(v interface{}) bool {
	if v == nil {
		return true
	}

	value := reflect.ValueOf(v)
	kind := value.Kind()
	isNilableKind := containsKind(
		[]reflect.Kind{
			reflect.Chan, reflect.Func,
			reflect.Interface, reflect.Map,
			reflect.Ptr, reflect.Slice},
		kind)

	if isNilableKind && value.IsNil() {
		return true
	}

	return false
}

func AssertNotNil(v interface{}, varname ...string) {
	var args []interface{}
	if len(varname) == 0 {
		args = []interface{}{"value"}
	} else {
		args = []interface{}{}
		for _, i := range varname {
			args = append(args, i)
		}
	}
	Assert(!IsNil(v), "Expected %s not to be nil.", args...)
}
