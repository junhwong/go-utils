package utils

import "testing"

func TestConv(t *testing.T) {
	var l int64 = 2342342424344565464
	c := &DefaultConverter{
		DefaultValue: l,
	}
	v, err := c.TryInt()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(v)
}
