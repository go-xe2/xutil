package xutil

import (
	"github.com/go-xe2/xtype"
	"reflect"
	"strings"
	"time"
)


func MapStruct(items *struct{}, fn MapFunc) map[string]interface{} {
	if items == nil {
		return nil
	}
	var results = make(map[string]interface{})
	vType := reflect.TypeOf(items)

	typeValue := reflect.ValueOf(items)
	count := vType.NumField()
	for i := 0; i < count; i++ {
		key := vType.Field(i).Name
		fieldValue := typeValue.Field(i)
		if fieldValue.IsNil() || !fieldValue.IsValid() {
			continue
		}
		tag := vType.Field(i).Tag
		value := fieldValue.Interface()
		item := fn(value, key, tag)
		results[key] = item
	}
	return results
}

func MapSlice(items []interface{}, fn MapFunc) []interface{} {
	var results []interface{}
	for i, v := range items {
		results = append(results, fn(v, i))
	}
	return results
}

func MapMap(items map[string]interface{}, fn MapFunc) map[string]interface{}  {
	var results = make(map[string]interface{})
	for key, val := range items {
		results[key] = fn(val, key)
	}
	return results
}

func Map(items interface{}, fn MapFunc) interface{} {
	vType := reflect.TypeOf(items).Kind()
	switch vType {
	case reflect.Slice:
		arr := xtype.New(items).SliceInterface()
		return MapSlice(arr, fn)
	case reflect.Map:
		m := xtype.New(items).MapStringInterface()
		return MapMap(m, fn)
	case reflect.Struct:
		if v, ok := items.(struct{}); ok {
			return MapStruct(&v, fn)
		}
		if v, ok := items.(*struct{}); ok {
			return MapStruct(v, fn)
		}
	default:
		return items
	}
	return items
}

func EachSlice(items []interface{}, fn EachFunc )  {
	if items == nil || fn == nil {
		return
	}
	for i, v := range items {
		fn(v, i)
	}
}

func EachMap(items map[string]interface{}, fn EachFunc)  {
	if items == nil || fn == nil {
		return
	}
	for key, v := range items {
		fn(v, key)
	}
}

func EachStruct(items *struct{}, fn EachFunc) {
	if items == nil || fn == nil {
		return
	}
	vType 		:= reflect.TypeOf(items)
	typeValue 	:= reflect.ValueOf(items)
	for i := 0; i < vType.NumField(); i ++ {
		key := vType.Field(i).Name
		fieldValue := typeValue.Field(i)
		if !fieldValue.IsValid() || fieldValue.IsNil() {
			continue
		}
		v := fieldValue.Interface()
		fn(v, key)
	}
}


func Each(items interface{}, fn EachFunc) {
	if items == nil || fn == nil {
		return
	}
	vType := reflect.TypeOf(items).Kind()
	switch vType {
	case reflect.Slice:
		arr := xtype.New(items).SliceInterface()
		EachSlice(arr, fn)
		break
	case reflect.Map:
		m := xtype.New(items).MapStringInterface()
		EachMap(m, fn)
		break
	case reflect.Struct:
		if v, ok := items.(struct{}); ok {
			EachStruct(&v, fn)
		}
		if v, ok := items.(*struct{}); ok {
			EachStruct(v, fn)
		}
		break
	}
}

func FindIndex(items []interface{}, fn FilterFunc) int {
	for i, item := range items {
		if fn(item, i) {
			return i
		}
	}
	return -1
}

func Find(items []interface{}, fn FilterFunc) interface{}  {
	for i, item := range items {
		if fn(item, i) {
			return item
		}
	}
	return nil
}

