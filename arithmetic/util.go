package arithmetic

import (
	"math/rand"
	"strconv"
	"strings"
)

func detectDifficultyInt(num int) int {
	switch {
	case num > 1000:
		return 4
	case num > 100:
		return 3
	case num > 10:
		return 2
	default:
		return 1
	}
}

func joinInts(intArray []int, delim string) string {

	b := make([]string, len(intArray))
	for i, v := range intArray {
		b[i] = strconv.Itoa(v)
	}
	return strings.Join(b, delim)
}

//func randomDigits(digits int) int {
//	var min, max int
//	if digits <= 1 {
//		min = 0
//	} else {
//		min = int(math.Pow10(digits-1))
//	}
//	max = int(math.Pow10(digits))-1
//	return rand.Intn(max-min) + min
//}

func randomMinMax(min, max int) int {
	return rand.Intn(max-min) + min
}
