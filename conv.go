package utils

import (
	"fmt"
	"strconv"
)

// Convert 准备转换一个计算值。
// computedErr 表示事先计算时产生的错误，如：函数调用 `fn()(v int, err error)`，该错误始终生效，非 TryXXX 转换报 panic。
// zeroIfNil 表示在没有错误的情况下值为空，是否可以返回它们的默认值。
func Convert(v interface{}, computedErr error, zeroIfNil ...bool) Converter {
	b := true
	if len(zeroIfNil) > 0 {
		b = zeroIfNil[0]
	}

	return &DefaultConverter{
		DefaultValue: v,
		ComputedErr:  computedErr,
		ZeroIfNil:    b,
	}
}

func MoneyToYuan(v uint64) float64 {
	return float64(v) / float64(100)
}

// Converter 定义的一组转换函数，该接口用于指导其它依赖该转换的组件。
// 扩展 redis 或 sql 执行结果时刻返回该接口。
type Converter interface {
	Err() error

	Value() interface{}

	String() string
	TryString() (string, error)

	Bool() bool
	TryBool() (bool, error)

	Int16() int16
	TryInt16() (int16, error)
	Int() int
	TryInt() (int, error)
	Int64() int64
	TryInt64() (int64, error)

	Uint16() uint16
	TryUint16() (uint16, error)
	Uint() uint
	TryUint() (uint, error)
	Uint64() uint64
	TryUint64() (uint64, error)

	Float32() float32
	TryFloat32() (float32, error)
	Float64() float64
	TryFloat64() (float64, error)
}

type DefaultConverter struct {
	DefaultValue interface{}
	ComputedErr  error
	ZeroIfNil    bool
}

func panicIfError(err error) {
	if err != nil {
		//panic(err)
	}
}
func (c *DefaultConverter) checkDefault() (bool, error) {
	return c.ZeroIfNil && c.DefaultValue == nil, c.ComputedErr
}
func (c *DefaultConverter) Err() error {
	return c.ComputedErr
}

func (c *DefaultConverter) Value() interface{} {
	return c.DefaultValue
}

// ======== strings ========

func (c *DefaultConverter) str() string {
	if c.DefaultValue == nil {
		return ""
	}
	return fmt.Sprint(c.DefaultValue)
}

func (c *DefaultConverter) String() string {
	v, err := c.TryString()
	panicIfError(err)
	return v
}

func (c *DefaultConverter) TryString() (string, error) {
	if ok, err := c.checkDefault(); ok && err == nil {
		return "", nil
	}
	if v, ok := c.DefaultValue.(string); ok {
		return v, nil
	}
	return c.str(), nil
}

// ======== bools ========

func (c *DefaultConverter) Bool() bool {
	v, err := c.TryBool()
	panicIfError(err)
	return v
}

func (c *DefaultConverter) TryBool() (bool, error) {
	if ok, err := c.checkDefault(); ok && err == nil {
		return false, nil
	}
	if v, ok := c.DefaultValue.(bool); ok {
		return v, nil
	}
	return strconv.ParseBool(c.str())
}

// ======== int numbers ========

func (c *DefaultConverter) Int16() int16 {
	v, err := c.TryInt16()
	panicIfError(err)
	return v
}
func (c *DefaultConverter) TryInt16() (int16, error) {
	v, err := c.TryInt()
	if err != nil {
		return 0, err
	}
	return int16(v), nil
}

func (c *DefaultConverter) Int() int {
	v, err := c.TryInt()
	panicIfError(err)
	return v
}
func (c *DefaultConverter) TryInt() (int, error) {
	v, err := c.TryInt64()
	if err != nil {
		return 0, err
	}
	return int(v), nil
}

func (c *DefaultConverter) Int64() int64 {
	v, err := c.TryInt64()
	panicIfError(err)
	return v
}
func (c *DefaultConverter) TryInt64() (int64, error) {
	if ok, err := c.checkDefault(); ok && err == nil {
		return 0, nil
	}
	if v, ok := c.DefaultValue.(int64); ok {
		return v, nil
	}
	return strconv.ParseInt(c.str(), 10, 64)
}

func (c *DefaultConverter) Uint16() uint16 {
	v, err := c.TryUint16()
	panicIfError(err)
	return v
}
func (c *DefaultConverter) TryUint16() (uint16, error) {
	v, err := c.TryUint()
	if err != nil {
		return 0, err
	}
	return uint16(v), nil
}

func (c *DefaultConverter) Uint() uint {
	v, err := c.TryUint()
	panicIfError(err)
	return v
}
func (c *DefaultConverter) TryUint() (uint, error) {
	v, err := c.TryUint64()
	if err != nil {
		return 0, err
	}
	return uint(v), nil
}

func (c *DefaultConverter) Uint64() uint64 {
	v, err := c.TryUint64()
	panicIfError(err)
	return v
}
func (c *DefaultConverter) TryUint64() (uint64, error) {
	if ok, err := c.checkDefault(); ok && err == nil {
		return 0, nil
	}
	if v, ok := c.DefaultValue.(uint64); ok {
		return v, nil
	}
	return strconv.ParseUint(c.str(), 10, 64)
}

// ======== float numbers ========

func (c *DefaultConverter) Float32() float32 {
	v, err := c.TryFloat32()
	panicIfError(err)
	return v
}
func (c *DefaultConverter) TryFloat32() (float32, error) {
	v, err := c.TryFloat64()
	if err != nil {
		return 0, err
	}
	return float32(v), nil
}

func (c *DefaultConverter) Float64() float64 {
	v, err := c.TryFloat64()
	panicIfError(err)
	return v
}
func (c *DefaultConverter) TryFloat64() (float64, error) {
	if ok, err := c.checkDefault(); ok && err == nil {
		return 0, nil
	}

	if v, ok := c.DefaultValue.(float64); ok {
		return v, nil
	}
	return strconv.ParseFloat(c.str(), 64)
}

// https://baike.baidu.com/item/%E5%9B%9B%E8%88%8D%E5%85%AD%E5%85%A5%E4%BA%94%E6%88%90%E5%8F%8C/9062547?fr=aladdin
// https://studygolang.com/articles/8390
func (c *DefaultConverter) Money() float64 {
	return 1
}