// 比较大小
func Compare(v1 interface{}, v2 interface{}, comparer ...CompareFunc) int {
	if len(comparer) > 0 {
		return comparer[0](v1, v2)
	}
	switch val := v1.(type) {
	case int:
		v2 := xtype.Int(v2)
		if v1 == v2 {
			return 0
		} else {
			return If(val > v2, 1, -1).(int)
		}
	case int8:
		v2 := xtype.Int8(v2)
		if v1 == v2 {
			return 0
		} else {
			return If(val > v2, 1, -1).(int)
		}
	case int16:
		v2 := xtype.Int16(v2)
		if v1 == v2 {
			return 0
		} else {
			return If(val > v2, 1, -1).(int)
		}
	case int32:
		v2 := xtype.Int32(v2)
		if v1 == v2 {
			return 0
		} else {
			return If(val > v2, 1, -1).(int)
		}
	case int64:
		v2 := xtype.Int64(v2)
		if v1 == v2 {
			return 0
		} else {
			return If(val > v2, 1, -1).(int)
		}
	case uint:
		v2 := xtype.Uint(v2)
		if v1 == v2 {
			return 0
		} else {
			return If(val > v2, 1, -1).(int)
		}
	case uint8:
		v2 := xtype.Uint8(v2)
		if v1 == v2 {
			return 0
		} else {
			return If(val > v2, 1, -1).(int)
		}
	case uint16:
		v2 := xtype.Uint16(v2)
		if v1 == v2 {
			return 0
		} else {
			return If(val > v2, 1, -1).(int)
		}
	case uint32:
		v2 := xtype.Uint32(v2)
		if v1 == v2 {
			return 0
		} else {
			return If(val > v2, 1, -1).(int)
		}
	case uint64:
		v2 := xtype.Uint64(v2)
		if v1 == v2 {
			return 0
		} else {
			return If(val > v2, 1, -1).(int)
		}
	case float32:
		v2 := xtype.Float32(v2)
		if v1 == v2 {
			return 0
		} else {
			return If(val > v2, 1, -1).(int)
		}
	case float64:
		v2 := xtype.Float64(v2)
		if v1 == v2 {
			return 0
		} else {
			return If(val > v2, 1, -1).(int)
		}
	case bool:
		v2 := xtype.Bool(v2)
		if v1 == v2 {
			return 0
		} else {
			bv1 := If(val, 1, 0).(int)
			bv2 := If(v2, 1, 0).(int)
			return If(bv1 > bv2, 1, -1).(int)
		}
	case string:
		v2 := xtype.String(v2)
		if v1 == v2 {
			return 0
		} else {
			return If(val > v2, 1, -1).(int)
		}
	case time.Duration:
		v2 := xtype.Duration(v2)
		if v1 == v2 {
			return 0
		} else {
			return If(val > v2, 1, -1).(int)
		}
	}
	return 1
}

// 比较是否相等
func Equired(v1 interface{}, v2 interface{}, comparer ...CompareFunc) bool {
	v := Compare(v1, v2, comparer...)
	return v == 0
}

// 条件中断,如果express != v 时中断
func Assert(express interface{}, v interface{}, msg ...string) {
	switch val := express.(type) {
	case error:
		panic(val)
		break
	case func() error:
		if err := val(); err != nil {
			panic(err.Error())
		}
		break
	case func() interface{}:
		v2 := val()
		if !Equired(v2, v) {
			panic(If(len(msg) > 0, strings.Join(msg, ","), "assert"))
		}
		break
	default:
		if !Equired(val, v) {
			panic(If(len(msg)> 0, strings.Join(msg, ","), "assert"))
		}
	}
}

// try ... catch
func TryCatch(try func(), catch func(err XError, args ...interface{})) {
	defer func() {
		if r := recover(); r != nil {
			if catch != nil {
				if e, ok := r.(XError); ok {
					catch(e)
				} else {
					e := NewError(xtype.String(r), -1)
					catch(e)
				}
			}
		}
	}()
	try()
}

// try finally
func TryFinally(try func(), finally func(...interface{})) {
	defer func() {
		if finally != nil {
			finally()
		}
	}()
	try()
}


