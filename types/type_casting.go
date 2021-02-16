package types

func IntDef(v interface{}, _default int) int {
	result, ok := v.(int)
	if !ok {
		return _default
	}
	return result
}

func IntArrDef(v interface{}, _default []int) []int {
	intArr, ok := v.([]interface{})
	if !ok {
		return _default
	}
	result := make([]int, len(intArr))
	for i, v := range intArr {
		result[i] = v.(int)
	}
	return result
}

func StrArrDef(v interface{}, _default []string) []string {
	strArr, ok := v.([]interface{})
	if !ok {
		return _default
	}
	result := make([]string, len(strArr))
	for i, v := range strArr {
		result[i] = v.(string)
	}
	return result
}

func BoolDef(v interface{}, _default bool) bool {
	result, ok := v.(bool)
	if !ok {
		return _default
	}
	return result
}
