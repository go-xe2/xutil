package xutil

import (
	"github.com/go-xe2/xtype"
	"time"
)


// 获取两项中的大值
func Max(value1, value2 interface{}) interface{} {
	switch v1 := value1.(type) {
	case int:
		v2 := xtype.New(value2).Int()
		return If(v1 > v2, value1, value2)
	case int8:
		v2 := xtype.Int8(value2)
		return If(v1 > v2, value1, value2)
	case int16:
		v2 := xtype.Int16(value2)
		return If(v1 > v2, value1, value2)
	case int32:
		v2 := xtype.Int32(value2)
		return If(v1 > v2, value1, value2)
	case int64:
		v2 := xtype.Int64(value2)
		return If(v1 > v2, value1, value2)
	case uint:
		v2 := xtype.Uint(value2)
		return If(v1 > v2, value1, value2)
	case uint8:
		v2 := xtype.Uint8(value2)
		return If(v1 > v2, value1, value2)
	case uint16:
		v2 := xtype.Uint16(value2)
		return If(v1 > v2, value1, value2)
	case uint32:
		v2 := xtype.Uint32(value2)
		return If(v1 > v2, value1, value2)
	case uint64:
		v2 := xtype.Uint64(value2)
		return If(v1 > v2, value1, value2)
	case bool:
		return If(v1, value1, value2)
	case string:
		v2 := xtype.String(value2)
		return If(v1 > v2, value1, value2)
	case float32:
		v2 := xtype.Float32(value2)
		return If(v1 > v2, value1, value2)
	case float64:
		v2 := xtype.Float64(value2)
		return If(v1 > v2, value1, value2)
	case time.Duration:
		v2 := xtype.Duration(value2)
		return If(v1 > v2, value1, value2)
	default:
		return value1
	}
}

// 获取两项中的小值
func Min(value1, value2 interface{}) interface{}  {
	switch v1 := value1.(type) {
	case int:
		v2 := xtype.New(value2).Int()
		return If(v1 < v2, value1, value2)
	case int8:
		v2 := xtype.Int8(value2)
		return If(v1 < v2, value1, value2)
	case int16:
		v2 := xtype.Int16(value2)
		return If(v1 < v2, value1, value2)
	case int32:
		v2 := xtype.Int32(value2)
		return If(v1 < v2, value1, value2)
	case int64:
		v2 := xtype.Int64(value2)
		return If(v1 < v2, value1, value2)
	case uint:
		v2 := xtype.Uint(value2)
		return If(v1 < v2, value1, value2)
	case uint8:
		v2 := xtype.Uint8(value2)
		return If(v1 < v2, value1, value2)
	case uint16:
		v2 := xtype.Uint16(value2)
		return If(v1 < v2, value1, value2)
	case uint32:
		v2 := xtype.Uint32(value2)
		return If(v1 < v2, value1, value2)
	case uint64:
		v2 := xtype.Uint64(value2)
		return If(v1 < v2, value1, value2)
	case bool:
		return If(v1, value1, value2)
	case string:
		v2 := xtype.String(value2)
		return If(v1 < v2, value1, value2)
	case float32:
		v2 := xtype.Float32(value2)
		return If(v1 < v2, value1, value2)
	case float64:
		v2 := xtype.Float64(value2)
		return If(v1 < v2, value1, value2)
	case time.Duration:
		v2 := xtype.Duration(value2)
		return If(v1 < v2, value1, value2)
	default:
		return value1
	}
}

// 获取多项同的大值
func MaxEx(values ...interface{}) interface{} {
	if len(values) == 0 {
		return nil
	}
	var max = values[0]
	for _, v := range values {
		max = Max(max, v)
	}
	return max
}

// 获取多值中的小值
func MinEx(values ...interface{}) interface{} {
	if len(values) == 0 {
		return nil
	}
	var min = values[0]
	for _, v := range values {
		min = Min(min, v)
	}
	return min
}




