package utils_test

import (
	"testing"

	utils "github.com/junhwong/go-utils"
)

type Exp struct {
}

func (*Exp) Error() string {
	return "gggg"
}

func TestError(t *testing.T) {
	var err utils.Error = utils.Err("out.code", utils.Err("inner.code", "msg"))
	// c := &DefaultConverter{
	// 	DefaultValue: l,
	// }
	// v, err := c.TryInt()
	// if err == nil {
	// 	t.Fatal("err")
	// }
	// panic(*err)
	t.Log(err.Error(), "g")
}
