package types

import "log"

func IntDef(v interface{}, _default int) int {
	result, err := v.(int)
	if err {
		log.Printf("Error IntDef: %v", err)
		return _default
	}
	return result
}

func IntArrDef(v interface{}, _default []int) []int {
	result, err := v.([]int)
	if err {
		log.Printf("Error IntArrDef: %v", err)
		return _default
	}
	return result
}
