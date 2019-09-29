package xutil

import (
	"reflect"
	"testing"
)

func TestParseTag(t *testing.T) {
	var tag = `name:"fdfdf", sex:"sex" age="{old:3,new:33}",number:56,str:"string1"`
	m := ParseTag(reflect.StructTag(tag))
	t.Log(m)

	v := FindTagValue(reflect.StructTag(tag), "number")
	t.Log(v)

	v = FindTagValue(reflect.StructTag(tag), "str")
	t.Log(v)
}
