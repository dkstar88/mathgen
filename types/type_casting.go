package types

import (
	"log"
)

func IntDef(v interface{}, _default int) int {
	result, ok := v.(int)
	if !ok {
		log.Printf("Error IntDef: %v", v)
		return _default
	}
	return result
}

func IntArrDef(v interface{}, _default []int) []int {
	intArr, ok := v.([]interface{})
	if !ok {
		log.Printf("Error IntArrDef: %v", v)
		return _default
	}
	result := make([]int, len(intArr))
	for i, v := range intArr {
		result[i] = v.(int)
	}
	return result
}

