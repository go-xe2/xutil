package xutil

func If(express bool, trueValue, falseValue interface{}) interface{} {
	if express {
		return trueValue
	} else {
		return falseValue
	}
}

func If_str(express bool, trueValue, falseValue string) string {
	return If(express, trueValue, falseValue).(string)
}

func If_int(express bool, trueValue, falseValue int) int  {
	return If(express, trueValue, falseValue).(int)
}

func If_int8(express bool, trueValue, falseValue int8) int8  {
	return If(express, trueValue, falseValue).(int8)
}

func If_int16(express bool, trueValue, falseValue int8) int8  {
	return If(express, trueValue, falseValue).(int8)
}

func If_int32(express bool, trueValue, falseValue int32) int32 {
	return If(express, trueValue, falseValue).(int32)
}

func If_int64(express bool, trueValue, falseValue int64) int64  {
	return If(express, trueValue, falseValue).(int64)
}

func If_uint(express bool, trueValue, falseValue uint) uint  {
	return If(express, trueValue, falseValue).(uint)
}

func If_uint8(express bool, trueValue, falseValue uint8) uint8  {
	return If(express, trueValue, falseValue).(uint8)
}

func If_uint16(express bool, trueValue, falseValue uint16) uint16  {
	return If(express, trueValue, falseValue).(uint16)
}

func If_uint32(express bool, trueValue, falseValue uint32) uint32  {
	return If(express, trueValue, falseValue).(uint32)
}

func If_uint64(express bool, trueValue, falseValue uint64) uint64  {
	return If(express, trueValue, falseValue).(uint64)
}

func If_float(express bool, trueValue, falseValue float32) float32  {
	return If(express, trueValue, falseValue).(float32)
}

func If_float64(express bool, trueValue, falseValue float64) float64  {
	return If(express, trueValue, falseValue).(float64)
}

func If_bool(express bool, trueValue, falseValue bool) bool  {
	return If(express, trueValue, falseValue).(bool)
}

func If_byte(express bool, trueValue, falseValue byte) byte  {
	return If(express, trueValue, falseValue).(byte)
}

func If_method(express bool, trueMethod, falseMethod func()) {
	if express {
		trueMethod()
	} else {
		falseMethod()
	}
}

func If_func(express bool, trueFunc, falseFunc func() O) O {
	if express {
		return trueFunc()
	} else {
		return falseFunc()
	}
}
