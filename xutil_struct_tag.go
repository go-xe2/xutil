package xutil

import (
	"encoding/json"
	"reflect"
)

// 解析struct tag
// 允许是json字符串，逗号分割，空格分割
func ParseTag(tag reflect.StructTag) map[string]interface{} {
	// 从缓存中读取
	var results = make(map[string]interface{})
	if tag == "" {
		return results
	}
	var v map[string]interface{}
	if err := json.Unmarshal([]byte(tag), &v); err == nil {
		return v
	}
	len := len(tag)
	var tmp string
	var key string
	var value string
	var strQuo = ""
	var searchMode = 0 // 0:查找键名,1:查找值,2:查找字符串
	for i := 0; i < len; i ++ {
		c := tag[i]
		switch c {
		case ':', '=':
			if searchMode == 0 && tmp != "" {
				key = tmp
				tmp = ""
				strQuo = ""
				searchMode = 1
			} else {
				tmp += string(c)
			}
			break
		case ',', ' ':
			// 项分割符
			if searchMode == 2 {
				// 字符串未结束，当字符串处理
				tmp += string(c)
			} else {
				// 键值对结束
				if searchMode == 1 && key != "" {
					value = tmp
					results[key] = value
				}
				searchMode = 0
				key = ""
				tmp = ""
				strQuo = ""
			}
		case '"', '\'':
			if searchMode == 1 {
				// 查找字符串
				strQuo = string(c)
				tmp = ""
				searchMode = 2
				if strQuo == "" {
					strQuo = string(c)
					tmp = ""
					searchMode = 2
				}
			} else {
				// 找到第二个引号
				if searchMode == 2 && string(c) == strQuo {
					searchMode = 1 // 字符串查找结束
					continue
				}
				tmp += string(c)
			}
			break
		default:
			tmp = tmp + string(c)
		}
	}
	if key != "" && searchMode <= 1 {
		// 最后一项
		results[key] = tmp
	}
	// 缓存
	return results
}

// 查找值StructTag中的键值
func FindTagValue(tag reflect.StructTag, field string) interface{} {
	m := ParseTag(tag)
	if s, ok := m[field]; ok {
		return s
	}
	return nil
}
