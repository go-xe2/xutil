package xutil

type O = interface {}
type M = map[string]O
type A = []O
type AM = []M

// 比较函数
type CompareFunc func(value1, value2 interface{}) int
type MapFunc func(item interface{}, indexOrKey interface{}, args ...interface{}) interface{}
type EachFunc func(item interface{}, indexOrKey interface{}) interface{}
type FilterFunc func(item interface{}, indexOrKey interface{}) bool
